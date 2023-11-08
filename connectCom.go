package pkg

import (
	"encoding/hex"
	"fmt"
	"github.com/tarm/serial" // библиотека для работы с последовательным портом
	"io"
	"log"
	"strconv"
	"time"
)

func ReadTemp() int {
	//Открываем COM-порт
	config := &serial.Config{
		Name:        "COM10", // Замените на нужный COM-порт
		Baud:        9600,    // Установите нужную скорость передачи данных
		ReadTimeout: time.Millisecond * 10,
		Size:        8,
		Parity:      serial.Parity('E'),
		StopBits:    1,
	}
	port, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	a := []byte{0x05, 0x04, 0x01, 0xF4, 0x00, 0x02, 0x30, 0x41}
	port.Write(a)
	time.Sleep(time.Millisecond * 50)
	// Читаем данные из порта
	buf := make([]byte, 10)
	_, err = port.Read(buf)
	if err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
	}

	// Конвертация строк
	first := strconv.FormatInt(int64(buf[5]), 16)
	second := strconv.FormatInt(int64(buf[6]), 16)
	answer, err := hex.DecodeString(string(first[1]) + string(second[0]))
	if err != nil {
		log.Fatal(err)
	}

	return int(answer[0])
}

func ReadFlow() string {
	//Открываем COM-порт
	config := &serial.Config{
		Name:        "COM10", // Замените на нужный COM-порт
		Baud:        9600,    // Установите нужную скорость передачи данных
		ReadTimeout: time.Millisecond * 10,
		Size:        8,
		Parity:      serial.Parity('E'),
		StopBits:    1,
	}
	port, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	a := []byte{0x05, 0x04, 0x02, 0x08, 0x00, 0x02, 0xF0, 0x35}
	port.Write(a)
	time.Sleep(time.Millisecond * 50)
	// Читаем данные из порта
	buf := make([]byte, 10)
	_, err = port.Read(buf)
	if err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
	}

	// Конвертация строк
	first := strconv.FormatInt(int64(buf[5]), 16)
	second := strconv.FormatInt(int64(buf[6]), 16)
	if buf[5] == 0 {
		return "Нет потока"
	}
	answer, err := hex.DecodeString(string(first[1]) + string(second[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(answer)
	return "Скорость потока" + string(answer[0])
}
