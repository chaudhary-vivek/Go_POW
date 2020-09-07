/*
package blockchain


import (
	"bytes"
	"crypto/sha256"
	
)

//This struct stores a blockchain
//This struct is a one dimension array of pointers to blocks
type BlockChain struct{
	Blocks []*Block
}


// This struct stores a block
type Block struct{
	Hash		[]byte //hash in this block
	Data		[]byte  //the data in this block
	PrevHash	[]byte //hash of the last block
}

/*
//this method will make a hash based on Data and Prevhash for which we have to use bytes library.
func (b *Block) DeriveHash(){
	//making a two dimensional slice of Cata and previoshash and combining with a separator that is an empty slice of byte
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	//SHA256 is actually a much easier way of calculationg the 
	//now we use sum256 hashing function to make a hash out of this
	hash := sha256.Sum256(info)
	//now we push the created hash into the hash field of the struct
	b.Hash = hash[:] 
}
*/


/*



// This function will output a pointer to a block.
//This function will take a string (data) and the PrevHash
func CreateBlock(data string, prevHash []byte) *Block{
	//&Block is a reference to a block
	//For the hash field in Block struct we put an empty slice of bite
	//The Data field in  the struct will be the data string converted to a slice of bytes using []byte(data)
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	//Now we feed the block to DeriveHash which returns a hash
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}


//Create a method to add a block to the chain
//The function takes chain which is a pointer to our blockchain
func(chain *BlockChain) AddBlock(data string){
	//this gives the previous block in "blocks" which is an array of blocks in Blockchain struct
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	//We use createblock function to make a block out of the string and the Hash in the prevBlock struct
	new := CreateBlock(data, prevBlock.Hash)
	//Now we will append the new block in "blocks" which is a one dimensional array of pointers to Blocks
	chain.Blocks = append(chain.Blocks, new)
}

//All blocks contain reference to a previous block
// So we will create the first block or a genesis block
// In this  function we craete a new block with the string genisis
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

//Now we will call a function to inititate a blockchain
//This returns a pointer to a blockchain
// We initiate it by creating an array of blocks with call to the genesis function
func InitBlockChain() *BlockChain{
	return &BlockChain{[]*Block{Genesis()}}
}


*/

package blockchain

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}