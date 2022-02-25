package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCorpBasicInfoResponse struct {
	// 企业id

	Id *string `json:"id,omitempty"`
	// 企业名称

	Name *string `json:"name,omitempty"`
	// 企业所在地

	Address *string `json:"address,omitempty"`
	// 管理员名称

	AdminName *string `json:"adminName,omitempty"`
	// 管理员账号

	Account *string `json:"account,omitempty"`
	// 管理员手机

	Phone *string `json:"phone,omitempty"`
	// 管理员手机所属的国家

	Country *string `json:"country,omitempty"`
	// 管理员邮箱

	Email *string `json:"email,omitempty"`
	// 是否发送短信

	EnableSMS *bool `json:"enableSMS,omitempty"`
	// 是否开启云盘

	EnableCloudDisk *bool `json:"enableCloudDisk,omitempty"`
	// 是否具有pstn功能

	EnablePstn *bool `json:"enablePstn,omitempty"`
	// 是否支持自动开户

	AutoUserCreate *bool `json:"autoUserCreate,omitempty"`
	// 企业类型

	CorpType       *int32 `json:"corpType,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowCorpBasicInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCorpBasicInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowCorpBasicInfoResponse", string(data)}, " ")
}
