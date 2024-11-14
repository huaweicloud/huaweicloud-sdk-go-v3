package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObsDataRepositoryPolicy 后端存储自动同步策略
type ObsDataRepositoryPolicy struct {
	AutoExportPolicy *AutoExportPolicy `json:"auto_export_policy,omitempty"`
}

func (o ObsDataRepositoryPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsDataRepositoryPolicy struct{}"
	}

	return strings.Join([]string{"ObsDataRepositoryPolicy", string(data)}, " ")
}
