package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPrivacyStatementResponse Response Object
type ShowPrivacyStatementResponse struct {

	// 隐私声明版本号。
	Version *string `json:"version,omitempty"`

	// 隐私声明内容。
	Content *string `json:"content,omitempty"`

	// 隐私声明发布时间。
	PublishTime *sdktime.SdkTime `json:"publish_time,omitempty"`

	// 隐私声明对应的语言。
	Language *string `json:"language,omitempty"`

	// 是否签署过历史隐私声明，非第一次签署，提示用户旧版过期需重新签署。
	SignedHistoryRecord *bool `json:"signed_history_record,omitempty"`

	// 是否签署过当前最新的隐私申明 没有签署需需要提醒用户重新签署。
	SignedNewestRecord *bool `json:"signed_newest_record,omitempty"`
	HttpStatusCode     int   `json:"-"`
}

func (o ShowPrivacyStatementResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivacyStatementResponse struct{}"
	}

	return strings.Join([]string{"ShowPrivacyStatementResponse", string(data)}, " ")
}
