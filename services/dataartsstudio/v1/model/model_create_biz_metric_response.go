package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBizMetricResponse Response Object
type CreateBizMetricResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CreateBizMetricResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBizMetricResponse struct{}"
	}

	return strings.Join([]string{"CreateBizMetricResponse", string(data)}, " ")
}
