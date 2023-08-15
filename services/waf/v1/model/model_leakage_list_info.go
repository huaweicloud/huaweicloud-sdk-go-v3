package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LeakageListInfo struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 规则应用的url
	Url *string `json:"url,omitempty"`

	// 类别（响应码：code，敏感信息：sensitive）
	Category *string `json:"category,omitempty"`

	// 规则内容
	Contents *[]string `json:"contents,omitempty"`

	// 创建规则时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 规则状态，0：关闭，1：开启
	Status *int32 `json:"status,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`
}

func (o LeakageListInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LeakageListInfo struct{}"
	}

	return strings.Join([]string{"LeakageListInfo", string(data)}, " ")
}
