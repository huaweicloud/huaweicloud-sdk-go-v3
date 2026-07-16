package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolMetadataUpdate 资源池元数据更新信息。
type PoolMetadataUpdate struct {
	Annotations *PoolMetadataUpdateAnnotations `json:"annotations,omitempty"`
}

func (o PoolMetadataUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolMetadataUpdate struct{}"
	}

	return strings.Join([]string{"PoolMetadataUpdate", string(data)}, " ")
}
