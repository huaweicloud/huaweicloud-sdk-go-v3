package model

import (
	"encoding/json"

	"strings"
)

// databases返回参数
type GetDatabases struct {
	// 分片数。

	Dbslot int32 `json:"dbslot"`
	// 分片名称.

	Name string `json:"name"`
	// 状态。

	Status string `json:"status"`
	// 创建时间。

	Created string `json:"created"`
	// 最近更新时间。

	Updated string `json:"updated"`
	// 所在RDS的id。

	Id string `json:"id"`
	// 所在RDS的名称。

	IdName string `json:"idName"`
}

func (o GetDatabases) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GetDatabases struct{}"
	}

	return strings.Join([]string{"GetDatabases", string(data)}, " ")
}
