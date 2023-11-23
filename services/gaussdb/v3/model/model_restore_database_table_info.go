package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreDatabaseTableInfo struct {

	// 表名。
	Name *string `json:"name,omitempty"`
}

func (o RestoreDatabaseTableInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreDatabaseTableInfo struct{}"
	}

	return strings.Join([]string{"RestoreDatabaseTableInfo", string(data)}, " ")
}
