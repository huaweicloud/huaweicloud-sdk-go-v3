package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatasetInfo struct {

	// 所属云服务,例如主机安全就填写hss
	Csvc string `json:"csvc"`

	// 启用状态：0未启用；1启用
	Enable string `json:"enable"`

	// 位置信息：1 region；0 global
	IsRegion int64 `json:"is_region"`

	Reference *DatasetInfoReference `json:"reference"`

	// 数据源ID
	SourceId int64 `json:"source_id"`

	// 数据源名称
	SourceName string `json:"source_name"`

	// 目标管道信息
	Target *interface{} `json:"target"`

	// 订阅类型：1租户订阅；2租户行管订阅；3平台行管(当前都为1)
	Type int64 `json:"type"`
}

func (o DatasetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatasetInfo struct{}"
	}

	return strings.Join([]string{"DatasetInfo", string(data)}, " ")
}
