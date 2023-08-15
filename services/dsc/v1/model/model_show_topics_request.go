package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTopicsRequest Request Object
type ShowTopicsRequest struct {

	// 页码
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowTopicsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopicsRequest struct{}"
	}

	return strings.Join([]string{"ShowTopicsRequest", string(data)}, " ")
}
