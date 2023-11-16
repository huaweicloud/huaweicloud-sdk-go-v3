package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePipelineGroupResponse Response Object
type UpdatePipelineGroupResponse struct {

	// 操作是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdatePipelineGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePipelineGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdatePipelineGroupResponse", string(data)}, " ")
}
