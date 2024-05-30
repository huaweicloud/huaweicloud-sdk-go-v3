package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BizDatasourceRelationVo 数据源信息。
type BizDatasourceRelationVo struct {

	// 编码，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 业务对象信息，填写String类型替代Long类型。
	BizId *string `json:"biz_id,omitempty"`

	BizType *BizTypeEnum `json:"biz_type,omitempty"`

	// 数据连接类型，对应表所在的数仓类型，取值可以为DLI、DWS、MRS_HIVE、POSTGRESQL、MRS_SPARK、CLICKHOUSE、MYSQL、ORACLE和DORIS等。
	DwType string `json:"dw_type"`

	// 数据连接ID。
	DwId string `json:"dw_id"`

	// 数据连接名，只读。
	DwName *string `json:"dw_name,omitempty"`

	// 数据库名。
	DbName *string `json:"db_name,omitempty"`

	// dli数据连接执行sql所需的队列，数据连接类型为DLI时必须。
	QueueName *string `json:"queue_name,omitempty"`

	// DWS类型需要。
	Schema *string `json:"schema,omitempty"`
}

func (o BizDatasourceRelationVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BizDatasourceRelationVo struct{}"
	}

	return strings.Join([]string{"BizDatasourceRelationVo", string(data)}, " ")
}
