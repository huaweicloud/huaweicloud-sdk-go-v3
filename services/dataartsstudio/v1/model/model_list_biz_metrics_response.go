package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBizMetricsResponse Response Object
type ListBizMetricsResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListBizMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBizMetricsResponse struct{}"
	}

	return strings.Join([]string{"ListBizMetricsResponse", string(data)}, " ")
}
