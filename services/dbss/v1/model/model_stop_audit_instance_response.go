package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopAuditInstanceResponse Response Object
type StopAuditInstanceResponse struct {

	// 响应状态
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopAuditInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopAuditInstanceResponse struct{}"
	}

	return strings.Join([]string{"StopAuditInstanceResponse", string(data)}, " ")
}
