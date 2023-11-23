package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteJdbcDriverResponse Response Object
type DeleteJdbcDriverResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteJdbcDriverResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJdbcDriverResponse struct{}"
	}

	return strings.Join([]string{"DeleteJdbcDriverResponse", string(data)}, " ")
}
