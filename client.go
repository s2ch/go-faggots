package main

import (
        // "log"
        // "net"
        "crypto/tls"
        "fmt"

        irc "github.com/thoj/go-ircevent"
        // "crypto/x509"
        // "gopkg.in/irc.v3"
)

const channel = "#carmack-test"
const serverssl = "chat.freenode.net:6697"

func main() {
        irccon := irc.IRC("urp", "brs")
        irccon.VerboseCallbackHandler = true
        irccon.Debug = true
        irccon.UseTLS = true
        irccon.Password = ""
        irccon.TLSConfig = &tls.Config{InsecureSkipVerify: true}
        irccon.AddCallback("001", func(e *irc.Event) { irccon.Join(channel) })
        irccon.AddCallback("366", func(e *irc.Event) {})
        err := irccon.Connect(serverssl)

        if err != nil {
                fmt.Printf("Err %s", err)
                return
        }

        irccon.AddCallback("PRIVMSG", func(e *irc.Event) {
                go func(event *irc.Event) {
                        irccon.Privmsg("#carmack-test", e.Message())
                }(e)
        })

        irccon.Loop()
}
