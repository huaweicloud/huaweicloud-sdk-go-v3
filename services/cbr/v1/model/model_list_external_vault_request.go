package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExternalVaultRequest Request Object
type ListExternalVaultRequest struct {

	// 其他区域的项目ID
	ExternalProjectId string `json:"external_project_id"`

	// 每页显示条目数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移值
	Offset *int32 `json:"offset,omitempty"`

	// [保护类型。取值为backup，replication和hybrid。](tag:hws,hws_hk) [保护类型。取值为backup和replication。](tag:ocb) [保护类型。取值为backup。](tag:g42,hk-g42,sbc,dt,fcs_vm,ctc,tm,tlf,cmcc,hcso_dt)
	ProtectType *string `json:"protect_type,omitempty"`

	// 区域ID
	RegionId string `json:"region_id"`

	// 资源类型
	ObjcetType *string `json:"objcet_type,omitempty"`

	// [云类型。取值为public和hybrid。](tag:hws,hws_hk) [云类型。取值为public。](tag:g42,hk-g42,sbc,dt,fcs_vm,ctc,ocb,tm,tlf,cmcc,hcso_dt)
	CloudType *string `json:"cloud_type,omitempty"`

	// 存储库ID，指定存储ID时其他过滤条件不生效。
	VaultId *string `json:"vault_id,omitempty"`
}

func (o ListExternalVaultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExternalVaultRequest struct{}"
	}

	return strings.Join([]string{"ListExternalVaultRequest", string(data)}, " ")
}
