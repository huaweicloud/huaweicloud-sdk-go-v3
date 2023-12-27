package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAttentionsRequest Request Object
type ListAttentionsRequest struct {

	// 页码
	PageNo *int32 `json:"page_no,omitempty"`

	// 每页大小
	PageSize *int32 `json:"page_size,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`
}

func (o ListAttentionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAttentionsRequest struct{}"
	}

	return strings.Join([]string{"ListAttentionsRequest", string(data)}, " ")
}
