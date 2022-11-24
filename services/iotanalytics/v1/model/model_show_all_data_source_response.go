package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAllDataSourceResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 数据源列表
	Datasources    *[]DatasourceRestDto `json:"datasources,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowAllDataSourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAllDataSourceResponse struct{}"
	}

	return strings.Join([]string{"ShowAllDataSourceResponse", string(data)}, " ")
}
