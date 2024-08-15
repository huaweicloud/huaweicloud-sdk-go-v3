package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstancePatchItemsResponse Response Object
type ShowInstancePatchItemsResponse struct {

	// 总条数
	Count *int64 `json:"count,omitempty"`

	// 补丁合规信息
	ComplianceItems *[]ComplianceItem `json:"compliance_items,omitempty"`
	HttpStatusCode  int               `json:"-"`
}

func (o ShowInstancePatchItemsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstancePatchItemsResponse struct{}"
	}

	return strings.Join([]string{"ShowInstancePatchItemsResponse", string(data)}, " ")
}
