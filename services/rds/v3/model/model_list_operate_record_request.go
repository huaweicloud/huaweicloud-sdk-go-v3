package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOperateRecordRequest Request Object
type ListOperateRecordRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`

	Body *ListOperateRecordRequestBody `json:"body,omitempty"`
}

func (o ListOperateRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOperateRecordRequest struct{}"
	}

	return strings.Join([]string{"ListOperateRecordRequest", string(data)}, " ")
}
