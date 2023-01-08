package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func onlineRequest() {
	resp, err := http.Get("https://www.shufersal.co.il/online/search?text=" + url.QueryEscape("קישוא עגול"))

	if err != nil {
		fmt.Println(err.Error())
	} else {
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			// plain := html2text.HTML2Text(body)
			// doc, err := html.
			// html.Parse(body)

			doc, err := html.Parse(strings.NewReader(body))
			fmt.Println(string(body))
		}
	}
}

func storeRequest() {
	resp, err := http.Get("http://pricesprodpublic.blob.core.windows.net/pricefull/PriceFull7290027600007-001-202301080300.gz?" +
		"sv=2014-02-14&sr=b&sig=7K47Y8Kk%2BY1ECR9o0x53etHuay%2BK4%2Fs7UvvQuDphxM4%3D&se=2023-01-08T14%3A17%3A30Z&sp=r")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		// zr, err := gzip.NewReader(resp.Body)

		// if err != nil {
		// 	fmt.Println(err.Error())
		// } else {
		// 	fmt.Println(zr)
		// }
		fmt.Println(resp.Header)
		// fmt.Println(resp.b)

		// client := new(http.Client)

		// request, err := http.NewRequest("GET", "http://stackoverflow.com", nil)
		// request.Header.Add("Accept-Encoding", "gzip")

		// response, err := client.Do(request)
		// defer response.Body.Close()

		// // Check that the server actually sent compressed data
		// var reader io.ReadCloser
		// switch response.Header.Get("Content-Encoding") {
		// case "gzip":
		//     reader, err = gzip.NewReader(response.Body)
		//     defer reader.Close()
		// default:
		//     reader = response.Body
		// }

		// io.Copy(os.Stdout, reader) // print html to standard out

	}
}

func main() {
	// storeRequest()
	onlineRequest()
}
