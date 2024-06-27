package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBasicAwInfoListSupportsSearchRequest Request Object
type ListBasicAwInfoListSupportsSearchRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 当前页数
	PageNo string `json:"page_no"`

	// 每页多少记录
	PageSize string `json:"page_size"`

	// 父目录ID
	ParentId *string `json:"parent_id,omitempty"`

	Body *SearchRequest `json:"body,omitempty"`
}

func (o ListBasicAwInfoListSupportsSearchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBasicAwInfoListSupportsSearchRequest struct{}"
	}

	return strings.Join([]string{"ListBasicAwInfoListSupportsSearchRequest", string(data)}, " ")
}
