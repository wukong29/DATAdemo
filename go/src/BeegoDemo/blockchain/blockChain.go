package blockchain

import (
	"github.com/boltdb/bolt"
)

//桶的名称，该桶用于装区块信息
var BUCKET_NAME = "blocks"
//表示最新的区块的key名
var LAST_KEY = "lasthash"

/**
 * 区块链结构体实例定义:用于表示代表一条区块链
 * 该区块链包含以下功能:
 		① 将新产生的区块与已有的区块链接起来，并保存
		② 可以查询某个区块的信息
		③ 可以将所有区块进行遍历，输出区块信息
*/
type BlockChain struct {
	LastHash []byte //最新区块hash
	BoltDb   *bolt.DB
}

/**
 * 用于创建一条区块链，并返回该区块链实例
	解释：由于区块链就是由一个一个的区块组成的，因此,如果要创建一条区块链，那么必须要先
		创建一个区块，该区块作为该条区块链的创世区块。
*/
func NewBlockChain() BlockChain {
	//1、创建创世区块
	genesis := CreateGenesisBlock() //创世区块
	//2、创建一个存储区块数据的文件
	db, err := bolt.Open("chain.db", 0600, nil)
	if err != nil {
		panic(err.Error())
	}
	bl := BlockChain{
		LastHash: genesis.Hash,
		BoltDb:   db,
	}
	//3、把新创建的创世区块存入到chain.db当中的一个桶中
	db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucket([]byte(BUCKET_NAME))
		if err != nil {
			panic(err.Error())
		}
		//将创世区块保存到桶中
		serialBlock, err := genesis.Serialize()
		if err != nil {
			panic(err.Error())
		}
		//把创世区块存入到桶中
		bucket.Put(genesis.Hash, serialBlock)
		//更新指向最新区块的Hash值
		bucket.Put([]byte(LAST_KEY), genesis.Hash)
		return nil
	})
	return bl
}

/**
 * 调用BlockChain的该SaveBlock方法，该方法可以将一个生成的新区块保存到chain.db文件中
 */
func (bc BlockChain) SaveData(data []byte) {
	db := bc.BoltDb

	var lastBlock *Block
	//先查询chain.db中存储的最新的区块
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			panic("boltdb未创建，请重试!")
		}
		lastHash := bucket.Get([]byte(LAST_KEY))
		lastBlockBytes := bucket.Get(lastHash)
		lastBlock, _ = DeSerialize(lastBlockBytes)

		return nil
	})

	//1、先生成一个区块 把data存入到新生成的区块中
	newBlock := NewBlock(lastBlock.Height+1, data, lastBlock.Hash)

	//更新chain.db
	db.Update(func(tx *bolt.Tx) error {

		return nil
	})

	return
}
