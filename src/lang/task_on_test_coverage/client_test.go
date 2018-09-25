package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"
)

func parseArgs(values url.Values) (*SearchRequest, error) {
	req := &SearchRequest{}
	var err error

	req.Limit, err = strconv.Atoi(values.Get("limit"))
	if err != nil {
		return nil, err
	}

	req.Offset, err = strconv.Atoi(values.Get("offset"))
	if err != nil {
		return nil, err
	}

	req.Query = values.Get("query")
	req.OrderField = values.Get("order_field") // TODO: check only needed fields
	req.OrderBy, err = strconv.Atoi(values.Get("order_by"))
	if err != nil {
		return nil, err
	}
	return req, nil
}

type UserXML struct {
	Id        int    `xml:"id"`
	FirstName string `xml:"first_name"`
	LastNane  string `xml:"last_name"`
	Age       int    `xml:"age"`
	About     string `xml:"about"`
	Gender    string `xml:"gender"`
}

type UsersXML struct {
	XMLName xml.Name  `xml:"root"`
	Users   []UserXML `xml:"row"`
}

// TODO: use filters and another futures
func getUsers() []User {
	xmlFile, err := os.Open("dataset.xml")
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	users := UsersXML{}
	xml.Unmarshal(byteValue, &users)

	var result []User
	for _, user := range users.Users {
		result = append(result, User{
			Id:     user.Id,
			Name:   user.FirstName + " " + user.LastNane,
			Age:    user.Age,
			About:  user.About,
			Gender: user.Gender,
		})
	}
	return result
}

type SearchHandler struct {
	Users []User
}

func SearchServer(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("AccessToken") != "123" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// TODO: filtering
	users := getUsers()
	req, err := parseArgs(r.URL.Query())
	if err != nil {
		panic(err) // TODO: may be return internal error
	}

	if req.OrderField != "Id" && req.OrderField != "Age" && req.OrderField != "Name" && req.OrderField != "" {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := SearchErrorResponse{Error: "ErrorBadOrderField"}
		res, err := json.Marshal(errorResponse)
		if err != nil {
			panic(err)
		}

		w.Write(res)
		return
	}

	res, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	w.Write(res)
}

func TestSimple(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(SearchServer))
	client := SearchClient{AccessToken: "123", URL: server.URL}
	resp, err := client.FindUsers(SearchRequest{
		Limit:      3,
		Offset:     0,
		Query:      "esse",
		OrderField: "Age",
		OrderBy:    OrderByAsc,
	})
	if err != nil {
		t.Fatal("err response", err)
	}
	//for _, user := range resp.Users {
	//	fmt.Println(user)
	//}
	if len(resp.Users) == 0 {
		t.Fatal("users not found")
	}
}

func TestBadAccessToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(SearchServer))
	client := SearchClient{AccessToken: "BAD_ACCESS_TOKEN", URL: server.URL}
	_, err := client.FindUsers(SearchRequest{
		Limit:      3,
		Offset:     0,
		Query:      "esse",
		OrderField: "Age",
		OrderBy:    OrderByAsc,
	})
	if err == nil && err !=  fmt.Errorf("Bad AccessToken"){
		t.Fatal("err response", err)
	}
}

