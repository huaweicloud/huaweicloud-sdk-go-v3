package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStandardTemplateResultData data，统一的返回结果的最外层数据结构。
type CreateStandardTemplateResultData struct {
	Value *StandElementFieldVo `json:"value,omitempty"`
}

func (o CreateStandardTemplateResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStandardTemplateResultData struct{}"
	}

	return strings.Join([]string{"CreateStandardTemplateResultData", string(data)}, " ")
}
