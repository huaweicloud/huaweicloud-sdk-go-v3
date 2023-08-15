package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResolveTaskResponse Response Object
type CreateResolveTaskResponse struct {

	// 任务ID。
	TaskId *string `json:"task_id,omitempty"`

	// 智能信息模板ID，由9位数字组成。
	TplId *string `json:"tpl_id,omitempty"`

	// 短信签名列表。
	SmsSigns *[]string `json:"sms_signs,omitempty"`

	// 用户创建时提交的最大解析次数。
	ResolvingTimes *int32 `json:"resolving_times,omitempty"`

	// 实际已解析数量统计。  > 预留字段。
	ResolvedTimes *int32 `json:"resolved_times,omitempty"`

	// 智能信息编码类型。 - group：群发 - individual：个性化
	AimCodeType *string `json:"aim_code_type,omitempty"`

	// 自定义短链域名，由大小写字母和数字组成的二级域名。
	Domain *string `json:"domain,omitempty"`

	// 失效时间（天）。
	ExpirationTime *int32 `json:"expiration_time,omitempty"`

	// 短链列表。
	Params         *[]CreateResolveTaskParamMode `json:"params,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o CreateResolveTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResolveTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateResolveTaskResponse", string(data)}, " ")
}
