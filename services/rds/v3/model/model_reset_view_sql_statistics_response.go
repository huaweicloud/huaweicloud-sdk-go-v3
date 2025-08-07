package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetViewSqlStatisticsResponse Response Object
type ResetViewSqlStatisticsResponse struct {

	// 参数解释： 调用结果。 取值范围： 调用正常时，返回“successful”。
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResetViewSqlStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetViewSqlStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ResetViewSqlStatisticsResponse", string(data)}, " ")
}
