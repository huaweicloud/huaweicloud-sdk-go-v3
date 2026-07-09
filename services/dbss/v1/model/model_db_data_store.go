package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DbDataStore struct {

	// 数据库类型
	Type *string `json:"type,omitempty"`

	// 数据库版本
	Version *string `json:"version,omitempty"`
}

func (o DbDataStore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbDataStore struct{}"
	}

	return strings.Join([]string{"DbDataStore", string(data)}, " ")
}
