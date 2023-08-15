package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrganizationalUnitsRequest Request Object
type ListOrganizationalUnitsRequest struct {

	// 父节点（根或组织单元）的唯一标识符（ID）。
	ParentId *string `json:"parent_id,omitempty"`

	// 页面中最大结果数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListOrganizationalUnitsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrganizationalUnitsRequest struct{}"
	}

	return strings.Join([]string{"ListOrganizationalUnitsRequest", string(data)}, " ")
}
