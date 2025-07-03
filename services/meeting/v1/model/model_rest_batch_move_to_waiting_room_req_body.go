package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestBatchMoveToWaitingRoomReqBody 移动到等候室请求类
type RestBatchMoveToWaitingRoomReqBody struct {

	// 需要移入等候室的全部与会者pid
	BatchParticipants *[]string `json:"batchParticipants,omitempty"`
}

func (o RestBatchMoveToWaitingRoomReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestBatchMoveToWaitingRoomReqBody struct{}"
	}

	return strings.Join([]string{"RestBatchMoveToWaitingRoomReqBody", string(data)}, " ")
}
