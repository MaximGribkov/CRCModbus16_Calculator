# CRCModbus16_Calculator
CRC 16 Modbus

Function usage example:

// device address Modbus
	address := byte(0x02)

// function code for reading holding registers
	functionCode := byte(0x03)

// register address to be read (1F4 в hex = 500 в dec)
	registerAddressHi := byte(0x01)
	registerAddressLo := byte(0xF4)

// number of registers to read (2)
	registerCountHi := byte(0x00)
	registerCountLo := byte(0x02)

// collect data package
	packet := []byte{address, functionCode, registerAddressHi, registerAddressLo, registerCountHi, registerCountLo}

// CRC-16
	crc := CalculateCRC(packet)
