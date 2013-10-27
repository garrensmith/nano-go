package nano

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "fmt"
  "net/url"
)

func TestConnectToDB(t *testing.T) {

  couch := Setup("http://localhost:5984")

  out := couch.Version()
  assert.Equal(t, out["couchdb"], "Welcome")
}


func TestAllDatabases(t *testing.T) {
  couch := Setup("http://localhost:5984")

  dbs := couch.Alldbs()

  assert.Equal(t, dbs[0], "_replicator")
}

func TestCreateDB(t *testing.T) {

  couch := Setup("http://garren:password@localhost:5984")

  resp, err := couch.CreateDb("newdb")
  defer couch.DestroyDb("newdb")

  fmt.Println(resp)
  assert.Nil(t, err)
  assert.True(t, resp["ok"])
}

func TestRemoveDB(t *testing.T) {
  couch := Setup("http://garren:password@localhost:5984")
  couch.CreateDb("newdbdelete")
  resp, err := couch.DestroyDb("newdbdelete")
  fmt.Println(resp["reason"])

  assert.Nil(t, err)
  assert.True(t, resp["ok"])
}

func TestUseDB(t *testing.T) {
  couch := Setup("http://garren:password@localhost:5984")
  db, err := couch.UseDb("avol10")

  assert.Nil(t, err)
  assert.Equal(t, db.name, "avol10")
}

func TestGetDoc(t *testing.T) {
  couch := Setup("http://garren:password@localhost:5984")
  couch.UseDb("avol10")

  doc, err := couch.Db.Get("19836cb7b7776aa4ebc590492e1e9543")
  assert.Nil(t, err)
  assert.Equal(t, doc["_id"], "19836cb7b7776aa4ebc590492e1e9543")
}

type DocTest struct {
  Id string `json:"_id"`
  Rev string `json:"_rev"`
  Type string `json:"type"`
  Hello string `json:"hello"`
}

func TestGetDocMarshalled(t *testing.T) {
  couch := Setup("http://garren:password@localhost:5984")
  couch.UseDb("avol10")

  var doc DocTest
  _, err := couch.Db.GetFor("19836cb7b7776aa4ebc590492e1e9543", &doc)

  assert.Nil(t, err)
  assert.Equal(t, doc.Id, "19836cb7b7776aa4ebc590492e1e9543")
  assert.Equal(t, doc.Type, "cool-doc")
  assert.Equal(t, doc.Hello, "world")
}

type ViewResponse struct {
  TotalRows int `json:"total_rows"`
  Rows []interface{} `json:"rows"`
}

func TestGetView(t *testing.T) {
  couch := Setup("http://garren:password@localhost:5984")
  couch.UseDb("avol10")

  params := url.Values{
                        "reduce": []string{"false"},
                        "include_docs": []string{"true"},
                      }

  var resp  ViewResponse
  err := couch.Db.View("ddoc", "all", &params, &resp)
  assert.Nil(t, err)

  assert.Equal(t, len(resp.Rows), 2)
}

type ViewResponse2 struct {
  TotalRows int `json:"total_rows"`
  Rows []map[string]string `json:"rows"`
}

func TestGetView2(t *testing.T) {
  couch := Setup("http://garren:password@localhost:5984")
  couch.UseDb("tickle_development")
  
  params := url.Values{ 
    "reduce": []string{"false"},
    "include_docs": []string{"true"},
    "limit": []string{"20"},
   "startkey": []string{"[\"1\", \"0\"]"},
   "endkey": []string{"[\"1\", {}]"},
  }

  resp, err := couch.Db.View2("ActivityItem", "by_organisation_id_and_updated_at", &params)
  rows, err := resp.Get("rows").Array()
  assert.Nil(t, err)
  assert.Equal(t, len(rows), 20)
}
