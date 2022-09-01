package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 企业注册信息
type CorpBasicInfoDto struct {

	// 企业id
	Id *string `json:"id,omitempty" xml:"id"`

	// 企业名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 企业所在地
	Address *string `json:"address,omitempty" xml:"address"`

	// 管理员名称
	AdminName *string `json:"adminName,omitempty" xml:"adminName"`

	// 管理员账号
	Account *string `json:"account,omitempty" xml:"account"`

	// 管理员手机
	Phone *string `json:"phone,omitempty" xml:"phone"`

	// 管理员手机所属的国家
	Country *string `json:"country,omitempty" xml:"country"`

	// 管理员邮箱
	Email *string `json:"email,omitempty" xml:"email"`

	// 是否发送短信
	EnableSMS *bool `json:"enableSMS,omitempty" xml:"enableSMS"`

	// 是否开启云盘
	EnableCloudDisk *bool `json:"enableCloudDisk,omitempty" xml:"enableCloudDisk"`

	// 是否具有pstn功能
	EnablePstn *bool `json:"enablePstn,omitempty" xml:"enablePstn"`

	// 是否支持自动开户
	AutoUserCreate *bool `json:"autoUserCreate,omitempty" xml:"autoUserCreate"`

	// 企业类型
	CorpType *int32 `json:"corpType,omitempty" xml:"corpType"`
}

func (o CorpBasicInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CorpBasicInfoDto struct{}"
	}

	return strings.Join([]string{"CorpBasicInfoDto", string(data)}, " ")
}
