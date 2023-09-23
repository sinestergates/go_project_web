package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var list_urls = [5]string{
	"https://golangify.com/array", 
	"https://dev-gang.ru/article/golang-making-http-requests/", 
	"https://ru.stackoverflow.com/questions/988086",
	"https://im.astralinux.ru/astralinux/channels/town-square",
	"https://vegaspro-rus.ru/kak-uznat-xarakteristiki-kompyutera-v-astra-linux-poshagovaya-instrukciya/",
}


func def(test string, test2 string){
	fmt.Printf("%s, %s!", test, test2)

}


func getRoot(w http.ResponseWriter, r *http.Request) {
	//fmt.Print(r)
	io.WriteString(w, "This is my website!\n")
}

func getApi(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r)
	io.WriteString(w, "This is my website!\n")
}

func get_responce(w http.ResponseWriter, r *http.Request){
    c:= make(chan string)

	for i, url := range list_urls {
		go gorutins(url, c, i)
		//gorutins(url, c, i)
	}
	for range list_urls {
        url := <-c // Получает значение от канала
        io.WriteString(w, url)
    }
	
}


func gorutins(url string, c chan string, i int) {
	resp, err:= http.Get(url)
	fmt.Print(i)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	bodyBytes, error := io.ReadAll(resp.Body)
    if error != nil {
        fmt.Print(error)
    }
    bodyString := string(bodyBytes)
	c <- bodyString
	
}

func main() {
	
	def("Вася", "пупкин")
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/get_to_sites", get_responce)
	err := http.ListenAndServe(":3334", nil)
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}