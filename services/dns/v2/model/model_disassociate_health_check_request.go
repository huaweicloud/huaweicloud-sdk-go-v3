package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DisassociateHealthCheckRequest struct {
	RecordsetId string `json:"recordset_id"`

	Body *AssociateHealthCheckReq `json:"body,omitempty"`
}

func (o DisassociateHealthCheckRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateHealthCheckRequest struct{}"
	}

	return strings.Join([]string{"DisassociateHealthCheckRequest", string(data)}, " ")
}
