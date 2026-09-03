package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportInstanceListNewRequest Request Object
type ExportInstanceListNewRequest struct {
	Body *ExportInstanceListNewRequestBody `json:"body,omitempty"`
}

func (o ExportInstanceListNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportInstanceListNewRequest struct{}"
	}

	return strings.Join([]string{"ExportInstanceListNewRequest", string(data)}, " ")
}
