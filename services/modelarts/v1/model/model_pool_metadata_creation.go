package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolMetadataCreation 创建资源池时metadata信息。
type PoolMetadataCreation struct {
	Labels *PoolLabelsCreation `json:"labels"`

	Annotations *PoolAnnotationsCreation `json:"annotations,omitempty"`
}

func (o PoolMetadataCreation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolMetadataCreation struct{}"
	}

	return strings.Join([]string{"PoolMetadataCreation", string(data)}, " ")
}
