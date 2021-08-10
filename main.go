package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var jsonStr = []byte(`{"platform_source":"web", "time_spent_in_seconds":5, "verbose_name":"word_bank"}`)
	req, err := http.NewRequest("POST", "https://app.voxy.com/api/v0/time-on-tasks/", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	//* * * * *   /usr/local/bin/go run ~/Desktop/projects/main.go >> ~/Desktop/cron.txt 2>&1
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjQxZTc2ZDAwOWYyOWZkYWYwYjkyMGI1MDIwZWU0OSIsInVzZXJfaWQiOjYwMTk1OTksImVtYWlsIjoiYWxpLnBhcm5hQG1pbmljbGlwLmNvbSIsImV4cCI6MTYzMTI4MzQyMX0.hXK9kPncMzX7tahehaWpbcyVecvpOriZFayE3i43uOU")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36")
	req.Header.Set("Origin", "https://app.voxy.com")
	req.Header.Set("Referer", "https://app.voxy.com/v2/")
	req.Header.Set("Host", "https://app.voxy.com/v2/")
	req.Header.Set("Content-Length", "78")
	req.Header.Set("Set-cookie", "sessionid=k7p3hcd2btyx7nt801ql3jr730lr82bm; expires=Fri, 10-Sep-2021 16:54:56 GMT; httponly; Max-Age=2678400; Path=/")

	client := &http.Client{}
	resp, err := client.Do(req)

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}
