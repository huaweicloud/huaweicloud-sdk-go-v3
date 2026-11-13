package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupabaseInstanceCreateRequest 新建托管PG实例请求（manage_type=newly_instance时必填）。
type SupabaseInstanceCreateRequest struct {
	Datastore *SupabaseDatastore `json:"datastore"`

	ChargeInfo *ChargeInfo `json:"charge_info"`

	// PG实例规格码。
	FlavorRef string `json:"flavor_ref"`

	Volume *SupabaseVolumeConfig `json:"volume"`

	// 可用区，Single模式下为1个，需与顶层主AZ一致。
	AvailabilityZone string `json:"availability_zone"`

	// 参数模板ID（可选）。
	ConfigurationId *string `json:"configuration_id,omitempty"`

	// 数据库端口，2100~9500，默认5432。
	Port *int32 `json:"port,omitempty"`

	// 安全组ID。
	SecurityGroupId string `json:"security_group_id"`

	// 数据库密码。
	Password string `json:"password"`

	// 内网IP（可选）。
	DataVip *string `json:"data_vip,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 标签列表，最多20个。
	Tags *[]TagWithKeyValue `json:"tags,omitempty"`
}

func (o SupabaseInstanceCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupabaseInstanceCreateRequest struct{}"
	}

	return strings.Join([]string{"SupabaseInstanceCreateRequest", string(data)}, " ")
}
