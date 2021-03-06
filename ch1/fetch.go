package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "get url %s error %v\n", url, err)
			os.Exit(1)
		}
		content, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "read body %s error %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", content)
	}
}