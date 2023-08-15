package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobInstancesRequest Request Object
type ListJobInstancesRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`
}

func (o ListJobInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListJobInstancesRequest", string(data)}, " ")
}
