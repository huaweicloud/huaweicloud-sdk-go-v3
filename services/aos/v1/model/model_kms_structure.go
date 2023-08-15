package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KmsStructure 如果用户给与的var_value是经过KMS加密的，可以通过传递相关加密信息，资源编排服务在使用的时候会帮助用户进行KMS解密  更多关于KMS加密以及KMS加密的样例代码请见：[[KMS加密使用场景介绍](https://support.huaweicloud.com/productdesc-dew/dew_01_0006.html)](tag:hws)[[KMS加密使用场景介绍](https://support.huaweicloud.com/intl/zh-cn/productdesc-dew/dew_01_0006.html)](tag:hws_hk)[[KMS加密使用场景介绍](https://support.huaweicloud.com/eu/productdesc-dew/dew_01_0006.html)](tag:hws_eu)   **注意：**   * 请确保用户给予资源编排服务的委托中包含对指定秘钥ID的操作权限   * KMS每个月有免费试用的额度，如果超过，则KMS需要收费，此费用不是资源编排服务收取，具体标准见：[[https://www.huaweicloud.com/pricing.html?tab=detail#/dew](https://www.huaweicloud.com/pricing.html?tab=detail#/dew)](tag:hws)[[https://www.huaweicloud.com/intl/zh-cn/pricing/index.html#/kms](https://www.huaweicloud.com/intl/zh-cn/pricing/index.html#/kms)](tag:hws_hk)[[https://www.huaweicloud.com/eu/pricing/index.html#/dew](https://www.huaweicloud.com/eu/pricing/index.html#/dew)](tag:hws_eu)   * KMS加密只代表资源编排服务在存储和传输的时候传递的是密文，但是在stack-events中依然是明文，如果希望在log中以密文形式对待，请在模板中声明sensitive，更多关于sensitive的介绍见：[https://learn.hashicorp.com/tutorials/terraform/sensitive-variables](https://learn.hashicorp.com/tutorials/terraform/sensitive-variables)
type KmsStructure struct {

	// 解密时，资源编排服务应该使用的KMS秘钥的ID，通常是加密时所使用的秘钥ID
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
