package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListStatisticsResponse struct {

	// 月度调用次数
	Count *[]MonthUsed `json:"count,omitempty" xml:"count"`

	// 月度资源用量
	Gbs *[]MonthUsed `json:"gbs,omitempty" xml:"gbs"`

	Statistics     *ListFunctionStatisticsResponseBody `json:"statistics,omitempty" xml:"statistics"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListStatisticsResponse", string(data)}, " ")
}
