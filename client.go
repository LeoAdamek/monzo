package monzo

import (
    "net/http"
    "net/url"
    "log"
    "time"
    "os"
    "encoding/json"
    "io/ioutil"
)


const baseURLstr = "https://api.monzo.com/"
var baseURL *url.URL


func init() {
    baseURL, _ = url.Parse(baseURLstr)
}

// Client represents a Monzo client instance
type Client struct {
    
    HTTP *http.Client
    Log *log.Logger
    
    token string
}

// Pagination is used to control pagination of requests.
type Pagination struct {
    Limit int `json:"limit"`
    Since *time.Time `json:"since"`
    Before *time.Time `json:"before"`
}

// New creates a new Monzo client
func New(token string) *Client {
    
    return &Client{
        HTTP: &http.Client{Timeout: 2 * time.Second},
        Log: log.New(os.Stdout, "monzo ", log.LstdFlags),
        token: token,
    }
}

func (c Client) do(req *http.Request) (*http.Response, error) {
    
    if req.Header == nil {
        req.Header = make(http.Header)
    }
    
    req.Header.Set("User-Agent", "Monzo Go +https://github.com/LeoAdamek/monzo")
    req.Header.Set("Authorization", "Bearer " + c.token)
    
    st := time.Now()
    
    res, err := c.HTTP.Do(req)
    
    dt := time.Now().Sub(st)
    
    if err != nil {
        c.Log.Println("Error sending request: ", err)
        return res, err
    }
    
    c.Log.Println(req.Method, req.URL.String(), res.StatusCode, res.ContentLength, dt)
    return res, err
}

func (c Client) json(req *http.Request, into interface{}) error {
    
    if req.Header == nil {
        req.Header = make(http.Header)
        req.Header.Set("Content-Type", "application/json;charset=utf-8")
    }
    
    req.Header.Set("Accept", "application/json")
    req.Header.Add("Accept", "text/*;q=0.4")
    req.Header.Add("Accept", "*/*;q=0.2")

    res, err := c.do(req)
    
    if err != nil {
        return err
    }
    
    data, err := ioutil.ReadAll(res.Body)
    
    if err != nil {
        return err
    }
    
    err = json.Unmarshal(data, into)
    
    if err != nil {
        c.Log.Println("Unable to parse response:", err)
        c.Log.Println("Response data:", string(data))
    }
    
    return err
}
