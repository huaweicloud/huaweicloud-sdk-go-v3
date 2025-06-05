package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJarPackageStatisticsResponse Response Object
type ListJarPackageStatisticsResponse struct {

	// **参数解释**: 中间件包的统计信息总数 **取值范围**: 取值0-10000
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 中间件包统计信息列表
	DataList       *[]JarPackageStatisticsResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListJarPackageStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJarPackageStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListJarPackageStatisticsResponse", string(data)}, " ")
}
