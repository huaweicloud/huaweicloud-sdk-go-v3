package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVariablesRequest Request Object
type ListVariablesRequest struct {

	// group_id
	GroupId *string `json:"group_id,omitempty"`

	// 当前页数
	PageNo *string `json:"page_no,omitempty"`

	// 每页多少记录
	PageSize *string `json:"page_size,omitempty"`

	// 工程id
	ProjectId string `json:"project_id"`
}

func (o ListVariablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVariablesRequest struct{}"
	}

	return strings.Join([]string{"ListVariablesRequest", string(data)}, " ")
}
