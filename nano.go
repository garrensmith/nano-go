package nano

import (
  "net/http"
  "encoding/json"
  "io/ioutil"
  "net/url"
  "errors"
  "github.com/bitly/go-simplejson"
  "fmt"
)

type Couch struct {
  url string
  Db Database
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
  c.Db = db
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

func (c *Couch) Uuids(count int) (uuids []string) {
  var bodyJson map[string][]string

  params := url.Values{ "count": []string {fmt.Sprintf("%d",count)} }
  uuidUrl := url.URL{ Path: c.url + "/_uuids"}
  uuidUrl.RawQuery = params.Encode()
  fmt.Println(uuidUrl.String())
  body, _ := get(uuidUrl.String())
  fmt.Println("uu %v", string(body))
  json.Unmarshal(body, &bodyJson)
  uuids = bodyJson["uuids"]
  return
}

func (db *Database) GetFor(name string, doc interface{}) (interface{}, error) {
  body, err := get(db.url + "/" + name)

  if err != nil {
    return nil, err
  }

  json.Unmarshal(body, &doc)
  return doc, nil
}

func (db *Database) Get(name string) (doc map[string]interface{}, err error) {
  body, err := get(db.url + "/" + name)

  if err != nil {
    return
  }

  json.Unmarshal(body, &doc)
  return
}

func (db *Database) View(ddoc string, view string, params *url.Values, response interface{}) (err error) {
  viewUrl := url.URL{ Path: db.url + "/_design/" + ddoc + "/_view/" + view}

  if params != nil {
    viewUrl.RawQuery = params.Encode()
  }

  body, err := get(viewUrl.String())

  if err != nil {
    return
  }
  err = json.Unmarshal(body, &response)
  return
}

func (db *Database) ViewJson(ddoc string, view string, params *url.Values) (resp *simplejson.Json, err error) {
  viewUrl := url.URL{ Path: db.url + "/_design/" + ddoc + "/_view/" + view}

  if params != nil {
    viewUrl.RawQuery = params.Encode()
  }

  body, err := get(viewUrl.String())

  if err != nil {
    return
  }

  resp, err = simplejson.NewJson(body)
  fmt.Println("")
  if _,ok := resp.CheckGet("error"); ok {
    err = errors.New(string(body))
  }

  return
}


func get(url string) (body []byte, err error) {
  resp, err := http.Get(url)
  if err != nil {
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
