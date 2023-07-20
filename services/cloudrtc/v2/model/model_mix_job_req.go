package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MixJobReq 合流任务参数，转推和录制至少选一个
type MixJobReq struct {
	MixParam *MixParam `json:"mix_param"`

	RecordParam *RecordParam `json:"record_param,omitempty"`
}

func (o MixJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MixJobReq struct{}"
	}

	return strings.Join([]string{"MixJobReq", string(data)}, " ")
}
