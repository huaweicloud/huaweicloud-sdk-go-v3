package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSyncRequest IAM同步请求的指定参数
type UpdateSyncRequest struct {

	// 是否是全量同步。true为全量同步，false为指定用户、用户组同步。默认值为false。
	IsAllSync *bool `json:"is_all_sync,omitempty"`

	// 指定同步的IAM用户组。
	GroupNames *[]string `json:"group_names,omitempty"`

	// 指定同步的IAM用户。
	UserNames *[]string `json:"user_names,omitempty"`
}

func (o UpdateSyncRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSyncRequest struct{}"
	}

	return strings.Join([]string{"UpdateSyncRequest", string(data)}, " ")
}
