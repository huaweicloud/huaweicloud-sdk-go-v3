package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlowLogsResponse Response Object
type ListFlowLogsResponse struct {
	Data           *HttpQueryCfwFlowLogsResponseDtoData `json:"data,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ListFlowLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlowLogsResponse struct{}"
	}

	return strings.Join([]string{"ListFlowLogsResponse", string(data)}, " ")
}
