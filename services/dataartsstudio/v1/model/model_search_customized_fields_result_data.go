package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCustomizedFieldsResultData data，统一的返回结果的最外层数据结构。
type SearchCustomizedFieldsResultData struct {
	Value *SearchCustomizedFieldsResultDataValue `json:"value,omitempty"`
}

func (o SearchCustomizedFieldsResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCustomizedFieldsResultData struct{}"
	}

	return strings.Join([]string{"SearchCustomizedFieldsResultData", string(data)}, " ")
}
