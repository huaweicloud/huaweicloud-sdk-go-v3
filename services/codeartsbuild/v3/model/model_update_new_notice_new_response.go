package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNewNoticeNewResponse Response Object
type UpdateNewNoticeNewResponse struct {

	// 返回错误信息
	Result *string `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateNewNoticeNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNewNoticeNewResponse struct{}"
	}

	return strings.Join([]string{"UpdateNewNoticeNewResponse", string(data)}, " ")
}
