package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DwType 数据连接类型，对应表所在的数仓类型，取值可以为DLI、DWS、MRS_HIVE、POSTGRESQL、MRS_SPARK、CLICKHOUSE、MYSQL、ORACLE和DORIS等。
type DwType struct {
}

func (o DwType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DwType struct{}"
	}

	return strings.Join([]string{"DwType", string(data)}, " ")
}
