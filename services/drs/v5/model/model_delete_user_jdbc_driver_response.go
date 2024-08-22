package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteUserJdbcDriverResponse Response Object
type DeleteUserJdbcDriverResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteUserJdbcDriverResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUserJdbcDriverResponse struct{}"
	}

	return strings.Join([]string{"DeleteUserJdbcDriverResponse", string(data)}, " ")
}
