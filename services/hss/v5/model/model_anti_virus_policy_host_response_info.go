package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AntiVirusPolicyHostResponseInfo 自定义查杀策略关联主机信息
type AntiVirusPolicyHostResponseInfo struct {

	// 服务器ID
	HostId *string `json:"host_id,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器私有IP
	PrivateIp *string `json:"private_ip,omitempty"`

	// 弹性公网IP地址
	PublicIp *string `json:"public_ip,omitempty"`

	// 资产重要性，包含如下3种   - important ：重要资产   - common ：一般资产   - test ：测试资产
	AssetValue *string `json:"asset_value,omitempty"`
}

func (o AntiVirusPolicyHostResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AntiVirusPolicyHostResponseInfo struct{}"
	}

	return strings.Join([]string{"AntiVirusPolicyHostResponseInfo", string(data)}, " ")
}
