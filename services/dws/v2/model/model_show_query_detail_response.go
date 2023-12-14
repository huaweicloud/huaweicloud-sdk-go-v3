package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQueryDetailResponse Response Object
type ShowQueryDetailResponse struct {

	// 响应码。
	Code *int32 `json:"code,omitempty"`

	// 响应信息。
	Msg *string `json:"msg,omitempty"`

	Data           *ListQueriesDto `json:"data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowQueryDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQueryDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowQueryDetailResponse", string(data)}, " ")
}
