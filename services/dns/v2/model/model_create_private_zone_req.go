package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePrivateZoneReq struct {

	// 待创建的域名。
	Name string `json:"name" xml:"name"`

	// 域名的描述信息。
	Description *string `json:"description,omitempty" xml:"description"`

	// 域名类型。取值：private。
	ZoneType string `json:"zone_type" xml:"zone_type"`

	// 管理该zone的管理员邮箱。
	Email *string `json:"email,omitempty" xml:"email"`

	// 用于填写默认生成的SOA记录中有效缓存时间，以秒为单位。
	Ttl *int32 `json:"ttl,omitempty" xml:"ttl"`

	Router *Router `json:"router" xml:"router"`

	// 资源标签。
	Tags *[]Tag `json:"tags,omitempty" xml:"tags"`

	// 域名关联的企业项目ID，长度不超过36个字符。  默认值为0。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`
}

func (o CreatePrivateZoneReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateZoneReq struct{}"
	}

	return strings.Join([]string{"CreatePrivateZoneReq", string(data)}, " ")
}
