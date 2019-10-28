package main

import (
	"encoding/asn1"
	"encoding/base64"
	"fmt"
	"github.com/hyperledger/fabric/common/util"
	"os"
	"strconv"
)

type Block struct {
	Header *BlockHeader
}

type BlockHeader struct {
	Number       int64
	PreviousHash []byte
	DataHash     []byte
}

func (b *BlockHeader) Bytes() []byte {
	asn1Header := BlockHeader{
		PreviousHash: b.PreviousHash,
		DataHash:     b.DataHash,
	}

	asn1Header.Number = int64(b.Number)

	result, err := asn1.Marshal(asn1Header)
	if err != nil {
		panic(err)
	}
	return result
}

func (b *BlockHeader) Hash() []byte {
	return util.ComputeSHA256(b.Bytes())
}

func genPreviousHash(number string, dataHash string, previousHash string) string {

	block := &Block{}
	block.Header = &BlockHeader{}
	block.Header.Number, _ = strconv.ParseInt(number, 10, 64)
	block.Header.DataHash, _ = base64.StdEncoding.DecodeString(dataHash)
	block.Header.PreviousHash, _ = base64.StdEncoding.DecodeString(previousHash)

	hashed := block.Header.Hash()

	encoded := base64.StdEncoding.EncodeToString(hashed)

	return encoded
}

func main() {
	if len(os.Args) != 4 {
		panic("Wrong argument number.")
	}

	args := os.Args[1:]

	number := args[0]
	dataHash := args[1]
	previousHash := args[2]

	result := genPreviousHash(number, dataHash, previousHash)

	fmt.Println("number: ", number)
	fmt.Println("dataHash: ", dataHash)
	fmt.Println("previousHash: ", previousHash)
	fmt.Println("\nCalculated: ", result)
}
