package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建公网zone请求
type CreatePublicZoneReq struct {

	// Zone名称
	Name string `json:"name" xml:"name"`

	// 描述
	Description *string `json:"description,omitempty" xml:"description"`

	// Zone类型,取值public。
	ZoneType *string `json:"zone_type,omitempty" xml:"zone_type"`

	// 管理该zone的管理员邮箱
	Email *string `json:"email,omitempty" xml:"email"`

	// 用于填写默认生成的SOA记录中有效缓存时间，以秒为单位.
	Ttl *int32 `json:"ttl,omitempty" xml:"ttl"`

	// 域名关联的企业项目ID，长度不超过36个字符.
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 资源标签。
	Tags *[]Tag `json:"tags,omitempty" xml:"tags"`
}

func (o CreatePublicZoneReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicZoneReq struct{}"
	}

	return strings.Join([]string{"CreatePublicZoneReq", string(data)}, " ")
}
