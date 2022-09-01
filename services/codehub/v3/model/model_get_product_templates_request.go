package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type GetProductTemplatesRequest struct {

	// 项目ID
	ProjectUuid string `json:"project_uuid" xml:"project_uuid"`

	// 分页页数
	PageNo *int32 `json:"page_no,omitempty" xml:"page_no"`

	// 每页数据数
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size"`
}

func (o GetProductTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetProductTemplatesRequest struct{}"
	}

	return strings.Join([]string{"GetProductTemplatesRequest", string(data)}, " ")
}
