package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VaultCreateResource struct {
	Billing *Billing `json:"billing" xml:"billing"`

	// 存储库自定义描述信息。
	Description *string `json:"description,omitempty" xml:"description"`

	// 存储库ID
	Id string `json:"id" xml:"id"`

	// 存储库名称
	Name string `json:"name" xml:"name"`

	// 项目ID
	ProjectId string `json:"project_id" xml:"project_id"`

	// 存储库资源类型id
	ProviderId string `json:"provider_id" xml:"provider_id"`

	// 存储库资源
	Resources []ResourceResp `json:"resources" xml:"resources"`

	// 存储库标签
	Tags *[]Tag `json:"tags,omitempty" xml:"tags"`

	// 企业项目id，默认为‘0’。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 是否自动绑定，默认为false，不支持。
	AutoBind *bool `json:"auto_bind,omitempty" xml:"auto_bind"`

	BindRules *VaultBindRules `json:"bind_rules,omitempty" xml:"bind_rules"`

	// 用户id
	UserId *string `json:"user_id,omitempty" xml:"user_id"`

	// 创建时间,例如:\"2020-02-05T10:38:34.209782\"
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 是否开启存储库自动扩容能力（只支持按需存储库）
	AutoExpand *bool `json:"auto_expand,omitempty" xml:"auto_expand"`

	// 存储库smn消息通知开关
	SmnNotify *bool `json:"smn_notify,omitempty" xml:"smn_notify"`

	// 存储库容量阈值，已用容量占总容量达到此百分比即发送相关通知
	Threshold *int32 `json:"threshold,omitempty" xml:"threshold"`

	// 包周期创建错误信息
	ErrText *string `json:"errText,omitempty" xml:"errText"`

	// 包周期订购结果
	RetCode *string `json:"retCode,omitempty" xml:"retCode"`

	// 包周期创建订单信息
	Orders *[]CbcOrderResult `json:"orders,omitempty" xml:"orders"`
}

func (o VaultCreateResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultCreateResource struct{}"
	}

	return strings.Join([]string{"VaultCreateResource", string(data)}, " ")
}
