package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListShareAppsSnapshotResponseBodyShareApps struct {

	// 应用包名
	PackageName *string `json:"package_name,omitempty"`

	// 应用对应的版本号列表
	Versions *[]string `json:"versions,omitempty"`
}

func (o ListShareAppsSnapshotResponseBodyShareApps) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListShareAppsSnapshotResponseBodyShareApps struct{}"
	}

	return strings.Join([]string{"ListShareAppsSnapshotResponseBodyShareApps", string(data)}, " ")
}
