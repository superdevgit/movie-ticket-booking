package model

import (
    "log"
    "errors"
    "cinema/mail"
    "database/sql"
    "cinema/config"
)

type Seat struct {
    Id           int            `db:"id,omitempty"`
    SeatNumber   int32          `db:"seat_number"`
    Status       int32          `db:"status"`
    BookedBy     sql.NullInt64  `db:"booked_by"`
}

type Response struct {
    Success string `json:"success"`
    Message string `json:"message"`
}

type SeatResponse struct {
    Success string `json:"success"`
    Message string `json:"message"`
    Seats    []int  `json:"seats"`
}

type CancelResponse struct {
    Success string    `json:"success"`
    Message string    `json:"message"`
    Seats   []string  `json:"seats"`
}

// get all seats
func GetAllSeats() ([]Seat, error){
    db, err := config.ConnectDB()
    if err != nil {
        log.Printf("%s", err)
    }
    defer db.Close()

    var seats []Seat
    err = db.Select(&seats, "SELECT * FROM seats")
    if err != nil {
        log.Printf("%s", err)
    }
    return seats, nil
}

// book selected seats
func BookSeat(name string, email string, seats []string) (Response, error) {

    var response Response

    db, err := config.ConnectDB()
    if err != nil {
        log.Printf("%s", err)
        return response, err
    }
    defer db.Close()

    if name == "" || email == "" || len(seats) == 0 {
        err := errors.New("Enter all details correct!")
        response.Success = "fail"
        response.Message = "Enter all details correct!"
        return response, err
    }

    tx, _ := db.Begin()

    var userId []int64
    var id int64

    err = db.Select(&userId, "SELECT id FROM users WHERE email = ?", email)
    if err != nil {
        log.Printf("%s", err)
    }

    if len(userId) > 0 {
        id = userId[0]
    } else {
        res, err := tx.Exec(`INSERT INTO users (name, email) VALUES (?,?)`, name, email)
        if err != nil {
            log.Printf("%s", err)
            return response, err
        }
        id, err = res.LastInsertId()
        if err != nil {
            log.Printf("%s", err)
            return response, err
        }
    }

    for _, s := range seats {
        rows, err := tx.Exec(`UPDATE seats SET status= ?, booked_by = ? WHERE seat_number = ? AND status = 0`, 1, id, s)
        if err != nil {
            log.Printf("%s", err)
            return response, err
        }

        num, _:= rows.RowsAffected()
        if num == 0 {
            tx.Rollback()
            err := errors.New("Seat already booked!")
            response.Success = "fail"
            response.Message = "Seat already booked!"
            return response, err
        } 
        
    }

    tx.Commit()

    //send booking confirmation mail
    go mail.SendBookingMail(name, email, seats)

    response.Success = "ok"
    response.Message = "Seat booked successfuly"
    return response, nil
    
}

// get all booking seats
func GetBooking(seatId string) (SeatResponse, error){
    var response SeatResponse

    db, err := config.ConnectDB()
    if err != nil {
        log.Printf("%s", err)
        return response, err
    }
    defer db.Close()
    
    var seats []int

    err = db.Select(&seats, `SELECT s.seat_number FROM seats INNER JOIN seats s ON s.booked_by = seats.booked_by WHERE seats.seat_number = ?`, seatId)
    if err != nil {
        log.Printf("%s", err)
        return response, err
    }
    
    response.Success = "ok"
    response.Seats = seats
    
    return response, nil
    
}

// cancel booking
func CancelBooking(seatIds []string, email string) (CancelResponse, error){
    var response CancelResponse
    
    db, err := config.ConnectDB()
    if err != nil {
        log.Printf("%s", err)
        return response, err
    }
    
    if len(seatIds) == 0 {
        err := errors.New("Could not cancel booking!")
        response.Success = "fail"
        response.Message = "Seat not selected!"
        return response, err
    }
    defer db.Close()

    tx, _ := db.Begin()

    for _, s := range seatIds {
        rows, err := tx.Exec(`UPDATE seats SET status = 0, booked_by = ? WHERE seat_number = ? AND booked_by IN (SELECT id FROM users WHERE email = ?)`,nil, s, email)
        if err != nil {
            log.Printf("%s", err)
            response.Success = "fail"
            return response, err
        }
        num, _:= rows.RowsAffected()
        if num == 0 {
            tx.Rollback()
            err := errors.New("Could not cancel booking!")
            response.Success = "fail"
            response.Message = "Email not matched with this booking!"
            return response, err
        }
    }

    tx.Commit()

    response.Success = "ok"
    response.Seats = seatIds
    response.Message = "Booking cancelled successfuly"

    return response, nil
}
