package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GeneralImageVulOperationsResponseInfo 漏洞处置记录详情
type GeneralImageVulOperationsResponseInfo struct {

	// **参数解释**： 镜像id **取值范围**： 字符长度0-128位
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**： 镜像名称 **取值范围**： 字符长度0-128位
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**： 处置用户名称 **取值范围**： 字符长度0-256位
	UserName *string `json:"user_name,omitempty"`

	// 处置时间，时间单位：毫秒（ms）
	HandleTime *int64 `json:"handle_time,omitempty"`

	// **参数解释**： 操作类型 **取值范围**： - ignore：忽略 - not_ignore：取消忽略 - add_to_whitelist：加入白名单
	HandleType *string `json:"handle_type,omitempty"`

	// **参数解释**： 漏洞当前状态 **取值范围**： - vul_status_unfix：未处理 - vul_status_ignored：已忽略
	Status *string `json:"status,omitempty"`

	// **参数解释**： 软件名称 **取值范围**： 字符长度0-256位
	AppName *string `json:"app_name,omitempty"`

	// **参数解释**： 软件版本 **取值范围**： 字符长度0-256位
	AppVersion *string `json:"app_version,omitempty"`

	// **参数解释**： 软件路径 **取值范围**： 字符长度0-256位
	AppPath *string `json:"app_path,omitempty"`

	// **参数解释**： 备注 **取值范围**： 字符长度0-256位
	Remark *string `json:"remark,omitempty"`

	// **参数解释**： 镜像标识 **取值范围**： 字符长度0-256位
	ImageDigest *string `json:"image_digest,omitempty"`

	// **参数解释**： 镜像版本 **取值范围**： 字符长度0-256位
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释**： Agent ID **取值范围**： 字符长度0-512位
	AgentId *string `json:"agent_id,omitempty"`
}

func (o GeneralImageVulOperationsResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeneralImageVulOperationsResponseInfo struct{}"
	}

	return strings.Join([]string{"GeneralImageVulOperationsResponseInfo", string(data)}, " ")
}
