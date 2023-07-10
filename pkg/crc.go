package pkg

// Функция принимает на вход байтовый массив и возвращает айтовый массив - расчитанную контрольную сумму для протокола Modbus.
//
// Алгоритм расчета CRC контрольной суммы для протокола Modbus основан на полиномиальном сдвиге. Коэффициент полинома равен 0xA001.
// Для каждого байта входных данных выполняется операция XOR с текущим значением crc, затем происходит циклический сдвиг на 1 бит и проверка младшего бита.
// Если он равен 1, то происходит операция XOR с коэффициентом полинома. Алгоритм завершается после прохода всех байтов входных данных.

func CalculateCRC(data []byte) []byte {
	var crc uint16 = 0xFFFF

	for _, b := range data {
		crc ^= uint16(b)

		for i := 0; i < 8; i++ {
			if (crc & 0x0001) != 0 {
				crc >>= 1
				crc ^= 0xA001
			} else {
				crc >>= 1
			}
		}
	}
	return []byte{byte(crc & 0xFF), byte(crc >> 8)}
}