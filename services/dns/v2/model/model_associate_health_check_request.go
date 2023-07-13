package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateHealthCheckRequest Request Object
type AssociateHealthCheckRequest struct {

	// 待查询的recordset ID。
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
