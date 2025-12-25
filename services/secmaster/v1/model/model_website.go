package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Website 网站
type Website struct {

	// 网站url
	Value string `json:"value"`

	// 主域名
	MainDomain string `json:"main_domain"`

	// WAF开启状态：OPEN | CLOSE
	ProtectedStatus string `json:"protected_status"`

	// 外网或内网 true：外网 false: 内网
	IsPublic bool `json:"is_public"`

	// 网站备注
	Remark *string `json:"remark,omitempty"`

	// 网站服务器列表
	NameServer []string `json:"name_server"`

	ExtendProperties *WebsiteExtendProperties `json:"extend_properties,omitempty"`
}

func (o Website) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Website struct{}"
	}

	return strings.Join([]string{"Website", string(data)}, " ")
}
