package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartAuditInstanceResponse Response Object
type StartAuditInstanceResponse struct {

	// 响应状态
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartAuditInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartAuditInstanceResponse struct{}"
	}

	return strings.Join([]string{"StartAuditInstanceResponse", string(data)}, " ")
}
