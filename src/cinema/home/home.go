package home

import(
    "fmt"
    "net/http"
    "html/template"
    "encoding/json"
    "cinema/model"
    _ "database/sql"
    "github.com/acoshift/flash"
    "github.com/gorilla/sessions"
    "log"
)

type Flash struct{
    FlashMessage string
}

var sessionStore sessions.Store

var f = flash.New()

type Seat struct {
    Id                   int    `db:"id,omitempty"`
    SeatNumber           int32  `db:"seat_number"`
    Status               int32  `db:"status"`
    BookedBy             int64  `db:"booked_by"`
}

// index page action for view seats
func IndexAction(w http.ResponseWriter, r *http.Request) {
   
    seats, err :=  model.GetAllSeats()
    if err != nil {
        log.Printf("%s", err)
    }

    tmpl,_ := template.ParseFiles("./layout/layout.html", "./home/view/index.html")

    var fm Flash
    if f.Has("message") == true {
        s := fmt.Sprintf("%s", f.Get("message"))
        fm.FlashMessage = s
        f.Clear()
    }

    data := []interface{}{
        seats,
        fm,
    }

    tmpl.Execute(w, data)
}

// save booking for selected seats
func SaveBooking(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    id := r.FormValue("id")
    name := r.FormValue("name")
    email := r.FormValue("email")

    var seatNo []string
    _ = json.Unmarshal([]byte(id), &seatNo)
    
    res, err := model.BookSeat(name, email, seatNo)
    if err != nil {
        log.Printf("%s", err)
    }
    js, err := json.Marshal(res)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}

// get all seats booked by seat number
func GetBooking(w http.ResponseWriter, r *http.Request) {

    r.ParseForm()
    id := r.FormValue("id")

    res, err := model.GetBooking(id)

    js, err := json.Marshal(res)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}

// cancel selected seat booking by seat number and email id
func CancelBooking(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    seatIds := r.FormValue("ids")
    email := r.FormValue("email")
    
    var ids []string
    _ = json.Unmarshal([]byte(seatIds), &ids)

    res, err := model.CancelBooking(ids, email)

    js, err := json.Marshal(res)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)

 }
