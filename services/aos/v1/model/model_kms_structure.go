package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 如果用户给与的var_value是经过KMS加密的，可以通过传递相关加密信息，RF在使用的时候会帮助用户进行KMS解密  更多关于KMS加密介绍见：https://support.huaweicloud.com/productdesc-dew/dew_01_0006.html  * 注意：KMS每个月有免费试用的额度，如果超过，则KMS需要收费，此费用不是RF收取， 具体标准见：https://www.huaweicloud.com/pricing.html?tab=detail#/dew  * 注意：KMS加密只代表RF在存储和传输的时候传递的是密文，但是在stack-events中依然是明文，如果希望在log中以密文形式对待， 请在模板中声名sensitive，更多关于sensitive的介绍见：https://learn.hashicorp.com/tutorials/terraform/sensitive-variables
type KmsStructure struct {

	// 解密时，RF应该使用的KMS秘钥的ID，通常是加密时所使用的秘钥ID
	Id string `json:"id"`

	// 数据加密秘钥所对应的密文
	CipherText string `json:"cipher_text"`
}

func (o KmsStructure) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KmsStructure struct{}"
	}

	return strings.Join([]string{"KmsStructure", string(data)}, " ")
}
