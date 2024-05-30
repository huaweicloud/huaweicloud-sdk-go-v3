package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchSubjectResultData data，统一的返回结果的最外层数据结构。
type SearchSubjectResultData struct {
	Value *SearchSubjectResultDataValue `json:"value,omitempty"`
}

func (o SearchSubjectResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchSubjectResultData struct{}"
	}

	return strings.Join([]string{"SearchSubjectResultData", string(data)}, " ")
}
