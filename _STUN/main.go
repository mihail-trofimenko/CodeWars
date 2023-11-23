package main

import (
	"fmt"

	"gortc.io/stun"
)

func main() {
	// Создаём соединение с сервером ОГЛУШЕНИЯ
	c, err := stun.Dial("udp", "stun.l.google.com:19302")
	if err != nil {
		panic(err)
	}
	// Создаём BindingRequest с рандомным TransactionID
	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)
	// Посылаем запрос на сервер ОГЛУШЕНИЯ, ожидаем ответного "message"
	if err := c.Do(message, func(res stun.Event) {
		if res.Error != nil {
			panic(res.Error)
		}
		// Декодируем аттрибут XOR-MAPPED-ADDRESS из "message"
		var xorAddr stun.XORMappedAddress
		if err := xorAddr.GetFrom(res.Message); err != nil {
			panic(err)
		}
		fmt.Printf("Мой IP:port = %s:%d", xorAddr.IP, xorAddr.Port)
	}); err != nil {
		panic(err)
	}
}
