package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowAllRepositoryByTwoProjectIdRequest struct {

	// 分页索引，从1开始计数
	PageIndex *int32 `json:"page_index,omitempty" xml:"page_index"`

	// 每页条目数
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size"`

	// 项目的uuid
	ProjectUuid string `json:"project_uuid" xml:"project_uuid"`

	// 搜索关键字
	Search *string `json:"search,omitempty" xml:"search"`
}

func (o ShowAllRepositoryByTwoProjectIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAllRepositoryByTwoProjectIdRequest struct{}"
	}

	return strings.Join([]string{"ShowAllRepositoryByTwoProjectIdRequest", string(data)}, " ")
}
