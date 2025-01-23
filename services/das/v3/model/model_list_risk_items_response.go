package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRiskItemsResponse Response Object
type ListRiskItemsResponse struct {

	// 数据库类型
	DatastoreType *string `json:"datastore_type,omitempty"`

	// 风险指标项
	Items          *[]QueryRiskItemsItems `json:"items,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListRiskItemsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRiskItemsResponse struct{}"
	}

	return strings.Join([]string{"ListRiskItemsResponse", string(data)}, " ")
}
