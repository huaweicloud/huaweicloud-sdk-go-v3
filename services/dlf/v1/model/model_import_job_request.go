package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportJobRequest Request Object
type ImportJobRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	Body *ImportFileReq `json:"body,omitempty"`
}

func (o ImportJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportJobRequest struct{}"
	}

	return strings.Join([]string{"ImportJobRequest", string(data)}, " ")
}
