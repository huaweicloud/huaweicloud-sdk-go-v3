package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrderIncidentResponse Response Object
type ListOrderIncidentResponse struct {

	// 数据总量
	TotalCount *int32 `json:"totalCount,omitempty"`

	// 工单列表
	IncidentInfoList *[]IncidentInfo `json:"incidentInfoList,omitempty"`
	HttpStatusCode   int             `json:"-"`
}

func (o ListOrderIncidentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrderIncidentResponse struct{}"
	}

	return strings.Join([]string{"ListOrderIncidentResponse", string(data)}, " ")
}
