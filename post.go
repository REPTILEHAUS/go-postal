​package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
)

func main() {
    // Declare http client
    client := &http.Client{}
	
    hdr := map[string]string{"Host": "www.example.com", "Content-Type": "application/x-www-form-urlencoded", "User-Agent": "Mozilla/5.0 (X11; Linux i686; rv:52.0) Gecko/20100101 Firefox/52.0"}

    hdr["Host"] = "www.example.com"

    // Declare post data
    PostData := strings.NewReader("subject=hey+you&message=how+are+you+doing&uid=123456789")

    // Declare HTTP Method and Url
    req, err := http.NewRequest("POST", "http://www.example.com/messages_send.php?id=123456789", PostData)

    // Set cookie
    req.Header.Set("Host", "www.example.com")
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux i686; rv:52.0) Gecko/20100101 Firefox/52.0")
    req.Header.Set("Referer", "http://www.example.com/messages.php?id=123456789")
    req.Header.Set("Cookie", "PHPSESSID=123456789; AWSELB=123456789")
    req.Header.Set("Content-Length", "63")
    req.Header.Set("Upgrade-Insecure-Requests", "1")
    req.Header.Set("Connection", "close")
    req.Header.Set("Accept-Language", "en-US,en;q=0.5")
    req.Header.Set("Accept-Encoding", "gzip, deflate")

    resp, err := client.Do(req)
    // Read response
    data, err := ioutil.ReadAll(resp.Body)

    // error handle
    if err != nil {
        fmt.Printf("error = %s \n", err)
    }

    // Print response
    fmt.Printf("Response = %s", string(data))
}