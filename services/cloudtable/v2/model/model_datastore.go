package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建集群数据库参数。
type Datastore struct {

	// controller版本号，默认1.0.6
	Version string `json:"version"`

	// 集群数据库类型
	Type string `json:"type"`
}

func (o Datastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Datastore struct{}"
	}

	return strings.Join([]string{"Datastore", string(data)}, " ")
}
