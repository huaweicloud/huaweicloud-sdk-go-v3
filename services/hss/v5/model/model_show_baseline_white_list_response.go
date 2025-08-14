package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBaselineWhiteListResponse Response Object
type ShowBaselineWhiteListResponse struct {

	// 基线检查白名单的规则范围 - specific_host：部分主机 - all_host：全部主机
	RuleType *string `json:"rule_type,omitempty"`

	// 基线检查的操作系统 - Linux - Windows
	OsType *string `json:"os_type,omitempty"`

	// 基线检查的检查项标识
	IndexVersion *string `json:"index_version,omitempty"`

	// 基线检查的基线名称
	CheckType *string `json:"check_type,omitempty"`

	// 标准类型，包含如下:   - cn_standard : 等保合规标准   - hw_standard : 云安全实践标准   - cis_standard : 通用安全标准
	Standard *string `json:"standard,omitempty"`

	// 基线检查中检查项的检查类型 - 访问控制 - 服务配置
	Tag *string `json:"tag,omitempty"`

	// 基线检查中检查项的名称
	CheckRuleName *string `json:"check_rule_name,omitempty"`

	// 白名单主机id列表
	HostIdList *[]string `json:"host_id_list,omitempty"`

	// 基线白名单备注
	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBaselineWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBaselineWhiteListResponse struct{}"
	}

	return strings.Join([]string{"ShowBaselineWhiteListResponse", string(data)}, " ")
}
