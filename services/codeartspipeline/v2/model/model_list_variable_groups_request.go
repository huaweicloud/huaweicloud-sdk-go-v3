package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVariableGroupsRequest Request Object
type ListVariableGroupsRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	Body *ListVariableGroupsReq `json:"body,omitempty"`
}

func (o ListVariableGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVariableGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListVariableGroupsRequest", string(data)}, " ")
}
