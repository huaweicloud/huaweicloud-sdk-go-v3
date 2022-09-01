package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CdnIps struct {

	// 需查询的IP地址。
	Ip *string `json:"ip,omitempty" xml:"ip"`

	// 是否是华为云CDN节点。（true:是华为云CDN节点，false:不是华为云CDN节点）
	Belongs *bool `json:"belongs,omitempty" xml:"belongs"`

	// IP归属地省份。（Unknown:表示未知归属地）
	Region *string `json:"region,omitempty" xml:"region"`

	// 运营商名称。如果IP归属地未知，该字段返回null。
	Isp *string `json:"isp,omitempty" xml:"isp"`

	// 平台。如果IP归属地未知，该字段返回null。
	Platform *string `json:"platform,omitempty" xml:"platform"`
}

func (o CdnIps) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdnIps struct{}"
	}

	return strings.Join([]string{"CdnIps", string(data)}, " ")
}
