package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPrivateZonesRequest struct {
	// 待查询的zone的类型。  取值范围：private。

	Type string `json:"type"`
	// 每页返回的资源个数，取值范围：0~500

	Limit *int32 `json:"limit,omitempty"`
	// 分页查询起始的资源ID，为空时为查询第一页

	Marker *string `json:"marker,omitempty"`
	// 分页查询起始偏移量，表示从偏移量的下一个资源开始查询

	Offset *int32 `json:"offset,omitempty"`
	// 资源标签。

	Tags *string `json:"tags,omitempty"`
	// zone名称。

	Name *string `json:"name,omitempty"`
	// 资源状态。

	Status *string `json:"status,omitempty"`
	// 域名关联的企业项目ID，长度不超过36个字符。  默认值为0。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListPrivateZonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateZonesRequest struct{}"
	}

	return strings.Join([]string{"ListPrivateZonesRequest", string(data)}, " ")
}
