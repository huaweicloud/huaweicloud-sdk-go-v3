package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDirectoriesResultData data，统一的返回结果的最外层数据结构。
type ListDirectoriesResultData struct {

	// value，统一的返回结果的外层数据结构。
	Value *[]DirectoryVo `json:"value,omitempty"`
}

func (o ListDirectoriesResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDirectoriesResultData struct{}"
	}

	return strings.Join([]string{"ListDirectoriesResultData", string(data)}, " ")
}
