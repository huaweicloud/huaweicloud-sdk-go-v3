package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePipelineGroupResponse Response Object
type DeletePipelineGroupResponse struct {

	// 操作是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o DeletePipelineGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePipelineGroupResponse struct{}"
	}

	return strings.Join([]string{"DeletePipelineGroupResponse", string(data)}, " ")
}
