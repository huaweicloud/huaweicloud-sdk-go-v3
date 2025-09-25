package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VaultGet struct {
	Billing *Billing `json:"billing"`

	// 存储库自定义描述信息。
	Description *string `json:"description,omitempty"`

	// 存储库ID
	Id string `json:"id"`

	// 存储库名称
	Name string `json:"name"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 存储库资源类型id
	ProviderId string `json:"provider_id"`

	// 资源
	Resources []ResourceResp `json:"resources"`

	// 存储库标签
	Tags *[]Tag `json:"tags,omitempty"`

	// 企业项目id，默认为‘0’。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 是否自动绑定，默认为false，不支持。
	AutoBind *bool `json:"auto_bind,omitempty"`

	BindRules *VaultBindRules `json:"bind_rules,omitempty"`

	// 用户id
	UserId *string `json:"user_id,omitempty"`

	// 创建时间,例如:\"2020-02-05T10:38:34.209782\"
	CreatedAt *string `json:"created_at,omitempty"`

	// 是否开启存储库自动扩容能力（只支持按需存储库）。
	AutoExpand *bool `json:"auto_expand,omitempty"`

	// 存储库smn消息通知开关
	SmnNotify *bool `json:"smn_notify,omitempty"`

	// 存储库容量阈值，已用容量占总容量达到此百分比即发送相关通知
	Threshold *int32 `json:"threshold,omitempty"`

	// 用于标识SMB服务
	SysLockSourceService *string `json:"sys_lock_source_service,omitempty"`

	// 用于标识该存储库是否已锁定
	Locked *bool `json:"locked,omitempty"`

	// 存储库可用区信息，最大支持32字符。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 更新时间,例如:\"2020-02-05T10:38:34.209782\"
	UpdatedAt string `json:"updated_at"`

	// 版本
	Version *string `json:"version,omitempty"`
}

func (o VaultGet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultGet struct{}"
	}

	return strings.Join([]string{"VaultGet", string(data)}, " ")
}
