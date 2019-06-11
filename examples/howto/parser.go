package main

import "github.com/trrtly/go-solr/solr"
import "fmt"

func main() {

	si, _ := solr.NewSolrInterface("http://localhost:8983/solr", "collection1")

	r, err := si.Query().Q("title:add sucess 1").
		Start(0).Rows(15).
		Search().Result(nil)

	if err != nil {
		fmt.Println("Error when querying solr:", err.Error())
		return
	}

	fmt.Println(r.Results.Docs)
}
