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
		log.Println(err)
		panic(err)

	}
	req.Header.Add("Accept", "application/json")

	/****************/
	//Read table1 config from a json file
	/****************/
	pathtofile := "../config/table1config.json"
	file, e := ioutil.ReadFile(pathtofile)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	//fmt.Printf("%s\n", string(file))

	var mobj componentjson

	err = json.Unmarshal(file, &mobj)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	log.Println("first attempt")
	fmt.Printf("Results: %v\n\n\n", mobj)
	log.Println(mobj)
	//numberofcomp := len(mobj.Components)
	// fmt.Printf("Results: %v\n\n\n", mobj.Components)
	// log.Println(mobj.Components)

	/***************************/
	//run the rest url to get json file
	/***************************/
	var m Message1
	//mm := make([]Message1, numberofcomp)
	urlstr3 := "http://teamcity.cvs-a.ula.comcast.net:8111/guestAuth/app/rest/builds/buildType:(id:"

	fileHandle, err := os.Create("../table1.md")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	fmt.Fprintln(writer, "\nThe following table represents the build order (top to bottom), current versions, and"+
		"the property names in the poms that represent those versions. As you modify POMs, please make sure to use these property names."+
		"Note that components with the same build order number can be built in parallel.\nTo add a new component make sure you edit **table1config.json** file\n\n")

	fmt.Fprintln(writer, "Build Order | Component | Property Name | Current Version")
	fmt.Fprintln(writer, "------------|-----------|---------------|------------------")
	writer.Flush()

	for i := range mobj.Components {
		log.Println(">>>loop :" + strconv.Itoa(i))

		urlstr4 := urlstr3 + mobj.Components[i].BuildTypeId + ")"
		log.Println(urlstr4)

		req, err = http.NewRequest("GET", urlstr4, nil)
		if err != nil {
			log.Println(err)
			panic(err)
		}

		req.Header.Add("Accept", "application/json")

		resp, err := netClient.Do(req)
		if err != nil {
			log.Println(err)
			panic(err)
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

	/****************/
	//Read table2 config from a json file
	/****************/
	pathtofile = "../config/table2config.json"
	file2, e := ioutil.ReadFile(pathtofile)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	//fmt.Printf("%s\n", string(file2))

	var mobj2 componentjson

	err = json.Unmarshal(file2, &mobj2)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	log.Println("first attempt")
	fmt.Printf("Results: %v\n\n\n", mobj2)
	log.Println(mobj2)
	//numberofcomp := len(mobj2.Components)
	// fmt.Printf("Results: %v\n\n\n", mobj2.Components)
	// log.Println(mobj2.Components)

	/***************************/
	//run the rest url to get json file
	/***************************/
	var m2 Message1
	//mm := make([]Message1, numberofcomp)
	urlstr3 = "http://teamcity.cvs-a.ula.comcast.net:8111/guestAuth/app/rest/builds/buildType:(id:"

	fileHandle2, err := os.Create("../table2.md")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	writer = bufio.NewWriter(fileHandle2)
	defer fileHandle2.Close()

	fmt.Fprintln(writer, "\nAdditional properties for components that are NOARCH or testing only and may not necessarily need to be rebuilt for the project.\n"+
		"To add a new component make sure you edit **table2config.json** file\n\n")

	fmt.Fprintln(writer, "Build Order | Component | Property Name | Current Version")
	fmt.Fprintln(writer, "------------|-----------|---------------|------------------")
	writer.Flush()

	for i2 := range mobj2.Components {
		log.Println(">>>loop :" + strconv.Itoa(i2))

		urlstr4 := urlstr3 + mobj2.Components[i2].BuildTypeId + ")"
		log.Println(urlstr4)

		req, err = http.NewRequest("GET", urlstr4, nil)
		if err != nil {
			log.Println(err)
			panic(err)
		}

		req.Header.Add("Accept", "application/json")

		resp, err := netClient.Do(req)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		// _, err = io.Copy(os.Stdout, resp.Body)
		// fmt.Printf("\nError: %v\n", err)
		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&m2)
		if (err != nil) && (err != io.EOF) {
			log.Println(err)
			panic(err)
			//log.Fatalln(err)
		}
		log.Println(">>" + strconv.Itoa(i2))
		//log.Println(mm)
		log.Println(m2)
		log.Println(m2.Number)
		//log.Println(m[i])

		stringprint := fmt.Sprintf("%v | %v | %v | [%v](%v)", mobj2.Components[i2].BuildOrder, mobj2.Components[i2].Name, mobj2.Components[i2].PropertyName, (m2.Number), m2.WebUrl)
		fmt.Fprintln(writer, stringprint)
		writer.Flush()

	}

}
