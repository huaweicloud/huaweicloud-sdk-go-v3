package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLtsLogsResponse Response Object
type ListLtsLogsResponse struct {

	// 日志开启状态。
	AccessStatus *string `json:"access_status,omitempty"`

	// LTS日志列表。
	LtsAccessList *[]LtslogInfo `json:"lts_access_list,omitempty"`

	// 总数量
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListLtsLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsLogsResponse struct{}"
	}

	return strings.Join([]string{"ListLtsLogsResponse", string(data)}, " ")
}
