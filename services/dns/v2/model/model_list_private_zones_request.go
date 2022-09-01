package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPrivateZonesRequest struct {

	// 待查询的zone的类型。  取值范围：private。
	Type string `json:"type" xml:"type"`

	// 每页返回的资源个数，取值范围：0~500
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 分页查询起始的资源ID，为空时为查询第一页
	Marker *string `json:"marker,omitempty" xml:"marker"`

	// 分页查询起始偏移量，表示从偏移量的下一个资源开始查询
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 资源标签。
	Tags *string `json:"tags,omitempty" xml:"tags"`

	// zone名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 资源状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 域名关联的企业项目ID，长度不超过36个字符。  默认值为0。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`
}

func (o ListPrivateZonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateZonesRequest struct{}"
	}

	return strings.Join([]string{"ListPrivateZonesRequest", string(data)}, " ")
}
