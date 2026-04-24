package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLifeCycleTypeConfigurationsRequest Request Object
type ListLifeCycleTypeConfigurationsRequest struct {

	// 触发场景类型。POST_CREATE_DESKTOP_SUCCESS：创建桌面成功后，POST_REBUILD_DESKTOP_SUCCESS：重建桌面成功后，POST_REATTACH_DESKTOP_SUCCESS：触发重建的分配用户任务成功后，POST_DESKTOP_DISCONNECTED：桌面断开连接后，POST_DESKTOP_CANNOT_CONNECT：桌面无法连接。不传时返回所有配置项。
	Type *string `json:"type,omitempty"`
}

func (o ListLifeCycleTypeConfigurationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLifeCycleTypeConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"ListLifeCycleTypeConfigurationsRequest", string(data)}, " ")
}
