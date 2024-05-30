package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTableModelsResultData data，统一的返回结果的最外层数据结构。
type ListTableModelsResultData struct {
	Value *ListTableModelsResultDataValue `json:"value,omitempty"`
}

func (o ListTableModelsResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableModelsResultData struct{}"
	}

	return strings.Join([]string{"ListTableModelsResultData", string(data)}, " ")
}
