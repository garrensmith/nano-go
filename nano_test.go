package nano

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "net/url"
)

var couchUrl = "http://garren:password@localhost:5984"

func TestConnectToDB(t *testing.T) {
  
  couch := Setup(couchUrl)

  out := couch.Version()
  assert.Equal(t, out["couchdb"], "Welcome")
}


func TestAllDatabases(t *testing.T) {
  couch := Setup(couchUrl)

  dbs := couch.Alldbs()

  assert.Equal(t, dbs[0], "_replicator")
}

func TestCreateDB(t *testing.T) {

  couch := Setup(couchUrl)

  resp, err := couch.CreateDb("newdb")
  defer couch.DestroyDb("newdb")

  assert.Nil(t, err)
  assert.True(t, resp["ok"])
}

func TestRemoveDB(t *testing.T) {
  couch := Setup(couchUrl)
  couch.CreateDb("newdbdelete")
  resp, err := couch.DestroyDb("newdbdelete")

  assert.Nil(t, err)
  assert.True(t, resp["ok"])
}

func TestUseDB(t *testing.T) {
  couch := Setup(couchUrl)
  db, err := couch.UseDb("nano-go-tests")

  assert.Nil(t, err)
  assert.Equal(t, db.name, "nano-go-tests")
}

func TestGetDoc(t *testing.T) {
  id := "16b57b197b487766dc2973e9806bd0c7"
  couch := Setup(couchUrl)
  couch.UseDb("nano-go-tests")

  doc, err := couch.Db.Get(id)
  assert.Nil(t, err)
  assert.Equal(t, doc["_id"],id)
}

type DocTest struct {
  Id string `json:"_id"`
  Rev string `json:"_rev"`
  Type string `json:"type"`
  UserName string `json:"user_name"`
}

func TestGetDocMarshalled(t *testing.T) {
  id := "16b57b197b487766dc2973e9806bd0c7"
  couch := Setup(couchUrl)
  couch.UseDb("nano-go-tests")

  var doc DocTest
  _, err := couch.Db.GetFor(id, &doc)

  assert.Nil(t, err)
  assert.Equal(t, doc.Id, id)
  assert.Equal(t, doc.Type, "ActivityItem")
  assert.Equal(t, doc.UserName, "Bill Pearl")
}

type ViewResponse struct {
  TotalRows int `json:"total_rows"`
  Rows []interface{} `json:"rows"`
}

func TestGetView(t *testing.T) {
  couch := Setup(couchUrl)
  couch.UseDb("nano-go-tests")

  params := url.Values{
                        "reduce": []string{"false"},
                      }

  var resp  ViewResponse
  err := couch.Db.View("ActivityItem", "testView", &params, &resp)
  assert.Nil(t, err)

  assert.Equal(t, len(resp.Rows), 3)
}

type ViewResponse2 struct {
  TotalRows int `json:"total_rows"`
  Rows []map[string]string `json:"rows"`
}

func TestGetViewJson(t *testing.T) {
  couch := Setup(couchUrl)
  couch.UseDb("nano-go-tests")
  
  params := url.Values{ 
    "reduce": []string{"false"},
    "include_docs": []string{"true"},
    "limit": []string{"20"},
   "startkey": []string{"[\"1\", \"0\"]"},
   "endkey": []string{"[\"1\", {}]"},
  }

  resp, err := couch.Db.ViewJson("ActivityItem", "testView", &params)
  assert.Nil(t, err)

  rows := resp.Get("rows")
  rowArray, _ := rows.Array()
  assert.Equal(t, len(rowArray), 3)
  id, _ := rows.GetIndex(0).Get("id").String()
  assert.Equal(t, id, "16b57b197b487766dc2973e9806bf62a")
}

func TestGetViewJsonError(t *testing.T) {
  couch := Setup(couchUrl)
  couch.UseDb("nano-go-tests")
  
  params := url.Values{ 
    "reduce": []string{"false"},
    "include_docs": []string{"true"},
    "limit": []string{"20"},
   "startkey": []string{"[\"1\", \"0\"]"},
   "endkey": []string{"[\"1\", {}]"},
   "descending": []string{"true"},
  }

  _, err := couch.Db.ViewJson("ActivityItem", "testView", &params)
  assert.Equal(t, err.Error(), "{\"error\":\"query_parse_error\",\"reason\":\"No rows can match your key range, reverse your start_key and end_key or set descending=false\"}\n")
}

func TestGenerateUUID(t *testing.T) {
  couch := Setup(couchUrl)
  couch.UseDb("nano-go-tests")
  
  uuid := couch.Uuids(3)
  assert.Equal(t, len(uuid), 3)
}
