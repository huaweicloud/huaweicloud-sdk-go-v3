package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobNoticeConfigInfoResponse Response Object
type ShowJobNoticeConfigInfoResponse struct {

	// 查询模板结果
	Result *[]QueryJobNoticeItems `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowJobNoticeConfigInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobNoticeConfigInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowJobNoticeConfigInfoResponse", string(data)}, " ")
}
