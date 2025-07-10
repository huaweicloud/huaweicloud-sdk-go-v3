package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetJobInfoResponseBodyJobEntitiesInstanceDatastore 数据库信息。
type GetJobInfoResponseBodyJobEntitiesInstanceDatastore struct {

	// 数据库引擎。
	Type *string `json:"type,omitempty"`

	// 数据库版本。
	Version *string `json:"version,omitempty"`
}

func (o GetJobInfoResponseBodyJobEntitiesInstanceDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetJobInfoResponseBodyJobEntitiesInstanceDatastore struct{}"
	}

	return strings.Join([]string{"GetJobInfoResponseBodyJobEntitiesInstanceDatastore", string(data)}, " ")
}
