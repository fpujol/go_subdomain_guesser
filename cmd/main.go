package main

import (
	"fmt"

	"github.com/miekg/dns"
)

func main() {
	var msg dns.Msg
	domain := dns.Fqdn("adssl.com")

	msg.SetQuestion(domain, dns.TypeA)

	in, err := dns.Exchange(&msg, "8.8.8.8:53")

	if err!=nil {
		panic(err)
	}

	if len(in.Answer) <1 {
		fmt.Println("No records!")
		return
	}

	for _, answer := range in.Answer {
		if a,ok := answer.(*dns.A);ok {
			fmt.Print("Ip Found ---> ", a.A)
		}
	}
}
