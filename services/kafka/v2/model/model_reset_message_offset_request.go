package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ResetMessageOffsetRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 消费组名称。
	Group string `json:"group" xml:"group"`

	Body *ResetMessageOffsetReq `json:"body,omitempty" xml:"body"`
}

func (o ResetMessageOffsetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetMessageOffsetRequest struct{}"
	}

	return strings.Join([]string{"ResetMessageOffsetRequest", string(data)}, " ")
}
