package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBizMetricResponse Response Object
type DeleteBizMetricResponse struct {
	Data           *DeleteResultData `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteBizMetricResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBizMetricResponse struct{}"
	}

	return strings.Join([]string{"DeleteBizMetricResponse", string(data)}, " ")
}
