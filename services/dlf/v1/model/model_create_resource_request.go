package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourceRequest Request Object
type CreateResourceRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	Body *ResourceInfo `json:"body,omitempty"`
}

func (o CreateResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceRequest struct{}"
	}

	return strings.Join([]string{"CreateResourceRequest", string(data)}, " ")
}
