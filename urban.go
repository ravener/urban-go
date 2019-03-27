// Urban Go
// A Golang package to search urban dictionary.
package urban

import (
  "time"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "strconv"
  "fmt"
  "errors"
)

// The version of the package.
const Version = "0.0.1"

// Entry represents an entry in the dictionary.
type Entry struct {
  Definition string `json:"definition"`
  Example string `json:"example"`
  Author string `json:"author"`
  Word string `json:"word"`
  DefID int `json:"defid"`
  PermaLink string `json:"permalink"`
  SoundURLs []string `json:"sound_urls"`
  ThumbsUp int `json:"thumbs_up"`
  ThumbsDown int `json:"thumbs_down"`
  CurrentVote string `json:"current_vote"`
  WrittenOn time.Time `json:"written_on"`
}

type response struct {
  List []*Entry `json:"list"`
}

func get(query string) ([]*Entry, error) {
  req, err := http.NewRequest("GET", "https://api.urbandictionary.com/v0/define" + query, nil)
  if err != nil {
    return nil, err
  }
  req.Header.Set("User-Agent", "Urban Go (https://github.com/pollen5/urban-go) v" + Version)
  res, err := http.DefaultClient.Do(req)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    return nil, errors.New(fmt.Sprintf("err: non 200 status: %d", res.StatusCode))
  }
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return nil, err
  }
  var data response
  err = json.Unmarshal(body, &data)
  if err != nil {
    return nil, err
  }
  return data.List, nil
}

// Define gets the definitions for a term.
func Define(term string) ([]*Entry, error) {
  return get("?term=" + term)
}

// DefineByID returns definitions for a given defid
func DefineByID(defid int) ([]*Entry, error) {
  return get("?defid=" + strconv.Itoa(defid))
}
