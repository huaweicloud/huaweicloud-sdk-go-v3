package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTopObjectsResponse Response Object
type ShowTopObjectsResponse struct {

	// 更新时间
	CurTime *int64 `json:"cur_time,omitempty"`

	// 明细
	TopObjectList *[]TopObject `json:"top_object_list,omitempty"`

	// 总览
	TopObjectOverviewList *[]TopObjectOverview `json:"top_object_overview_list,omitempty"`
	HttpStatusCode        int                  `json:"-"`
}

func (o ShowTopObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopObjectsResponse struct{}"
	}

	return strings.Join([]string{"ShowTopObjectsResponse", string(data)}, " ")
}
