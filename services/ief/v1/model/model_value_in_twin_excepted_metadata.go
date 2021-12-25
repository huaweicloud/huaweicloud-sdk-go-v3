package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 期望值设置的时间信息
type ValueInTwinExceptedMetadata struct {
	// 修改时间，UNIX timestamp格式

	Timestamp *string `json:"timestamp,omitempty"`
}

func (o ValueInTwinExceptedMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValueInTwinExceptedMetadata struct{}"
	}

	return strings.Join([]string{"ValueInTwinExceptedMetadata", string(data)}, " ")
}
