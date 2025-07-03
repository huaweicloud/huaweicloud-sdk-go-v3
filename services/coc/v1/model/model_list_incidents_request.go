package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIncidentsRequest Request Object
type ListIncidentsRequest struct {
	Body *ListTicketParamsWithPageV2 `json:"body,omitempty"`
}

func (o ListIncidentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIncidentsRequest struct{}"
	}

	return strings.Join([]string{"ListIncidentsRequest", string(data)}, " ")
}
