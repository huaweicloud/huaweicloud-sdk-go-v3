package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDownloadAccessoryUrlResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 附件下载url
	AccessoryUrl   *string `json:"accessory_url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDownloadAccessoryUrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDownloadAccessoryUrlResponse struct{}"
	}

	return strings.Join([]string{"ShowDownloadAccessoryUrlResponse", string(data)}, " ")
}
