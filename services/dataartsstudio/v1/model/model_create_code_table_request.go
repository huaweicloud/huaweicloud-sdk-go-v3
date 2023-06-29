package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCodeTableRequest Request Object
type CreateCodeTableRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *CodeTableVo `json:"body,omitempty"`
}

func (o CreateCodeTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCodeTableRequest struct{}"
	}

	return strings.Join([]string{"CreateCodeTableRequest", string(data)}, " ")
}
