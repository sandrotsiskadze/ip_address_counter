package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
)

const max_count_address = 4294967296

func main() {
	file, err := os.Open("ip_addresses.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var bits big.Int
	var count uint64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		ip := net.ParseIP(text)
		var ip_int big.Int
		ip_int.SetBytes(ip.To4())

		if bits.Bit(int(ip_int.Int64())) == 0 {
			bits.SetBit(&bits, int(ip_int.Int64()), 1)
			count++
		}

		// if all of the unique addresses are seen, then we can't find any more unique address, so we can stop
		if count == max_count_address {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}
