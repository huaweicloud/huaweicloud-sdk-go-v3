package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTableDataRequest Request Object
type ShowTableDataRequest struct {

	// 实例ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Instance string `json:"instance"`

	// 资产guid，获取方法请参见[数据资产guid](dataartsstudio_02_0351.xml)。
	Guid string `json:"guid"`

	// 数据连接id。
	DataConnectionId *string `json:"data_connection_id,omitempty"`

	// 数据源类型。
	SourceType *string `json:"source_type,omitempty"`

	// db名称。
	Database *string `json:"database,omitempty"`

	// schema名称。
	Schema *string `json:"schema,omitempty"`

	// table名称。
	Table *string `json:"table,omitempty"`

	// 队列名称。
	Queue *string `json:"queue,omitempty"`
}

func (o ShowTableDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTableDataRequest struct{}"
	}

	return strings.Join([]string{"ShowTableDataRequest", string(data)}, " ")
}
