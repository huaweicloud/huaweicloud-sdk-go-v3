package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUpdateAndCheckinUsingPostResponse Response Object
type ShowUpdateAndCheckinUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]VersionModelViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowUpdateAndCheckinUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUpdateAndCheckinUsingPostResponse struct{}"
	}

	return strings.Join([]string{"ShowUpdateAndCheckinUsingPostResponse", string(data)}, " ")
}
