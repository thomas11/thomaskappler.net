package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/radovskyb/watcher"
	"github.com/thomas11/blog11"
)

var siteConfPath = flag.String("siteConfPath", "blog11.yml", "Path to the site configuration file")
var serve = flag.Bool("serve", false, "Start a localhost:9999 server for the site")
var watch = flag.Bool("watch", false, "Keep running and re-render the site on changes to the input directory.")

const siteUrl = "http://www.thomaskappler.net/"

var conf = blog11.SiteConf{
	Author:                       "Thomas Kappler",
	AuthorUri:                    siteUrl,
	BaseUrl:                      siteUrl,
	SiteTitle:                    "Thomas Kappler's site. Mostly programming and books.",
	CategoriesOutDir:             "categories",
	WritingFileExtension:         ".text",
	WritingFileDateStampFormat:   "2006-01-02",
	ImgOutDir:                    "img",
	WritingDir:                   "../writing",
	OutDir:                       "../thomas11.github.com",
	TemplateDir:                  "tmpl",
	MaxArticlesOnIndex:           20,
	NumFreqCategories:            5,
	MinArticlesForFreqCategories: 2,
	MaxAgeForFreqCategories:      time.Hour * 24 * 365 * 2,
}

func main() {
	flag.Parse()

	renderSite(&conf)

	if *watch {
		go rerenderOnChange(&conf)
	}

	if *serve {
		serveSite(conf.OutDir)
	}
}

func renderSite(conf *blog11.SiteConf) {
	site, err := blog11.ReadSite(conf)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Writing site to " + conf.OutDir)
	err = site.RenderAll()
	if err != nil {
		log.Fatal(err)
	}
}

func serveSite(dir string) {
	port := ":9999"

	http.Handle("/", http.FileServer(http.Dir(dir)))
	log.Printf("Serving %v on %v.", dir, port)
	http.ListenAndServe(port, nil)
}

func rerenderOnChange(siteConf *blog11.SiteConf) {
	log.Println("Watching " + siteConf.WritingDir + " for changes...")

	watcher := watcher.New()
	watcher.SetMaxEvents(1)

	go func() {
		for {
			select {
			case _ = <-watcher.Event:
				renderSite(siteConf)
			case err := <-watcher.Error:
				log.Println(err)
			case <-watcher.Closed:
				return
			}
		}
	}()

	if err := watcher.AddRecursive(conf.WritingDir); err != nil {
		log.Fatalln(err)
	}

	if err := watcher.Start(time.Millisecond * 200); err != nil {
		log.Fatalln(err)
	}
}
