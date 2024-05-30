package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTableModelResultData data，统一的返回结果的最外层数据结构。
type CreateTableModelResultData struct {
	Value *TableModelVo `json:"value,omitempty"`
}

func (o CreateTableModelResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTableModelResultData struct{}"
	}

	return strings.Join([]string{"CreateTableModelResultData", string(data)}, " ")
}
