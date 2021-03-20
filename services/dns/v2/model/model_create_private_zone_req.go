package model

import (
	"encoding/json"

	"strings"
)

type CreatePrivateZoneReq struct {
	// 待创建的域名。

	Name string `json:"name"`
	// 域名的描述信息。

	Description *string `json:"description,omitempty"`
	// 域名类型。取值：private。

	ZoneType string `json:"zone_type"`
	// 管理该zone的管理员邮箱。

	Email *string `json:"email,omitempty"`
	// 用于填写默认生成的SOA记录中有效缓存时间，以秒为单位。

	Ttl *string `json:"ttl,omitempty"`

	Router *Router `json:"router"`
	// 资源标签。

	Tags *[]Tag `json:"tags,omitempty"`
	// 域名关联的企业项目ID，长度不超过36个字符。  默认值为0。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreatePrivateZoneReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePrivateZoneReq struct{}"
	}

	return strings.Join([]string{"CreatePrivateZoneReq", string(data)}, " ")
}
