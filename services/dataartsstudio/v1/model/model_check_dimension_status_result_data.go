package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckDimensionStatusResultData data，统一的返回结果的最外层数据结构。
type CheckDimensionStatusResultData struct {
	Value *BatchOperationVo `json:"value,omitempty"`
}

func (o CheckDimensionStatusResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDimensionStatusResultData struct{}"
	}

	return strings.Join([]string{"CheckDimensionStatusResultData", string(data)}, " ")
}
