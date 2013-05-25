package main

import (
	"github.com/thomas11/blog11"
	"log"
)

const siteUrl = "http://www.thomaskappler.net/"

var conf = blog11.SiteConf{
	Author:                     "Thomas Kappler",
	AuthorUri:                  siteUrl,
	BaseUrl:                    siteUrl,
	SiteTitle:                  "Thomas Kappler's site. Mostly programming and books.",
	CategoriesOutDir:           "categories",
	WritingFileExtension:       ".text",
	WritingFileDateStampFormat: "2006-01-02",
	ImgOutDir:                  "img",
	WritingDir:                 "../writing-markdown",
	MaxArticlesOnIndex:         10,
	OutDir:                     "thomas11.github.com",
	TemplateDir:                "tmpl",
}

func main() {
	site, err := blog11.ReadSite(&conf)
	if err != nil {
		log.Fatal(err)
	}

	err = site.RenderAll()
	if err != nil {
		log.Fatal(err)
	}
}
