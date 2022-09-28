package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowJobEventResponse struct {

	// 条数
	Count *int32 `json:"count,omitempty"`

	// 作业事件列表
	Events         *[]JobEventRsp `json:"events,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowJobEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobEventResponse struct{}"
	}

	return strings.Join([]string{"ShowJobEventResponse", string(data)}, " ")
}
