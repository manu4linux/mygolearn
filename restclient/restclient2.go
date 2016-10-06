/*package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

func main() {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Hour)
	}))
	defer svr.Close()
	fmt.Println("making request")
	http.Get(svr.URL)
	fmt.Println("finished request")
}
*/

/*
//server code 

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
*/


/*
//client code

package main

import (
	"fmt"
	"github.com/go-resty/resty"
)

func main() {
	// GET request
	resp, err := resty.R().Get("http://teamcity.cvs-a.ula.comcast.net:8111/guestAuth/app/rest/builds/buildType:(id:CptServers_CptServersLegac_ReleaseBuildTvworksServerParentPom)/number")

	// explore response object
	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Recevied At: %v", resp.ReceivedAt())
	fmt.Printf("\nResponse Body: %v", resp) // or resp.String() or string(resp.Body())
}
*/
