package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	siteToReach := "https://google.com"
	resp, err := http.Get(siteToReach)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	} else {
		// The reponse is of a Reader interface type. It needs to be read
		// Basic way would be to allot a memory and read the response and store it in a byte slice
		val := basicPrint(resp)
		fmt.Println("Response from ", siteToReach, " is: \n", string(val))

		// alternatively we can use the writer interface to make it cleaner and consice
		io.Copy(os.Stdout, resp.Body) // Copy(writer , reader)
	}
}

func basicPrint(r *http.Response) []byte {
	bt := make([]byte, 99999) // 99999 is the size alloted for bt
	r.Body.Read(bt)
	return bt
}
