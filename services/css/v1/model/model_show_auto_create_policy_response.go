package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAutoCreatePolicyResponse struct {

	// 快照保留的天数。
	Keepday *int32 `json:"keepday,omitempty" xml:"keepday"`

	// 每天快照创建时刻。
	Period *string `json:"period,omitempty" xml:"period"`

	// 快照命名前缀，需要用户自己手动输入。
	Prefix *string `json:"prefix,omitempty" xml:"prefix"`

	// 快照存放的OBS桶的桶名。
	Bucket *string `json:"bucket,omitempty" xml:"bucket"`

	// 快照在OBS桶中的存放路径。
	BasePath *string `json:"basePath,omitempty" xml:"basePath"`

	// 访问OBS桶用到的委托。
	Agency *string `json:"agency,omitempty" xml:"agency"`

	// 是否开启自动创建快照策略。 - true：表示开启自动创建快照策略。 - false：表示关闭自动创建快照策略。
	Enable         *string `json:"enable,omitempty" xml:"enable"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAutoCreatePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoCreatePolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoCreatePolicyResponse", string(data)}, " ")
}
