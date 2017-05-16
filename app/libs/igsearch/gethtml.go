package igsearch

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

// GetHTML ...
func GetHTML(urlstr, ua string) (string, error) {
	//log.Print("getHTML")

	//NOTE: http://stackoverflow.com/questions/13130341/reading-gzipped-http-response-in-go

	if ua == "" {
		ua = "Googlebot"
	}
	req, _ := http.NewRequest("GET", urlstr, nil)
	// Set User-Agent to Googlebot
	req.Header.Set("User-Agent", ua)

	//gzip
	req.Header.Add("Accept-Encoding", "gzip")
	//req.Header.Add("Accept-Encoding", "gzip, deflate")

	// dump req header
	dump, _ := httputil.DumpRequestOut(req, true)
	fmt.Printf("%s", dump)
	// New Client
	tr := &http.Transport{
		//MaxIdleConns:       10,
		//IdleConnTimeout:    30 * time.Second,
		DisableCompression: false,
	}

	cl := &http.Client{
		Transport: tr,
	}
	// Send request
	resp, err := cl.Do(req)
	if err != nil {
		return "", err
	}
	//log.Print("defer resp.Body.Close()")
	defer resp.Body.Close()

	// dump resp header
	//dumpResp, _ := httputil.DumpResponse(resp, true)
	//fmt.Printf("%s", dumpResp)

	log.Printf("resp.Content-Length:%s", resp.Header.Get("Content-Length"))
	log.Printf("resp.Content-Encoding:%s", resp.Header.Get("Content-Encoding"))
	/*
		//log.Print("b, err := ioutil.ReadAll(resp.Body)")
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		//return string(b), nil
		if resp.Header.Get("Content-Encoding") == "gzip" {
			bb := bytes.NewBuffer(b)
			zipread, _ := gzip.NewReader(bb)

			defer zipread.Close()
			reader := bufio.NewReader(zipread)

			ret := ""
			var part []byte
			for {
				if part, _, err = reader.ReadLine(); err != nil {
					break
				}

				ret += string(part)

			}
			return ret, nil
		}
		return string(b), nil
	*/

	if resp.Header.Get("Content-Encoding") == "gzip" {

		zipread, e := gzip.NewReader(resp.Body)
		if e != nil {
			log.Print(e)
			return "", e
		}
		defer zipread.Close()

		reader := bufio.NewReader(zipread)
		var part []byte
		ret := ""

		for {
			if part, _, err = reader.ReadLine(); err != nil {
				break
			}
			ret += string(part)
		}
		log.Printf("unzipped resp body size = %d", len(ret))
		return ret, nil
	}

	//log.Print("b, err := ioutil.ReadAll(resp.Body)")
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	log.Printf("resp body size = %d", len(string(b)))
	return string(b), nil

}
