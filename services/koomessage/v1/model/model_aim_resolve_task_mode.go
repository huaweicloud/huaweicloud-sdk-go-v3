package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AimResolveTaskMode 查询发送任务短链请求信息
type AimResolveTaskMode struct {

	// 智能信息模板ID，由9位数字组成。
	TplId *string `json:"tpl_id,omitempty"`

	// 短链的最大解析次数。
	ResolveTimes *int32 `json:"resolve_times,omitempty"`

	// 智能信息编码类型。 - group：群发 - individual：个性化
	AimCodeType *string `json:"aim_code_type,omitempty"`

	// 自定义短链域名，由大小写字母和数字组成的二级域名。
	Domain *string `json:"domain,omitempty"`

	// 失效时间（天）。
	ExpirationTime *int32 `json:"expiration_time,omitempty"`

	// 短链解析详情列表。一次请求最多100个短链。
	Params *[]ResolveTaskParam `json:"params,omitempty"`
}

func (o AimResolveTaskMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AimResolveTaskMode struct{}"
	}

	return strings.Join([]string{"AimResolveTaskMode", string(data)}, " ")
}
