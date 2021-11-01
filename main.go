package main

import (
	"fmt"
	"log"
	"flag"
	"strings"
	"io"
	"os"
	"net/http"
	"path"
	"strconv"
	"sync"
	"github.com/PuerkitoBio/goquery"
	"github.com/schollz/progressbar"

)

func init() {
	flag.Usage = func() {
		h := []string{
			"Download Comic Books from zipcomic.com(concurrent with 2 goroutines by default)",
			"",
			"Usage:",
			"  zipcomic -u 'https://zipcomimic.com/comicbook_title'",
			"",
			"Options:",
			"  -u, -url <string> Specify a URL",
			"\tWith only -u, all comics will be downloaded in '/comicbook_title' in the current directory",
			"",
			"  -d, -dir <string> Change directory(should already be available)",
			"\tzipcomic -u 'https://zipcomimic.com/comicbook_title' -d '/directory'",
			"",
			"  -c, -concurrency <int> Number of Threads",
			"",
			"  -r, -range <string> Specify a Range",
			"\tzipcomic -u 'https://zipcomimic.com/comicbook_title' -r '5:10'\n",
		}

		fmt.Fprint(os.Stderr, strings.Join(h, "\n"))
	}
}

// check for errors
func checkErr(err error){
	if err != nil{
		log.Fatal(err)
	}
}

// struct for flags
type options struct{
	url string
	rangeN string
	dir string
	concurrency int
}

func mkDirNcd(opts options){

	newdir := path.Base(opts.url)
	err := os.Mkdir(newdir, 0755)
	checkErr(err)
	err = os.Chdir(newdir)
	checkErr(err)
}

//get all hrefs that starts with "/storage"
func getHref(opts options) []string{
	var listOfLinks []string
	doc, err := goquery.NewDocument(opts.url)
	checkErr(err)
	doc.Find("a[href]").Each(func(index int, item *goquery.Selection) {
        href, _ := item.Attr("href")
        if strings.HasPrefix(href, "/storage"){
        	//fmt.Println(href, item.Text()) //file and size
        	listOfLinks = append(listOfLinks, href)
        }
    })
    return listOfLinks
}

// if a range is provided
func calculateRange(opts options) (int, int){
	value := strings.Split(opts.rangeN, ":")
	from, _ := strconv.Atoi(value[0])
	to, _ := strconv.Atoi(value[1])
	if from > 0{
		from--
	}
	return from, to
}

func download(urlChan chan string, wg *sync.WaitGroup){ 

	for u := range urlChan {
		file := path.Base(u)

		req, _ := http.NewRequest("GET", u, nil)
		resp, _ := http.DefaultClient.Do(req)
		defer resp.Body.Close()

		f, _ := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()

		bar := progressbar.DefaultBytes(
			resp.ContentLength,
			file,
		)
		io.Copy(io.MultiWriter(f, bar), resp.Body)
	}
	wg.Done()
}

func sendUrlsOnChannel(urlChan chan string, urls []string, opts options){
	for _, v := range urls{
		file := fmt.Sprintf("%s%s", opts.url, v)
		urlChan <- file
	}
	close(urlChan)
}

func main(){

	var wg sync.WaitGroup

	opts := options{}
	flag.StringVar(&opts.url, "url", "", "")
	flag.StringVar(&opts.url, "u", "", "")
	flag.StringVar(&opts.rangeN, "range", "", "")
	flag.StringVar(&opts.rangeN, "r", "", "")
	flag.StringVar(&opts.dir, "dir", "", "")
	flag.StringVar(&opts.dir, "d", "", "")
	flag.IntVar(&opts.concurrency, "concurrency", 2, "number of goroutines")
	flag.IntVar(&opts.concurrency, "c", 2, "number of goroutines")
	flag.Parse()

	// get all urls that contains "/storage"
	urls := getHref(opts)

	// if a range is provided
	if opts.rangeN != ""{
		from, to := calculateRange(opts)
		urls = urls[from:to]
	}

	// if a directory is provided, change/cd into it
	if opts.dir != ""{
		err := os.Chdir(opts.dir)
		checkErr(err)
	}

	// create comicbook directory
	mkDirNcd(opts)

	// send all urls over the channel
	urlChan := make(chan string)
	go sendUrlsOnChannel(urlChan, urls, opts)

	// spawn N of goroutines
	for i := 0; i < opts.concurrency; i++{
		wg.Add(1)
		go download(urlChan, &wg)
	}

	wg.Wait()
}
