package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePortOption 更新端口请求体
type UpdatePortOption struct {

	// - 功能说明：IP/Mac对列表 - 约束：     IP地址不允许为 “0.0.0.0/0”     如果allowed_address_pairs配置地址池较大的CIDR（掩码小于24位），建议为该port配置一个单独的安全组。     如果allowed_address_pairs为“1.1.1.1/0”，表示关闭源目地址检查开关     为虚拟IP配置后端边缘实例场景，       allowed_address_pairs中配置的IP地址，必须为边缘实例网卡已有的IP地址，否则可能会导致虚拟IP通信异常。       被绑定的边缘实例网卡allowed_address_pairs填“1.1.1.1/0”
	AllowedAddressPairs *[]AllowedAddressPair `json:"allowed_address_pairs,omitempty"`

	// 安全组列表
	SecurityGroups *[]string `json:"security_groups,omitempty"`
}

func (o UpdatePortOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePortOption struct{}"
	}

	return strings.Join([]string{"UpdatePortOption", string(data)}, " ")
}
