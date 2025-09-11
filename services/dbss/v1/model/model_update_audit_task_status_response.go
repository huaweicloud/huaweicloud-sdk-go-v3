package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuditTaskStatusResponse Response Object
type UpdateAuditTaskStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAuditTaskStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuditTaskStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateAuditTaskStatusResponse", string(data)}, " ")
}
