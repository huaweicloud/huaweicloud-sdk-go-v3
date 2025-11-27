package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UrlAccessControlTaskRequestBody 需要解禁或封禁的URL信息
type UrlAccessControlTaskRequestBody struct {
	AccessControlUrls *AccessControlUrls `json:"access_control_urls"`
}

func (o UrlAccessControlTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UrlAccessControlTaskRequestBody struct{}"
	}

	return strings.Join([]string{"UrlAccessControlTaskRequestBody", string(data)}, " ")
}
