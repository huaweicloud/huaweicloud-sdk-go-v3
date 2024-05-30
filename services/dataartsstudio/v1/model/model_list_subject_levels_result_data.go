package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubjectLevelsResultData data，统一的返回结果的最外层数据结构。
type ListSubjectLevelsResultData struct {

	// value，统一的返回结果的外层数据结构。
	Value *[]CatalogLevelVo `json:"value,omitempty"`
}

func (o ListSubjectLevelsResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubjectLevelsResultData struct{}"
	}

	return strings.Join([]string{"ListSubjectLevelsResultData", string(data)}, " ")
}
