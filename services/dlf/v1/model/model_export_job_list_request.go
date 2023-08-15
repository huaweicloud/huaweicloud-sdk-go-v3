package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportJobListRequest Request Object
type ExportJobListRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	Body *ExportJobsReq `json:"body,omitempty"`
}

func (o ExportJobListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportJobListRequest struct{}"
	}

	return strings.Join([]string{"ExportJobListRequest", string(data)}, " ")
}
