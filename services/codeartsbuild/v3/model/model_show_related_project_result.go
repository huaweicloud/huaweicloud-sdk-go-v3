package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRelatedProjectResult 当前用户的项目信息列表
type ShowRelatedProjectResult struct {

	// 总数
	Total float32 `json:"total,omitempty"`

	// 项目信息列表
	ProjectInfoList *[]ShowRelatedProjectResultProjectInfoList `json:"project_info_list,omitempty"`
}

func (o ShowRelatedProjectResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRelatedProjectResult struct{}"
	}

	return strings.Join([]string{"ShowRelatedProjectResult", string(data)}, " ")
}
