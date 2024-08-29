package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPublicipPoolRequest Request Object
type ShowPublicipPoolRequest struct {

	// 公网IP池ID唯一标识
	PublicipPoolId string `json:"publicip_pool_id"`

	// 显示，形式为\"fields=id&fields=name&...\"  支持字段：id/name/size/used/project_id/status/billing_info/created_at/updated_at/type/shared/is_common/description/tags/enterprise_project_id/allow_share_bandwidth_types/public_border_group
	Fields *[]string `json:"fields,omitempty"`
}

func (o ShowPublicipPoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicipPoolRequest struct{}"
	}

	return strings.Join([]string{"ShowPublicipPoolRequest", string(data)}, " ")
}
