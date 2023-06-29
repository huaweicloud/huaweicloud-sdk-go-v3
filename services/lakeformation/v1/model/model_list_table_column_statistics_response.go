package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTableColumnStatisticsResponse Response Object
type ListTableColumnStatisticsResponse struct {
	Body           *[]ColumnStatisticsObj `json:"body,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListTableColumnStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableColumnStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListTableColumnStatisticsResponse", string(data)}, " ")
}
