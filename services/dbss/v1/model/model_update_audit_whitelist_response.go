package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuditWhitelistResponse Response Object
type UpdateAuditWhitelistResponse struct {

	// 状态  - SUCCESS：成功  - FAILED：失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAuditWhitelistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuditWhitelistResponse struct{}"
	}

	return strings.Join([]string{"UpdateAuditWhitelistResponse", string(data)}, " ")
}
