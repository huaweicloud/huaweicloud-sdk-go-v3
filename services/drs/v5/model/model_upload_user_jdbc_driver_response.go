package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadUserJdbcDriverResponse Response Object
type UploadUserJdbcDriverResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UploadUserJdbcDriverResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadUserJdbcDriverResponse struct{}"
	}

	return strings.Join([]string{"UploadUserJdbcDriverResponse", string(data)}, " ")
}
