package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubjectResultData data，统一的返回结果的最外层数据结构。
type CreateSubjectResultData struct {
	Value *CatalogVo `json:"value,omitempty"`
}

func (o CreateSubjectResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubjectResultData struct{}"
	}

	return strings.Join([]string{"CreateSubjectResultData", string(data)}, " ")
}
