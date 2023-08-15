package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIndividualJobReq 更新单流任务请求，转推和录制至少选一个
type UpdateIndividualJobReq struct {
	RecordParam *RecordParam `json:"record_param,omitempty"`
}

func (o UpdateIndividualJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIndividualJobReq struct{}"
	}

	return strings.Join([]string{"UpdateIndividualJobReq", string(data)}, " ")
}
