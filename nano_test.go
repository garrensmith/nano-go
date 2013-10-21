package nano

import (
  "testing"
  "github.com/stretchr/testify/assert"
  //"fmt"
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

}
