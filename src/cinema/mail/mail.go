package mail

import (
    "log" 
    "net/smtp"
    "cinema/config"
)

// function for send mail
func sendMail(body string, to string) {
    from := config.MailFrom
    pass := config.SmtpPassword
    smtpHost := config.SmtpHost
    smtpPort := config.SmtpPort

    msg := "From: " + from + "\n" +
        "To: " + to + "\n" +
        "Subject: Ticket Confirmation" +
        body

    err := smtp.SendMail(smtpHost + ":" + smtpPort,
        smtp.PlainAuth("", from, pass, smtpHost),
        from, []string{to}, []byte(msg))

    if err != nil {
        log.Printf("smtp error: %s", err)
        return
    }
    
}

// send booking confirmation email to user
func SendBookingMail(name string, email string, seats []string) {
    var seatNum string

    for _, s := range seats {
        seatNum = seatNum + " " + s
    }

    body := "Hi " + name + "," + "\n" +
        "You have successfuly booked ticket" + "\n" +
        "your seat number is " + seatNum + "\n" +
        "Thanks for booking"

    sendMail(body, email)
}

