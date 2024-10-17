package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuditInstanceResponse Response Object
type UpdateAuditInstanceResponse struct {

	// 响应状态
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAuditInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuditInstanceResponse struct{}"
	}

	return strings.Join([]string{"UpdateAuditInstanceResponse", string(data)}, " ")
}
