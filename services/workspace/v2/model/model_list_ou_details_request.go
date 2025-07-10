package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOuDetailsRequest Request Object
type ListOuDetailsRequest struct {

	// OU名称。
	OuName *string `json:"ou_name,omitempty"`

	// 用于分页查询，返回桌面数量限制。如果不指定，则返回所有符合条件的桌面。
	Limit *int32 `json:"limit,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListOuDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOuDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListOuDetailsRequest", string(data)}, " ")
}
