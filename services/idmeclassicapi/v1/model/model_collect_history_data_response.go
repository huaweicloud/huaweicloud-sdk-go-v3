package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectHistoryDataResponse Response Object
type CollectHistoryDataResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]StatisticsRvo `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CollectHistoryDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectHistoryDataResponse struct{}"
	}

	return strings.Join([]string{"CollectHistoryDataResponse", string(data)}, " ")
}
