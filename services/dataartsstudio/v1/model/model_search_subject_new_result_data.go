package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchSubjectNewResultData data，统一的返回结果的最外层数据结构。
type SearchSubjectNewResultData struct {
	Value *SearchSubjectNewResultDataValue `json:"value,omitempty"`
}

func (o SearchSubjectNewResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchSubjectNewResultData struct{}"
	}

	return strings.Join([]string{"SearchSubjectNewResultData", string(data)}, " ")
}
