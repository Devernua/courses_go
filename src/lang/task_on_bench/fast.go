package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type User struct {
	Browsers []string `json:"browsers"`
	Company string `json:"company"`
	Country string `json:"country"`
	Email string `json:"email"`
	Job string `json:"job"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	seenBrowsers := make(map[string]struct{})
	uniqueBrowsers := 0
	foundUsers := ""
	i := -1

	for scanner.Scan() {
		i++

		var user User
		err := json.Unmarshal([]byte(scanner.Text()), &user)
		if err != nil {
			panic(err)
		}

		hasAndroid := false
		hasMSIE := false
		for _, browser := range user.Browsers {
			isAndroid := strings.Contains(browser, "Android")
			isMSIE := strings.Contains(browser, "MSIE")

			if isAndroid || isMSIE {
				if _, ok := seenBrowsers[browser]; !ok {
					seenBrowsers[browser] = struct{}{}
					uniqueBrowsers++
				}
				hasAndroid = isAndroid || hasAndroid
				hasMSIE = isMSIE || hasMSIE
			}
		}
		if hasAndroid && hasMSIE {
			emailParts := strings.Split(user.Email, "@")
			foundUsers += fmt.Sprintf("[%d] %s <%s>\n", i, user.Name, strings.Join(emailParts, " [at] "))
		}
	}

	fmt.Fprintln(out, "found users:\n"+foundUsers)
	fmt.Fprintln(out, "Total unique browsers", uniqueBrowsers)
}