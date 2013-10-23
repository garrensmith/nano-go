package nano

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "fmt"
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

  doc, err := couch.db.get("19836cb7b7776aa4ebc590492e1e9543")
  assert.Nil(t, err)
  assert.Equal(t, doc["_id"], "19836cb7b7776aa4ebc590492e1e9543")
}
