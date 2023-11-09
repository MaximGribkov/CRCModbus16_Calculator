package pkg

import (
	"fmt"
	"github.com/goburrow/modbus"
	"log"
)

func ModbusTCP() {
	// Создаем клиент Modbus TCP
	client := modbus.TCPClient("192.168.127.254:502") // Замените IP-адрес на адрес вашего устройства

	// Чтение значения входа
	result, err := client.ReadCoils(0, 1) // Читаем один вход, начиная с адреса 0
	if err != nil {
		log.Fatalf("error in read %s", err)
	}
	fmt.Println("Значение входа:", result[0])

	// Запись значения входа	0xFF00 (ON)		0x0000 (OFF)
	res, err := client.WriteSingleCoil(0, 0xFF00) // Записываем значение True во вход с адресом 0
	if err != nil {
		log.Fatalf("error in write %s", err)
	}

	fmt.Printf("Значение входа %v", res)
}
