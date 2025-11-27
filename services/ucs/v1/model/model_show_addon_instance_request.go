package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAddonInstanceRequest Request Object
type ShowAddonInstanceRequest struct {

	// 插件示例id
	Id string `json:"id"`

	// 是否使用数据库存储的插件状态
	IsDatabaseStatus *string `json:"is_database_status,omitempty"`

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o ShowAddonInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAddonInstanceRequest struct{}"
	}

	return strings.Join([]string{"ShowAddonInstanceRequest", string(data)}, " ")
}
