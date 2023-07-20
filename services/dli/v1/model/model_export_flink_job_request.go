package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportFlinkJobRequest Request Object
type ExportFlinkJobRequest struct {
	Body *ExportFlinkJobRequestBody `json:"body,omitempty"`
}

func (o ExportFlinkJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportFlinkJobRequest struct{}"
	}

	return strings.Join([]string{"ExportFlinkJobRequest", string(data)}, " ")
}
