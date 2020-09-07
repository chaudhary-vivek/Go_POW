/*
import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)


// We secure the blockchain ny forcing them to do work to add blocks to the blockchain
// The work should be difficult to do but easy to prove
// The steps are ----------------------------
// Take data from the block
// Create a counter (nonce) which starts at 0
// Create a hash of the data plus the counter
// Check to see if hash meets the requirements
// We repeat till the hash meets requirements
// The requirements are:
//The first few bites must contain zero (This difficulty keeps changing in BTC)

const Difficulty = 12

// This is the proof of work struct
// This contains a pointer to a block
// This will also contain a pointer to a big int

type ProofOfWork struct{
	Block *Block
	Target *big.Int
}

// This function takes a pointer to a block and creates a pointer to a proof of work
func NewProof(b*Block) *ProofOfWork{
	targrt := big.NewInt(1)
	// 256 is the number of bytes in our hash
	// We shift it left by using Lsh function 
	target.Lsh(target, uint(256-Difficulty))
	//Now we take the target and put it in a Proof of work struct along with the block
	pow := &ProofOfWork{b, target}
	return pow
}

// This function will combine the proof of work and data to return a slice of byte
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		//We grb the hash of the previous hash and the data for the current hash and ccombine both using an empty slice of bytes
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

//The validation function which returns a boolean
func (pow *ProofOfWork) Validate() bool{
	var intHash big.Int
	data := pow.InitData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1

}


func ToHex(num int64) []byte{
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil{
		log.Panic(err)
	}
	return buff.Bytes()
}

//This function will create the hash from the counter plus the data
//This runs on the poofof work
func (pow *ProofOfWork) Run() (int, []byte){
	var intHash big.Int
	var hash [32]byte

	nonce := 0
	
	//Prepare data, hash it into SHA256, convert it into integer and compare with target integer
	for nonce < math.MaxInt64{
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])
		// This means that the target has been reached 
		if intHash.Cmp(pow.Target) == -1{
			break
		} else {
			//This is if the target has not been reached
			nonce++
		}
		fmt.Println()
		return nonce, hash[:]
	}

}
*/

package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

// Take the data from the block

// create a counter (nonce) which starts at 0

// create a hash of the data plus the counter

// check the hash to see if it meets a set of requirements

// Requirements:
// The First few bytes must contain 0s

const Difficulty = 18

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}

	}
	fmt.Println()

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)

	}

	return buff.Bytes()
}





