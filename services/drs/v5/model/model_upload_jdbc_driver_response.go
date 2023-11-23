package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadJdbcDriverResponse Response Object
type UploadJdbcDriverResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UploadJdbcDriverResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadJdbcDriverResponse struct{}"
	}

	return strings.Join([]string{"UploadJdbcDriverResponse", string(data)}, " ")
}
