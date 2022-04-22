package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 动态属性的元数据信息
type ValueInTwinResponseMetadata struct {

	// 修改时间，UNIX timestamp格式
	Timestamp *string `json:"timestamp,omitempty"`
}

func (o ValueInTwinResponseMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValueInTwinResponseMetadata struct{}"
	}

	return strings.Join([]string{"ValueInTwinResponseMetadata", string(data)}, " ")
}
