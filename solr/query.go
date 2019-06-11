package solr

import (
	"fmt"
	"net/url"
)

type Query struct {
	params *url.Values
	si     *SolrInterface
}

func NewQuery(si *SolrInterface) *Query {
	q := new(Query)
	q.params = &url.Values{}
	q.si = si
	return q
}

func (q *Query) AddParam(k, v string) *Query {
	q.params.Add(k, v)
	return q
}

func (q *Query) RemoveParam(k string) *Query {
	q.params.Del(k)
	return q
}

func (q *Query) GetParam(k string) string {
	return q.params.Get(k)
}

func (q *Query) SetParam(k, v string) *Query {
	q.params.Set(k, v)
	return q
}

// q parameter http://wiki.apache.org/solr/CommonQueryParameters
// Example: id:100
func (qq *Query) Q(q string) *Query {
	qq.params.Add("q", q)
	return qq
}

// sort parameter http://wiki.apache.org/solr/CommonQueryParameters
// Example: geodist() asc
func (q *Query) Sort(sort string) *Query {
	q.params.Add("sort", sort)
	return q
}

// fq (Filter Query) http://wiki.apache.org/solr/CommonQueryParameters
// Example: popularity:[10 TO *]
func (q *Query) FilterQuery(fq string) *Query {
	q.params.Add("fq", fq)
	return q
}

// fl (Field List ) parameter http://wiki.apache.org/solr/CommonQueryParameters
// Example: id,name,decsription
func (q *Query) FieldList(fl string) *Query {
	q.params.Add("fl", fl)
	return q
}

// f (Facet) https://cwiki.apache.org/confluence/display/solr/Faceting#Faceting-Thefacet.fieldParameter
// Example: category
func (q *Query) AddFacet(f string) *Query {
	q.params.Set("facet", "true")
	q.params.Add("facet.field", f)
	return q
}

// fq (FacetQuery) https://wiki.apache.org/solr/SimpleFacetParameters#facet.query_:_Arbitrary_Query_Faceting
// Example: price:[* TO 500]
func (q *Query) AddFacetQuery(fq string) *Query {
	q.params.Set("facet", "true")
	q.params.Add("facet.query", fq)
	return q
}

// mc (Facet min count) https://cwiki.apache.org/confluence/display/solr/Faceting#Faceting-Thefacet.mincountParameter
// Example: 5
func (q *Query) SetFacetMinCount(mc int) *Query {
	q.params.Set("facet.mincount", fmt.Sprintf("%d", mc))
	return q
}

// f (Facet) https://wiki.apache.org/solr/SimpleFacetParameters#facet.pivot
// Example: category
func (q *Query) AddFacetPivot(f string) {
	q.params.Add("facet.pivot", f)
}

// mc (Facet pivot min count) https://wiki.apache.org/solr/SimpleFacetParameters#facet.pivot
// Example: 5
func (q *Query) SetFacetPivotMinCount(mc int) *Query {
	q.params.Set("facet.pivot.mincount", fmt.Sprintf("%d", mc))
	return q
}

// jf (Json facet) https://cwiki.apache.org/confluence/display/solr/JSON+Request+API#JSONRequestAPI-FacetExample
// Example: {avg_price:"avg(price)"}
func (q *Query) AddJsonFacet(jf string) *Query {
	q.params.Add("json.facet", jf)
	return q
}

// geofilt - The distance filter http://wiki.apache.org/solr/SpatialSearch
// Output example: fq={!geofilt pt=45.15,-93.85 sfield=store d=5}
func (q *Query) Geofilt(latitude, longitude float64, sfield string, distance float64) *Query {
	q.params.Add("fq", fmt.Sprintf("{!geofilt pt=%#v,%#v sfield=%s d=%#v}", latitude, longitude, sfield, distance))
	return q
}

// defType http://wiki.apache.org/solr/CommonQueryParameters
// Example: dismax
func (q *Query) DefType(defType string) *Query {
	q.params.Add("defType", defType)
	return q
}

// bf (Boost Functions) parameter http://wiki.apache.org/solr/DisMaxQParserPlugin
// Example: ord(popularity)^0.5 recip(rord(price),1,1000,1000)^0.3
// Check this http://wiki.apache.org/solr/FunctionQuery for available functions
func (q *Query) BoostFunctions(bf string) *Query {
	q.params.Add("bf", bf)
	return q
}

// bq (Boost Query) parameter http://wiki.apache.org/solr/DisMaxQParserPlugin
func (q *Query) BoostQuery(bq string) *Query {
	q.params.Add("bq", bq)
	return q
}

// qf (Query Fields) parameter http://wiki.apache.org/solr/DisMaxQParserPlugin
// Example: features^20.0+text^0.3
func (q *Query) QueryFields(qf string) *Query {
	q.params.Add("qf", qf)
	return q
}

// Start sets start value which is the offset of result
func (q *Query) Start(start int) *Query {
	q.params.Set("start", fmt.Sprintf("%d", start))
	return q
}

// Rows sets value for rows which means set the limit for how many rows to return
func (q *Query) Rows(rows int) *Query {
	q.params.Set("rows", fmt.Sprintf("%d", rows))
	return q
}

// Search Return a new instace of Search
func (q *Query) Search() *Search {
	return NewSearch(q.si.conn, q)
}

func (q *Query) String() string {
	return q.params.Encode()
}
