package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCodeTableResultData data，统一的返回结果的最外层数据结构。
type CreateCodeTableResultData struct {
	Value *CodeTableVo `json:"value,omitempty"`
}

func (o CreateCodeTableResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCodeTableResultData struct{}"
	}

	return strings.Join([]string{"CreateCodeTableResultData", string(data)}, " ")
}
