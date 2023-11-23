package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstancesDatastoreResult 数据库信息。
type InstancesDatastoreResult struct {

	// 数据库引擎。
	Type string `json:"type"`

	// 数据库版本号。
	Version string `json:"version"`
}

func (o InstancesDatastoreResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstancesDatastoreResult struct{}"
	}

	return strings.Join([]string{"InstancesDatastoreResult", string(data)}, " ")
}
