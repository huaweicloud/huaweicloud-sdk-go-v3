package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRestorableListResponse Response Object
type ShowRestorableListResponse struct {

	// 可恢复的实例总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 可恢复的实例信息
	RestorableInstances *[]QueryRestoreList `json:"restorable_instances,omitempty"`
	HttpStatusCode      int                 `json:"-"`
}

func (o ShowRestorableListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRestorableListResponse struct{}"
	}

	return strings.Join([]string{"ShowRestorableListResponse", string(data)}, " ")
}
