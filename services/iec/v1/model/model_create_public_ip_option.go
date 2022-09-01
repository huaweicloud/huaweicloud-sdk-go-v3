package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建公网IP请求参数
type CreatePublicIpOption struct {

	// 边缘站点的ID。
	SiteId string `json:"site_id" xml:"site_id"`

	// 弹性公网IP的版本。目前IEC服务只支持4，即ipv4。
	IpVersion *string `json:"ip_version,omitempty" xml:"ip_version"`

	// 线路ID。 不传时默认取当前站点第一条线路
	Type *string `json:"type,omitempty" xml:"type"`
}

func (o CreatePublicIpOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicIpOption struct{}"
	}

	return strings.Join([]string{"CreatePublicIpOption", string(data)}, " ")
}
