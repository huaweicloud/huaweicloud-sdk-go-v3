package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDatasourceRequest Request Object
type DeleteDatasourceRequest struct {

	// 数据源id
	DatasourceId string `json:"datasource_id"`
}

func (o DeleteDatasourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatasourceRequest struct{}"
	}

	return strings.Join([]string{"DeleteDatasourceRequest", string(data)}, " ")
}
