package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
	"io"
	"os"
	"strings"
)

type User struct {
	Browsers []string `json:"browsers"`
	Email string `json:"email"`
	Name string `json:"name"`
}

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson2bc03518DecodeLangTaskOnBench(in *jlexer.Lexer, out *User) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "browsers":
			if in.IsNull() {
				in.Skip()
				out.Browsers = nil
			} else {
				in.Delim('[')
				if out.Browsers == nil {
					if !in.IsDelim(']') {
						out.Browsers = make([]string, 0, 4)
					} else {
						out.Browsers = []string{}
					}
				} else {
					out.Browsers = (out.Browsers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Browsers = append(out.Browsers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "email":
			out.Email = string(in.String())
		case "name":
			out.Name = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2bc03518DecodeLangTaskOnBench(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2bc03518DecodeLangTaskOnBench(l, v)
}

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	seenBrowsers := make(map[string]struct{})
	i := -1
	var user User
	var tmp int
	hasAndroid := false
	hasMSIE := false
	isAndroid := false
	isMSIE := false

	fmt.Fprintln(out, "found users:")

	for scanner.Scan() {
		i++
		user.UnmarshalJSON([]byte(scanner.Text()))

		hasAndroid = false
		hasMSIE = false
		for _, browser := range user.Browsers {
			isAndroid = strings.Contains(browser, "Android")
			isMSIE = strings.Contains(browser, "MSIE")

			if isAndroid || isMSIE {
				if _, ok := seenBrowsers[browser]; !ok {
					seenBrowsers[browser] = struct{}{}
				}
				hasAndroid = isAndroid || hasAndroid
				hasMSIE = isMSIE || hasMSIE
			}
		}
		if hasAndroid && hasMSIE {
			tmp = strings.IndexAny(user.Email, "@")
			fmt.Fprintf(out, "[%d] %s <%s [at] %s>\n", i, user.Name, user.Email[0:tmp], user.Email[tmp + 1: len(user.Email)])
			//emailParts := strings.Split(user.Email, "@")
			//fmt.Fprintf(out, "[%d] %s <%s>\n", i, user.Name, strings.Join(emailParts, " [at] "))
		}
	}

	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}