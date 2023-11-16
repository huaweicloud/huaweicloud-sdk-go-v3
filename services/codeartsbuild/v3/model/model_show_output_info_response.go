package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOutputInfoResponse Response Object
type ShowOutputInfoResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误
	Error *interface{} `json:"error,omitempty"`

	Result         *OutPutInfoResult `json:"result,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowOutputInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOutputInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowOutputInfoResponse", string(data)}, " ")
}
