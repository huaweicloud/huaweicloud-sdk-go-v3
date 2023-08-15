package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CorpBasicInfoDto 企业注册信息。
type CorpBasicInfoDto struct {

	// 企业id。
	Id *string `json:"id,omitempty"`

	// 企业名称。
	Name *string `json:"name,omitempty"`

	// 企业所在地。
	Address *string `json:"address,omitempty"`

	// 管理员名称。
	AdminName *string `json:"adminName,omitempty"`

	// 管理员的华为云会议帐号。
	Account *string `json:"account,omitempty"`

	// 管理员手机。
	Phone *string `json:"phone,omitempty"`

	// [[手机号所属的国家](https://support.huaweicloud.com/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hws)[[手机号所属的国家](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hk) 。
	Country *string `json:"country,omitempty"`

	// 管理员邮箱。
	Email *string `json:"email,omitempty"`

	// 是否通过短信形式发送会议通知。
	EnableSMS *bool `json:"enableSMS,omitempty"`

	// 是否开启云盘。
	EnableCloudDisk *bool `json:"enableCloudDisk,omitempty"`

	// 是否具有pstn功能。
	EnablePstn *bool `json:"enablePstn,omitempty"`

	// 是否支持自动开户。
	AutoUserCreate *bool `json:"autoUserCreate,omitempty"`

	// 企业类型。 * 0：旗舰版 * 5：免费版 * 6：标准版
	CorpType *int32 `json:"corpType,omitempty"`
}

func (o CorpBasicInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CorpBasicInfoDto struct{}"
	}

	return strings.Join([]string{"CorpBasicInfoDto", string(data)}, " ")
}
