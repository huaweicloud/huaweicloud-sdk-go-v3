package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListQueryHttpCodeResponse struct {

	// 基于时间轴的状态码
	DataSeries *[]HttpCodeSummary `json:"data_series,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListQueryHttpCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueryHttpCodeResponse struct{}"
	}

	return strings.Join([]string{"ListQueryHttpCodeResponse", string(data)}, " ")
}
