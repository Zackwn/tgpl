package main

type Bytes uint64

const (
	byte Bytes = 8
	KB         = 1024 * byte
	MB         = 1024 * KB
	GB         = 1024 * MB
	TB         = 1024 * GB
	PB         = 1024 * TB
	// ...
)
