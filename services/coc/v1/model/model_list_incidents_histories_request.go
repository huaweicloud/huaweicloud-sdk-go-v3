package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIncidentsHistoriesRequest Request Object
type ListIncidentsHistoriesRequest struct {

	// 事件单ID
	IncidentId string `json:"incident_id"`

	Body *ListTicketParamsV2 `json:"body,omitempty"`
}

func (o ListIncidentsHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIncidentsHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListIncidentsHistoriesRequest", string(data)}, " ")
}
