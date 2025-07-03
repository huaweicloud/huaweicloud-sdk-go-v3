package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestAllowAudienceReqBody 主持人允许观众入会入会请求消息体
type RestAllowAudienceReqBody struct {

	// 主持人是否允许入会 1：允许入会
	AllowAudience *int32 `json:"allowAudience,omitempty"`
}

func (o RestAllowAudienceReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestAllowAudienceReqBody struct{}"
	}

	return strings.Join([]string{"RestAllowAudienceReqBody", string(data)}, " ")
}
