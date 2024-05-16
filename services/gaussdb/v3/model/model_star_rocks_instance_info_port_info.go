package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StarRocksInstanceInfoPortInfo 端口信息。
type StarRocksInstanceInfoPortInfo struct {

	// MySQL端口号。默认3306。
	MysqlPort *int32 `json:"mysql_port,omitempty"`
}

func (o StarRocksInstanceInfoPortInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StarRocksInstanceInfoPortInfo struct{}"
	}

	return strings.Join([]string{"StarRocksInstanceInfoPortInfo", string(data)}, " ")
}
