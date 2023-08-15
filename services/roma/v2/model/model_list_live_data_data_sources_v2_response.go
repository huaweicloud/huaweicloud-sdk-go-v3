package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLiveDataDataSourcesV2Response Response Object
type ListLiveDataDataSourcesV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 数据源列表
	DataSources    *[]LdDatasourceInfo `json:"data_sources,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListLiveDataDataSourcesV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLiveDataDataSourcesV2Response struct{}"
	}

	return strings.Join([]string{"ListLiveDataDataSourcesV2Response", string(data)}, " ")
}
