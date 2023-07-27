package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProductTwoTemplatesRequest Request Object
type ListProductTwoTemplatesRequest struct {

	// 项目ID，获取方式请参见[获取项目ID](codehub_api_0014.xml)。
	ProjectUuid string `json:"project_uuid"`

	// 分页页数
	PageNo *int32 `json:"page_no,omitempty"`

	// 每页数据数
	PageSize *int32 `json:"page_size,omitempty"`
}

func (o ListProductTwoTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductTwoTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListProductTwoTemplatesRequest", string(data)}, " ")
}
