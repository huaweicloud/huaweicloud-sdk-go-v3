package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeSubjectsResultData data，统一的返回结果的最外层数据结构。
type ChangeSubjectsResultData struct {
	Value *CatalogLevelVoList `json:"value,omitempty"`
}

func (o ChangeSubjectsResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSubjectsResultData struct{}"
	}

	return strings.Join([]string{"ChangeSubjectsResultData", string(data)}, " ")
}
