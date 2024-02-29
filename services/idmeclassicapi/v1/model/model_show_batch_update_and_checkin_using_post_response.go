package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchUpdateAndCheckinUsingPostResponse Response Object
type ShowBatchUpdateAndCheckinUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]VersionModelViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowBatchUpdateAndCheckinUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchUpdateAndCheckinUsingPostResponse struct{}"
	}

	return strings.Join([]string{"ShowBatchUpdateAndCheckinUsingPostResponse", string(data)}, " ")
}
