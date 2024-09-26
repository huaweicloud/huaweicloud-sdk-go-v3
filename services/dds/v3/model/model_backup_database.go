package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupDatabase struct {

	// 数据库引擎。 取值：DDS-Community。
	Type string `json:"type"`

	// 数据库版本。取值：“3.4”、“4.0”、“4.2”、“4.4”、“5.0”。
	Version string `json:"version"`
}

func (o BackupDatabase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupDatabase struct{}"
	}

	return strings.Join([]string{"BackupDatabase", string(data)}, " ")
}
