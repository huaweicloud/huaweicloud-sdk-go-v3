package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Datastore 数据库版本信息。
type Datastore struct {

	// 数据库类型。
	Type string `json:"type"`

	// 数据库版本。
	Version string `json:"version"`
}

func (o Datastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Datastore struct{}"
	}

	return strings.Join([]string{"Datastore", string(data)}, " ")
}
