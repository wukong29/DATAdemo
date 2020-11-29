package des

import (
	"crypto/des"
	"CryptoHashCodeClass3/utils"
	"crypto/cipher"
)

/**
 * 使用秘钥key对明文data进行加密
 */
func DESEnCrypt(data []byte, key []byte) ([]byte, error) {
	//三要素：key、data、mode
	//DES:数据加密标准算法 Data Encryption Stardard
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//对明文进行尾部填充
	originText := utils.PKCS5EndPadding(data, block.BlockSize())
	//mode
	blockMode := cipher.NewCBCEncrypter(block, key)
	cipherText := make([]byte, len(originText))
	blockMode.CryptBlocks(cipherText, originText)
	return cipherText, nil
}

/**
 * 使用des算法和秘钥key对密文进行解密，并去除尾部填充
 */
func DESDeCrypt(data []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//mode实例化
	mode := cipher.NewCBCDecrypter(block, key)
	//原始的,最初的: original
	originalText := make([]byte, len(data))
	mode.CryptBlocks(originalText, data)
	//对解密后的明文进行尾部填充内容去除
	originalText = utils.ClearPKCS5Padding(originalText, block.BlockSize())
	return originalText, nil
}
