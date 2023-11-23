package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalVariablesRequest Request Object
type ListGlobalVariablesRequest struct {

	// 每页显示的返回信息的个数，默认值为“100”。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，默认值为“0”。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListGlobalVariablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalVariablesRequest struct{}"
	}

	return strings.Join([]string{"ListGlobalVariablesRequest", string(data)}, " ")
}
