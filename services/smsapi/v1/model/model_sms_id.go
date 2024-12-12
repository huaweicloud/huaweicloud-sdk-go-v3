package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmsId struct {

	// 创建时间
	CreateTime *string `json:"createTime,omitempty"`

	// 发送短信号码
	From *string `json:"from,omitempty"`

	// 接收短信号码
	OriginTo *string `json:"originTo,omitempty"`

	// 短信id
	SmsMsgId *string `json:"smsMsgId,omitempty"`

	// 短信当前状态
	Status *string `json:"status,omitempty"`

	// 国家码
	CountryId *string `json:"countryId,omitempty"`

	// 短信拆分条数
	Total *int32 `json:"total,omitempty"`
}

func (o SmsId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmsId struct{}"
	}

	return strings.Join([]string{"SmsId", string(data)}, " ")
}
