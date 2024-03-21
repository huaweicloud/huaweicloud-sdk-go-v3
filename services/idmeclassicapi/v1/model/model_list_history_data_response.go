package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHistoryDataResponse Response Object
type ListHistoryDataResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]HistoryDataModelHistoryViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors *[]string `json:"errors,omitempty"`

	PageInfo       *PageInfoViewDto `json:"pageInfo,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListHistoryDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryDataResponse struct{}"
	}

	return strings.Join([]string{"ListHistoryDataResponse", string(data)}, " ")
}
