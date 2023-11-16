package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNoticeResponse Response Object
type ListNoticeResponse struct {

	// 查询模板结果
	Result *[]QueryJobNoticeItems `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListNoticeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNoticeResponse struct{}"
	}

	return strings.Join([]string{"ListNoticeResponse", string(data)}, " ")
}
