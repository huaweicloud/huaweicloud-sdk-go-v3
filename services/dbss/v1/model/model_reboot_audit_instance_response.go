package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebootAuditInstanceResponse Response Object
type RebootAuditInstanceResponse struct {

	// 响应状态
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RebootAuditInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebootAuditInstanceResponse struct{}"
	}

	return strings.Join([]string{"RebootAuditInstanceResponse", string(data)}, " ")
}
