package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBizMetricResponse Response Object
type DeleteBizMetricResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteBizMetricResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBizMetricResponse struct{}"
	}

	return strings.Join([]string{"DeleteBizMetricResponse", string(data)}, " ")
}
