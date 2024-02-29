package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUpdateByAdminUsingPostResponse Response Object
type ShowUpdateByAdminUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]VersionModelViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowUpdateByAdminUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUpdateByAdminUsingPostResponse struct{}"
	}

	return strings.Join([]string{"ShowUpdateByAdminUsingPostResponse", string(data)}, " ")
}
