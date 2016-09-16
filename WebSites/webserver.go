package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var fileBytes, err = ioutil.ReadFile(r.URL.Path[1:])
	//var _, err = ioutil.ReadFile(r.URL.Path[1:])

	fmt.Println("\n------\n")
	fmt.Println([]byte("<html>\n<head>\n    <link rel=\"stylesheet\" type=\"text/css\" href=\"styles/style.css\">\n</head>\n<body>\n    <div>Hello, my name is John Carpenter and I love my wife, Lara Sozansky Carpenter</div>\n</body>\n</html>"))
	fmt.Println("\n------\n")
	fmt.Println(fileBytes)
	fmt.Println("\n------\n")

	if err == nil {
		//fmt.Fprintf(w, "<html>\n<body>\n    Hello, my name is John Carpenter and I love my wife, Lara Sozansky Carpenter\n</body>\n</html>")
		fmt.Fprintf(w, "\n%s\n", string(fileBytes))
	} else {
		fmt.Fprintf(w, "There was an error getting the page you requesting")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
