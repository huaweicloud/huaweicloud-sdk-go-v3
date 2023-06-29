package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBizMetricResponse Response Object
type UpdateBizMetricResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateBizMetricResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBizMetricResponse struct{}"
	}

	return strings.Join([]string{"UpdateBizMetricResponse", string(data)}, " ")
}
