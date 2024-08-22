package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncUserJdbcDriverResponse Response Object
type SyncUserJdbcDriverResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o SyncUserJdbcDriverResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncUserJdbcDriverResponse struct{}"
	}

	return strings.Join([]string{"SyncUserJdbcDriverResponse", string(data)}, " ")
}
