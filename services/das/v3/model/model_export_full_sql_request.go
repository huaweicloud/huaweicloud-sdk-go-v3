package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportFullSqlRequest Request Object
type ExportFullSqlRequest struct {
	Body *ExportFullSqlRequestBody `json:"body,omitempty"`
}

func (o ExportFullSqlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportFullSqlRequest struct{}"
	}

	return strings.Join([]string{"ExportFullSqlRequest", string(data)}, " ")
}
