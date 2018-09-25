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
	"strings"
	"testing"
	"time"
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

	res, err := json.Marshal(users[req.Offset:req.Limit])
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

	if len(resp.Users) == 0 {
		t.Fatal("users not found")
	}
}

func TestSimpleMaxLimit(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(SearchServer))
	client := SearchClient{AccessToken: "123", URL: server.URL}
	resp, err := client.FindUsers(SearchRequest{
		Limit:      40,
		Offset:     0,
		Query:      "esse",
		OrderField: "Age",
		OrderBy:    OrderByAsc,
	})
	if err != nil {
		t.Fatal("err response", err)
	}

	if len(resp.Users) > 25 {
		t.Fatal("users not found")
	}
}

func TestSubZeroLimit(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(SearchServer))
	client := SearchClient{AccessToken: "123", URL: server.URL}
	_, err := client.FindUsers(SearchRequest{
		Limit:      -3,
		Offset:     0,
		Query:      "esse",
		OrderField: "Age",
		OrderBy:    OrderByAsc,
	})
	if err == nil && err != fmt.Errorf("limit must be > 0") {
		t.Fatal("err response", err)
	}
}

func TestSubZeroOffset(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(SearchServer))
	client := SearchClient{AccessToken: "123", URL: server.URL}
	_, err := client.FindUsers(SearchRequest{
		Limit:      3,
		Offset:     -2,
		Query:      "esse",
		OrderField: "Age",
		OrderBy:    OrderByAsc,
	})
	if err == nil && err != fmt.Errorf("offset must be > 0") {
		t.Fatal("err response", err)
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
	if err == nil && err != fmt.Errorf("Bad AccessToken") {
		t.Fatal("err response", err)
	}
}

func TestTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){time.Sleep(time.Second)}))

	client := SearchClient{AccessToken: "123", URL: server.URL}
	_, err := client.FindUsers(SearchRequest{
		Limit:      3,
		Offset:     0,
		Query:      "esse",
		OrderField: "Age",
		OrderBy:    OrderByAsc,
	})
	fmt.Println(err)
	if err == nil || !strings.Contains(err.Error(), "timeout") {
		t.Fatal("err response:", err)
	}
}

func TestEmptyURL(t *testing.T) {
	//server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){fatal}))

	client := SearchClient{AccessToken: "123", URL: ""}
	_, err := client.FindUsers(SearchRequest{
		Limit:      3,
		Offset:     0,
		Query:      "esse",
		OrderField: "Age",
		OrderBy:    OrderByAsc,
	})
	fmt.Println(err)
	if err == nil || !strings.Contains(err.Error(), "unknown error") {
		t.Fatal("err response:", err)
	}
}


func TestFatalServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){w.WriteHeader(http.StatusInternalServerError)}))
	client := SearchClient{AccessToken: "123", URL: server.URL}
	_, err := client.FindUsers(SearchRequest{
		Limit:      3,
		Offset:     0,
		Query:      "esse",
		OrderField: "Age",
		OrderBy:    OrderByAsc,
	})
	fmt.Println(err)
	if err == nil || !strings.Contains(err.Error(), "SearchServer fatal error") {
		t.Fatal("err response:", err)
	}
}

func TestBadRequestWithBadBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){w.WriteHeader(http.StatusBadRequest)}))
	client := SearchClient{AccessToken: "123", URL: server.URL}
	_, err := client.FindUsers(SearchRequest{
		Limit:      3,
		Offset:     0,
		Query:      "esse",
		OrderField: "Age",
		OrderBy:    OrderByAsc,
	})
	fmt.Println(err)
	if err == nil || !strings.Contains(err.Error(), "cant unpack error json") {
		t.Fatal("err response:", err)
	}
}

func TestBadOrderField(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusBadRequest)
		body, _ :=json.Marshal(&SearchErrorResponse{Error: "ErrorBadOrderField"})
		w.Write(body)
	}))
	client := SearchClient{AccessToken: "123", URL: server.URL}
	_, err := client.FindUsers(SearchRequest{
		Limit:      3,
		Offset:     0,
		Query:      "esse",
		OrderField: "Age",
		OrderBy:    OrderByAsc,
	})
	fmt.Println(err)
	if err == nil || !strings.Contains(err.Error(), "OrderFeld") {
		t.Fatal("err response:", err)
	}
}

func TestUnknownErrorSearchResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusBadRequest)
		body, _ :=json.Marshal(&SearchErrorResponse{Error: "UnknownError"})
		w.Write(body)
	}))
	client := SearchClient{AccessToken: "123", URL: server.URL}
	_, err := client.FindUsers(SearchRequest{
		Limit:      3,
		Offset:     0,
		Query:      "esse",
		OrderField: "Age",
		OrderBy:    OrderByAsc,
	})
	fmt.Println(err)
	if err == nil || !strings.Contains(err.Error(), "unknown bad request error") {
		t.Fatal("err response:", err)
	}
}