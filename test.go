package main

import (
  "github.com/syndtr/goleveldb/leveldb"
  "fmt"
  "encoding/binary"
  "github.com/ethereum/go-ethereum/core/types"
  "bytes"
  "github.com/ethereum/go-ethereum/rlp"
)
type Hash [32]byte
func (h *Hash) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-32:]
	}
	copy(h[32-len(b):], b)
}
func (h Hash) Bytes() []byte { return h[:] } 
var  (
	headerPrefix       = []byte("h") // headerPrefix + num (uint64 big endian) + hash -> header
	headerHashSuffix   = []byte("n") // headerPrefix + num (uint64 big endian) + headerHashSuffix -> hash
)
func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetBytes(b)
	return h
}

func encodeBlockNumber(number uint64) []byte {
	enc := make([]byte, 8)
	binary.BigEndian.PutUint64(enc, number)
	return enc
}

func headerHashKey(number uint64) []byte {
	return append(append(headerPrefix, encodeBlockNumber(number)...), headerHashSuffix...)
}

func main() {
	// Connection to leveldb
	db, err := leveldb.OpenFile("/Users/XXX/Library/Ethereum/geth/chaindata", nil)

	if(err != nil){
		fmt.Printf("Step 1: %v",err)
		return
	}
	
	//Get Block Header
	data, err := db.Get(headerHashKey(40), nil)

	if(err != nil){
		fmt.Printf("Step 2: %v",err)
		return
	}
	fmt.Printf("%v", data);
	h := BytesToHash(data)
 
	headerKey := append(append(headerPrefix, encodeBlockNumber(40)...), h.Bytes()...)	

	fmt.Printf("Retrieved BlockHeader: %x\n", headerKey)

	//get Block Header data from db
	blockHeaderData, err := db.Get(headerKey, nil)
	if(err != nil) {
		fmt.Println("Error while fetching block from level db: %v\n", err)
		return
	}
	
	ethHeader := new(types.Header)
	tmpByteData := bytes.NewReader(blockHeaderData)
	rlp.Decode(tmpByteData, ethHeader)
	fmt.Printf("Details of Header for block number 40:- Type: %T  Hex: %x Value: %v\n", ethHeader, ethHeader, ethHeader)
}
