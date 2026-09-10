package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListShareAppsSnapshotResponse Response Object
type ListShareAppsSnapshotResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 采集时间。
	CollectTime *string `json:"collect_time,omitempty"`

	// 采集的共享应用信息
	ShareApps *[]ListShareAppsSnapshotResponseBodyShareApps `json:"share_apps,omitempty"`

	PageInfo       *ListCloudPhoneServersModelOfferingsResponseBodyPageInfo `json:"page_info,omitempty"`
	HttpStatusCode int                                                      `json:"-"`
}

func (o ListShareAppsSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListShareAppsSnapshotResponse struct{}"
	}

	return strings.Join([]string{"ListShareAppsSnapshotResponse", string(data)}, " ")
}
