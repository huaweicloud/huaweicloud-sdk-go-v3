package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportTableRequest Request Object
type ExportTableRequest struct {
	Body *ExportTableRequestBody `json:"body,omitempty"`
}

func (o ExportTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTableRequest struct{}"
	}

	return strings.Join([]string{"ExportTableRequest", string(data)}, " ")
}
