package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func title(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func main() {
	t1 := title("https://www.abordagemnoticias.com/", "https://www.queroquitar.com.br/")
	t2 := title("https://letscode.com.br/", "https://www.udemy.com/")

	fmt.Println("Primeiras respostas: ", <-t1, "&", <-t2)
	fmt.Println("Segundas respostas: ", <-t1, "&", <-t2)
}
