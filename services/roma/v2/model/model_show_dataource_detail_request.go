package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDataourceDetailRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 数据源ID
	DatasourceId string `json:"datasource_id" xml:"datasource_id"`
}

func (o ShowDataourceDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataourceDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowDataourceDetailRequest", string(data)}, " ")
}
