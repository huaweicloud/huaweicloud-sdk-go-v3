package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRulesRequest Request Object
type ShowRulesRequest struct {

	// 页码
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRulesRequest struct{}"
	}

	return strings.Join([]string{"ShowRulesRequest", string(data)}, " ")
}
