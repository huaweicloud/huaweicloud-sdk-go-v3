package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAddonInstancesRequest Request Object
type ListAddonInstancesRequest struct {

	// 集群id
	ClusterId string `json:"cluster_id"`

	// 是否使用数据库存储的插件状态
	AddonTemplateName *string `json:"addon_template_name,omitempty"`

	// 插件包名字
	IsDatabaseStatus *string `json:"is_database_status,omitempty"`
}

func (o ListAddonInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddonInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListAddonInstancesRequest", string(data)}, " ")
}
