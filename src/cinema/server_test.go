package main

import (
    "strings"
    "testing"
    "net/http"
    "io/ioutil"
    "cinema/home"
    "cinema/config"
    "encoding/json"
    "net/http/httptest"
)

type Response struct {
    Success string    `json:"success"`
    Mesaage string    `json:"message"`
}


var URL = config.DomainName + ":" + config.AppPort

func TestSaveBooking(t *testing.T) {
    //Create a request with data
    req, _ := http.NewRequest("POST", URL+"/save-booking",
        strings.NewReader("name=Bob&email=bob987@mailinator.com&id=[\"20\"]"))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

    //Create a response to store the data requested
    w := httptest.NewRecorder()
    //Call the function
    home.SaveBooking(w,req)
    //Read the response body
    data, _ := ioutil.ReadAll(w.Body)

    res := Response{}
    json.Unmarshal(data, &res)

    if res.Success != "ok" {
        t.Errorf("Fail, expected success and got %v", res.Success)
    }
}

func TestCancelBooking(t *testing.T) {
    //Create a request with data
    req, _ := http.NewRequest("POST", URL+"/cancel-booking",
        strings.NewReader("ids=[\"20\"]&email=bob987@mailinator.com"))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

    //Create a response to store the data requested
    w := httptest.NewRecorder()
    //Call the function
    home.CancelBooking(w,req)
    //Read the response body
    data, _ := ioutil.ReadAll(w.Body)

    res := Response{}
    json.Unmarshal(data, &res)

    if res.Success != "ok" {
        t.Errorf("Fail, expected success and got %v", res.Success)
    }
}
