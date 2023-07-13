package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDataSourceRequest Request Object
type ShowDataSourceRequest struct {

	// 数据源id
	DatasourceId string `json:"datasource_id"`
}

func (o ShowDataSourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataSourceRequest struct{}"
	}

	return strings.Join([]string{"ShowDataSourceRequest", string(data)}, " ")
}
