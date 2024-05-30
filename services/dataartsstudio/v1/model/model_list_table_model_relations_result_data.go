package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTableModelRelationsResultData data，统一的返回结果的最外层数据结构。
type ListTableModelRelationsResultData struct {
	Value *ListTableModelRelationsResultDataValue `json:"value,omitempty"`
}

func (o ListTableModelRelationsResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableModelRelationsResultData struct{}"
	}

	return strings.Join([]string{"ListTableModelRelationsResultData", string(data)}, " ")
}
