package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncJdbcDriverResponse Response Object
type SyncJdbcDriverResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o SyncJdbcDriverResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncJdbcDriverResponse struct{}"
	}

	return strings.Join([]string{"SyncJdbcDriverResponse", string(data)}, " ")
}
