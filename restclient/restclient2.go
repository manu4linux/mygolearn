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

/********************************************************/
//client code
/********************************************************/
/*
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

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var urlstr1 string = "http://teamcity.cvs-a.ula.comcast.net:8111/guestAuth/app/rest/builds/buildType:(id:CptServers_CptServersLegac_ReleaseBuildTvworksServerParentPom)"
var urlstr2 string = "http://teamcity.cvs-a.ula.comcast.net:8111/guestAuth/app/rest/builds/buildType:(id:CptServers_CptServersLegac_ReleaseBuildTvworksServerParentPom)/number"

type Message1 struct {
	Id            json.Number `json:"id,Number"`
	BuildTypeId   string
	Number        json.Number `json:"number,Number"`
	Status        string
	State         string
	BranchName    string
	DefaultBranch bool
	WebUrl        string
}

func main() {

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", urlstr1, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Accept", "application/json")

	resp, err := netClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	//_, err = io.Copy(os.Stdout, resp.Body)
	//fmt.Printf("Error: %v\n", err)

	var m Message1
	decoder := json.NewDecoder(resp.Body)
	//m := Fo{}
	err = decoder.Decode(&m)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(m)
	/*type Animal struct {
		Name  string
		Order string
	}

	var animals []Animal
	err := json.Unmarshal(resp, &animals)
	*/
	// explore response object

	//fmt.Printf("Response Body: %v\n", (resp.Body.Read)) // or resp.String() or string(resp.Body())
}
