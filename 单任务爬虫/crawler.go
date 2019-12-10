package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:status code:", resp.StatusCode)
		return
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%s\n", all)

	PrintCityList(all)
}

func PrintCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`)
	all := re.FindAllSubmatch(contents, -1)
	for _, c := range all {
		fmt.Printf("City\t%s\t,URL\t%s", c[2], c[1])
		fmt.Println()
	}
}
