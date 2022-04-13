package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StartTestDatasourceRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 数据源ID

	DatasourceId string `json:"datasource_id"`

	Body *DatasourceInfo `json:"body,omitempty"`
}

func (o StartTestDatasourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartTestDatasourceRequest struct{}"
	}

	return strings.Join([]string{"StartTestDatasourceRequest", string(data)}, " ")
}
