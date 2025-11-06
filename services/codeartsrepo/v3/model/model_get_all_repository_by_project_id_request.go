package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetAllRepositoryByProjectIdRequest Request Object
type GetAllRepositoryByProjectIdRequest struct {

	// 分页索引，从1开始计数
	PageIndex *int32 `json:"page_index,omitempty"`

	// 每页条目数
	PageSize *int32 `json:"page_size,omitempty"`

	// 项目ID，获取方式请参见[获取项目ID](codehub_api_0014.xml)。
	ProjectUuid string `json:"project_uuid"`

	// 搜索关键字
	Search *string `json:"search,omitempty"`
}

func (o GetAllRepositoryByProjectIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAllRepositoryByProjectIdRequest struct{}"
	}

	return strings.Join([]string{"GetAllRepositoryByProjectIdRequest", string(data)}, " ")
}
