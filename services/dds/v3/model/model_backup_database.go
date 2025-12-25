package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupDatabase struct {

	// **参数解释：** 数据库引擎 **取值范围：** DDS-Community。
	Type string `json:"type"`

	// **参数解释：** 数据库版本。 **取值范围：** 取值为“5.0”、“4.4”、“4.2”、“4.0”、“3.4”。
	Version string `json:"version"`
}

func (o BackupDatabase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupDatabase struct{}"
	}

	return strings.Join([]string{"BackupDatabase", string(data)}, " ")
}
