package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlowStatisticResponse Response Object
type ListFlowStatisticResponse struct {
	Data           *ListFlowStatisticRespData `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListFlowStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlowStatisticResponse struct{}"
	}

	return strings.Join([]string{"ListFlowStatisticResponse", string(data)}, " ")
}
