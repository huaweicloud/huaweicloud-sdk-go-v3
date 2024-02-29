package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGetAllVersionsUsingPostResponse Response Object
type ShowGetAllVersionsUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]VersionModelQueryViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors *[]string `json:"errors,omitempty"`

	PageInfo       *PageInfoViewDto `json:"pageInfo,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowGetAllVersionsUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGetAllVersionsUsingPostResponse struct{}"
	}

	return strings.Join([]string{"ShowGetAllVersionsUsingPostResponse", string(data)}, " ")
}
