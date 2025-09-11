package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddAuditWhitelistResponse Response Object
type BatchAddAuditWhitelistResponse struct {

	// 状态  - SUCCESS:成功  - FAILED:失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchAddAuditWhitelistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddAuditWhitelistResponse struct{}"
	}

	return strings.Join([]string{"BatchAddAuditWhitelistResponse", string(data)}, " ")
}
