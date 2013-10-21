package nano

import (
  "net/http"
  "encoding/json"
  "io/ioutil"
)

type Couch struct {
  url string
}

type Database struct {
  name string
}

func Setup(url string) (couch Couch) {
  couch.url = url
  return
}

func (couch *Couch) Version() (info map[string]string) {
  body, _ := get(couch.url)
  json.Unmarshal(body, &info)
  return
}

func (db *Couch) Alldbs() (dbs []string) {
  body, _ := get(db.url + "/_all_dbs")
  json.Unmarshal(body, &dbs)
  return
}

func get(url string) (body []byte, err error) {
  resp, err := http.Get(url)
  if err != nil {
    // handle error
    return
  }
  defer resp.Body.Close()
  body, err = ioutil.ReadAll(resp.Body)

  return
}
