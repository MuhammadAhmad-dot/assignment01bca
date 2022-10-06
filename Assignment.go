package assignment01bca

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

var LastBlockHash string

type Block struct {
	Tranaction       string
	Nonce            int
	PreviousHash     string
	CurrentblockHash string
}

type Chain struct {
	ChainBlock []Block
}

func Newblock(transaction string, Nonce int, PreviousHash string) *Block {
	blockobj := new(Block)
	blockobj.Tranaction = transaction
	blockobj.Nonce = Nonce
	blockobj.PreviousHash = PreviousHash
	blockobj.CreateHash(transaction + string(Nonce) + PreviousHash)
	return blockobj
}

func (b *Block) CreateHash(stringtohash string) {

	hashsum := sha256.Sum256([]byte(stringtohash))
	b.CurrentblockHash = hex.EncodeToString(hashsum[:])

}

func (b *Chain) ListBlocks() {
	//A method to print all the blocks in a nice format showing block data such
	//as transaction, Nonce, previous hash, current block hash
	LastBlockHash = b.ChainBlock[len(b.ChainBlock)-1].CurrentblockHash

	for i := range b.ChainBlock {
		fmt.Println("Block Number : ", i)
		fmt.Println("Transaction is : ", b.ChainBlock[i].Tranaction)
		fmt.Println("Nonce : ", b.ChainBlock[i].Nonce)
		fmt.Println("Hash of previous Block : ", b.ChainBlock[i].PreviousHash)
		fmt.Println("Current Block Hash : ", b.ChainBlock[i].CurrentblockHash)
		fmt.Printf("\n")
	}

}

func (b *Chain) ChangeBlock() {
	var BlockNum int
	fmt.Println("Enter the block number you want to change : ")
	fmt.Scanln(&BlockNum)
	fmt.Println("Enter the the updated transaction : ")
	reader := bufio.NewReader(os.Stdin)
	updateTransaction, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction in Block %d before update is %s \n", BlockNum, b.ChainBlock[BlockNum].Tranaction)
	b.ChainBlock[BlockNum].Tranaction = updateTransaction
	fmt.Printf("Transaction in Block %d After update is %s \n", BlockNum, b.ChainBlock[BlockNum].Tranaction)
	stringtohash := b.ChainBlock[BlockNum].Tranaction + string(b.ChainBlock[BlockNum].Nonce) + b.ChainBlock[BlockNum].PreviousHash
	b.ChainBlock[BlockNum].CreateHash(stringtohash)
}

func (b *Chain) VerifyChain() {
	var flag int = 0

	for i := 0; i < len(b.ChainBlock); i++ {
		if i == len(b.ChainBlock)-1 {
			if b.ChainBlock[i].CurrentblockHash != LastBlockHash {
				flag = 1
			}
		} else if b.ChainBlock[i].CurrentblockHash != b.ChainBlock[i+1].PreviousHash {
			flag = 1

		}

	}

	if flag == 0 {
		fmt.Printf("Block Chain is not modified. \n")
	} else if flag == 1 {
		fmt.Printf("Block Chain is modified.\n")
	}

}


