package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AssociateHealthCheckRequest struct {
	RecordsetId string `json:"recordset_id"`

	Body *AssociateHealthCheckReq `json:"body,omitempty"`
}

func (o AssociateHealthCheckRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateHealthCheckRequest struct{}"
	}

	return strings.Join([]string{"AssociateHealthCheckRequest", string(data)}, " ")
}
