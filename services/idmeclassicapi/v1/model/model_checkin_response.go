package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckinResponse Response Object
type CheckinResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]VersionModelViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CheckinResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckinResponse struct{}"
	}

	return strings.Join([]string{"CheckinResponse", string(data)}, " ")
}
