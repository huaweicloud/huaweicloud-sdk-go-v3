package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOttChannelInfoResponse Response Object
type DeleteOttChannelInfoResponse struct {

	// 错误码
	ResultCode *string `json:"result_code,omitempty"`

	// 错误描述
	ResultMsg *string `json:"result_msg,omitempty"`

	// 推流域名
	Domain *string `json:"domain,omitempty"`

	// 组名或应用名，为必填项
	AppName *string `json:"app_name,omitempty"`

	// 频道ID。频道唯一标识，为必填项
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteOttChannelInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOttChannelInfoResponse struct{}"
	}

	return strings.Join([]string{"DeleteOttChannelInfoResponse", string(data)}, " ")
}
