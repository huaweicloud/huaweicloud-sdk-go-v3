package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListShareAppsSnapshotRequest Request Object
type ListShareAppsSnapshotRequest struct {

	// 云手机服务器的唯一标识。
	ServerId string `json:"server_id"`

	// 每页返回的资源个数。取值范围：1~500（默认值为100）。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。
	Marker *string `json:"marker,omitempty"`

	// 应用包名称，只包含大小写字母、数字、下划线、点，不能以数字和下划线开头，点不能作为结尾且包名中至少有一个点，长度不超过128
	PackageName *string `json:"package_name,omitempty"`
}

func (o ListShareAppsSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListShareAppsSnapshotRequest struct{}"
	}

	return strings.Join([]string{"ListShareAppsSnapshotRequest", string(data)}, " ")
}
