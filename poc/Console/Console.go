package Console

import (
	"Weblogic_Scanner/UA"
	"fmt"
	"net/http"
)

func alive(ip string, port string) int {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	url := "http://" + ip + ":" + port + "/console/login/LoginForm.jsp"
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", UA.UA())
	if err != nil {
		return -1
	}
	response, err := client.Do(request)
	if err != nil {
		return -1
	}
	defer response.Body.Close()

	status := response.StatusCode
	return status
}

func Run(ip string, port string) {

	if alive(ip, port) == 200 {
		url := "http://" + ip + ":" + port + "/console/login/LoginForm.jsp"
		fmt.Println("[+] Weblogic console is exposed!")
		fmt.Printf("[+] The path is: %s \n", url)
	} else {
		fmt.Printf("[-] Weblogic console not found!\n")
	}

}
