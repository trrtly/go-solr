package main

import "github.com/trrtly/go-solr/solr"
import "fmt"

type result struct {
	Response struct {
		NumFound int `json:"numFound"`
		Start    int `json:"start"`
		Docs     []struct {
			IndusName string `json:"indus_name"`
		} `json:"docs"`
	} `json:"response"`
}

func main() {

	si, _ := solr.NewSolrInterface("http://localhost:8983/solr", "collection1")

	var data result

	si.Query().
		FieldList("*").
		Q("*").
		FilterQuery("NOT indus_type:2").
		FilterQuery("NOT g_id:0").
		FilterQuery("indus_pid:0").
		Sort("task_num desc").
		Start(0).
		Rows(1).
		SetResult(&data).
		Search().Result(nil)

	fmt.Println(data.Response.Docs[0].IndusName)
}
