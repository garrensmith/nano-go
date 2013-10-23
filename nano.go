package nano

import (
  "net/http"
  "encoding/json"
  "io/ioutil"
)

type Couch struct {
  url string
  db Database
}

type Database struct {
  name string
  url string
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

func (couch *Couch) CreateDb(name string) (reply map[string]bool, err error ) {
  body, err := put(couch.url + "/" + name)

  if err != nil {
    // handle error
    return
  }

  json.Unmarshal(body, &reply)

  return
}

func (c *Couch) UseDb(name string) (db Database, err error ) {
  err = nil
  db.name = name
  db.url = c.url + "/" + name
  c.db = db
  return
}

func (db *Couch) DestroyDb(name string) (reply map[string]bool, err error ) {
  body, err := delete(db.url + "/" + name)

  if err != nil {
    // handle error
    return
  }
  json.Unmarshal(body, &reply)

  return
}

func (db *Couch) Alldbs() (dbs []string) {
  body, _ := get(db.url + "/_all_dbs")
  json.Unmarshal(body, &dbs)
  return
}

func (db *Database) get(name string) (doc map[string]interface{}, err error) {
  body, err := get(db.url + "/" + name)

  if err != nil {
    return
  }

  json.Unmarshal(body, &doc)
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

func put(url string) (body []byte, err error) {
  body, err = request(url, "PUT")
  return
}

func delete(url string) (body []byte, err error) {
  body, err = request(url, "DELETE")
  return
}

func request(url string, method string) (body []byte, err error) {
  client := &http.Client{}
  request, err := http.NewRequest(method, url, nil)
  resp, err := client.Do(request)
  
  if err != nil {
    return
  }

  defer resp.Body.Close()
  body, err = ioutil.ReadAll(resp.Body)

  return
}
