package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteDatasourceInfoByIdRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 数据源ID

	DatasourceId string `json:"datasource_id"`
}

func (o DeleteDatasourceInfoByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatasourceInfoByIdRequest struct{}"
	}

	return strings.Join([]string{"DeleteDatasourceInfoByIdRequest", string(data)}, " ")
}
