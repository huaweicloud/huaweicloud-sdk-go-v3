package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTechnicalAssetsStatisticResponse Response Object
type ShowTechnicalAssetsStatisticResponse struct {

	// 数据连接总数
	Count *int32 `json:"count,omitempty"`

	// 数据连接统计信息
	DatasourceStatistics *[]DataSource `json:"datasource_statistics,omitempty"`
	HttpStatusCode       int           `json:"-"`
}

func (o ShowTechnicalAssetsStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTechnicalAssetsStatisticResponse struct{}"
	}

	return strings.Join([]string{"ShowTechnicalAssetsStatisticResponse", string(data)}, " ")
}
