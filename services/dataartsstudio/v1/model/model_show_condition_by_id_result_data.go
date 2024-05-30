package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConditionByIdResultData data，统一的返回结果的最外层数据结构。
type ShowConditionByIdResultData struct {
	Value *ConditionVo `json:"value,omitempty"`
}

func (o ShowConditionByIdResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConditionByIdResultData struct{}"
	}

	return strings.Join([]string{"ShowConditionByIdResultData", string(data)}, " ")
}
