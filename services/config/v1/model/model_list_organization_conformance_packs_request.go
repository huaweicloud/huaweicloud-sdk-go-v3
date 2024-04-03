package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrganizationConformancePacksRequest Request Object
type ListOrganizationConformancePacksRequest struct {

	// 组织ID。
	OrganizationId string `json:"organization_id"`

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`

	// 组织合规规则包ID。
	OrganizationConformancePackId *string `json:"organization_conformance_pack_id,omitempty"`

	// 合规规则包名称。
	ConformancePackName *string `json:"conformance_pack_name,omitempty"`
}

func (o ListOrganizationConformancePacksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrganizationConformancePacksRequest struct{}"
	}

	return strings.Join([]string{"ListOrganizationConformancePacksRequest", string(data)}, " ")
}
