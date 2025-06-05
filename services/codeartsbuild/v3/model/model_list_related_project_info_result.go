package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRelatedProjectInfoResult 项目列表
type ListRelatedProjectInfoResult struct {

	// 总数
	Total float32 `json:"total,omitempty"`

	// 项目列表
	ProjectInfoList *[]ListRelatedProjectInfoResultProjectInfoList `json:"project_info_list,omitempty"`
}

func (o ListRelatedProjectInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRelatedProjectInfoResult struct{}"
	}

	return strings.Join([]string{"ListRelatedProjectInfoResult", string(data)}, " ")
}
