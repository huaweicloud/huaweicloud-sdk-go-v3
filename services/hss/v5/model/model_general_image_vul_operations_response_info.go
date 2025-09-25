package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GeneralImageVulOperationsResponseInfo 漏洞处置记录详情
type GeneralImageVulOperationsResponseInfo struct {

	// 镜像id
	ImageId *string `json:"image_id,omitempty"`

	// 镜像名称
	ImageName *string `json:"image_name,omitempty"`

	// 处置用户名称
	UserName *string `json:"user_name,omitempty"`

	// 处置时间，时间单位：毫秒（ms）
	HandleTime *int64 `json:"handle_time,omitempty"`

	// 操作类型，包含如下：   -ignore : 忽略   -not_ignore : 取消忽略   -add_to_whitelist ：加入白名单
	HandleType *string `json:"handle_type,omitempty"`

	// 漏洞当前状态，包含如下：   -vul_status_unfix : 未处理   -vul_status_ignored : 已忽略
	Status *string `json:"status,omitempty"`

	// 软件名称
	AppName *string `json:"app_name,omitempty"`

	// 软件版本
	AppVersion *string `json:"app_version,omitempty"`

	// 软件路径
	AppPath *string `json:"app_path,omitempty"`

	// 备注
	Remark *string `json:"remark,omitempty"`

	// 镜像标识
	ImageDigest *string `json:"image_digest,omitempty"`

	// 镜像版本
	ImageVersion *string `json:"image_version,omitempty"`

	// Agent ID
	AgentId *string `json:"agent_id,omitempty"`
}

func (o GeneralImageVulOperationsResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeneralImageVulOperationsResponseInfo struct{}"
	}

	return strings.Join([]string{"GeneralImageVulOperationsResponseInfo", string(data)}, " ")
}
