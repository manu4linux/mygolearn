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
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var urlstr1 string = "http://teamcity.cvs-a.ula.comcast.net:8111/guestAuth/app/rest/builds/buildType:(id:CptServers_CptServersLegac_ReleaseBuildTvworksServerParentPom)"
var urlstr2 string = "http://teamcity.cvs-a.ula.comcast.net:8111/guestAuth/app/rest/builds/buildType:(id:CptServers_CptServersLegac_ReleaseBuildTvworksServerParentPom)/number"

type Message1 struct {
	Id            json.Number `json:"id,Number"`
	BuildTypeId   string      `json:"buildTypeId"`
	Number        string      `json:"number"`
	Status        string      `json:"status"`
	State         string      `json:"state"`
	BranchName    string      `json:"branchName"`
	DefaultBranch bool        `json:"defaultBranch"`
	WebUrl        string      `json:"webUrl"`
}

type component struct {
	Name           string      `json:"name"`
	PropertyName   string      `json:"property_name"`
	CurrentVersion string      `json:"current_version"`
	BuildOrder     json.Number `json:"build_order,Number"`
	BuildTypeId    string      `json:"BuildTypeId"`
}

type componentjson struct {
	Components []component `json:"components"`
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
	/*
		resp, err := netClient.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
	*/
	//_, err = io.Copy(os.Stdout, resp.Body)
	//fmt.Printf("Error: %v\n", err)
	/*

		decoder := json.NewDecoder(resp.Body)
		//m := Fo{}
		err = decoder.Decode(&m)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(m)
		//defer decoder.Close()
	*/

	/****************/
	//Read from a json file
	/****************/
	pathtofile := "./restclient/ServerRefresh.json"
	file, e := ioutil.ReadFile(pathtofile)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	//fmt.Printf("%s\n", string(file))

	var mobj componentjson

	err = json.Unmarshal(file, &mobj)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("first attempt")
	fmt.Printf("Results: %v\n\n\n", mobj)
	log.Println(mobj)
	//numberofcomp := len(mobj.Components)
	// fmt.Printf("Results: %v\n\n\n", mobj.Components)
	// log.Println(mobj.Components)
	// fmt.Printf("Results: %v\n\n\n", mobj.Components[2])
	// log.Println(mobj.Components[3])
	/*
		configFile, err := os.Open(pathtofile)
		if err != nil {
			log.Println("opening config file", err.Error())
		}

		decoder2 := json.NewDecoder(configFile)
		err = decoder2.Decode(&mobj)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("second attempt")
		fmt.Printf("Results: %v\n\n\n", mobj)
		log.Println(mobj)
	*/
	/***************************/
	//run the rest url to get json file
	/***************************/
	//resp := &http.Response{}
	//decoder := json.NewDecoder("")
	var m Message1
	//mm := make([]Message1, numberofcomp)
	urlstr3 := "http://teamcity.cvs-a.ula.comcast.net:8111/guestAuth/app/rest/builds/buildType:(id:"

	fileHandle, err := os.Create("table1.md")
	if err != nil {
		log.Fatalln(err)
	}
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	fmt.Fprintln(writer, "Build Order | Component | Property Name | Current Version")
	fmt.Fprintln(writer, "------------|-----------|---------------|------------------")
	writer.Flush()

	for i := range mobj.Components {
		log.Println(">>>loop :" + strconv.Itoa(i))

		urlstr4 := urlstr3 + mobj.Components[i].BuildTypeId + ")"
		log.Println(urlstr3)

		req, err = http.NewRequest("GET", urlstr4, nil)
		if err != nil {
			log.Fatalln(err)
		}

		req.Header.Add("Accept", "application/json")

		resp, err := netClient.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		// _, err = io.Copy(os.Stdout, resp.Body)
		// fmt.Printf("\nError: %v\n", err)
		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&m)
		if (err != nil) && (err != io.EOF) {
			log.Println(err)
			panic(err)
			//log.Fatalln(err)
		}
		log.Println(">>" + strconv.Itoa(i))
		//log.Println(mm)
		log.Println(m)
		log.Println(m.Number)
		//log.Println(m[i])

		stringprint := fmt.Sprintf("%v | %v | %v | [%v](%v)", mobj.Components[i].BuildOrder, mobj.Components[i].Name, mobj.Components[i].PropertyName, (m.Number), m.WebUrl)
		fmt.Fprintln(writer, stringprint)
		writer.Flush()

	}

	/*
		urlstr3 = urlstr3 + mobj.Components[0].BuildTypeId + ")"
		log.Println("urlstr3")
		log.Println(urlstr3)
		req, err = http.NewRequest("GET", urlstr3, nil)
		if err != nil {
			log.Fatalln(err)
		}
		req.Header.Add("Accept", "application/json")

		resp, err := netClient.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		//	_, err = io.Copy(os.Stdout, resp.Body)
		//fmt.Printf("Error: %v\n", err)

		decoder := json.NewDecoder(resp.Body)
		//m := Fo{}
		err = decoder.Decode(&m)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(m)
		log.Println(m.Number)
	*/

	//defer decoder.Close()

	/***************************/
	//write the file with latest verion numbers

	/* file will look like below
	   Build Order | Component | Property Name | Current Version
	   ----------------|-----------------|----------------------|-----------------------
	   1 | server-parent | None - this is the Parent Pom and must be directly referenced | [3.0.15](http://teamcity.cvs.ula.comcast.net:8111/viewLog.html?buildId=4062409&buildTypeId=CptServers_CptServersLegac_ReleaseBuildTvworksServerParentPom)
	   2 | external | cvsp.external.version |  3.0.20
	   3 | ace-lib | cvsp.ace-lib.version | 3.0.6
	   4 | c-libs | cvsp.c-libs.version | 3.0.14
	*/
	/***************************/
	//////////
	/*	fileHandle, err := os.Create("table1.md")
		if err != nil {
			log.Fatalln(err)
		}
		writer := bufio.NewWriter(fileHandle)
		defer fileHandle.Close()

		fmt.Fprintln(writer, "Build Order | Component | Property Name | Current Version")
		fmt.Fprintln(writer, "------------|-----------|---------------|------------------")
		writer.Flush()
		stringprint := fmt.Sprintf("%v | %v | %v | [%v](%v)", mobj.Components[0].Build_order, mobj.Components[0].Name, mobj.Components[0].Property_name, (m.Number), m.WebUrl)
		fmt.Fprintln(writer, stringprint)
		writer.Flush()
		stringprint = fmt.Sprintf("%v | %v | %v | [%v](%v)", mobj.Components[0].Build_order, mobj.Components[0].Name, mobj.Components[0].Property_name, (m.Number), m.WebUrl)
		fmt.Fprintln(writer, stringprint)
		writer.Flush()
	*/

}
