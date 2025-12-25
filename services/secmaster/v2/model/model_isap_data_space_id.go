package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapDataSpaceId 数据空间ID
type IsapDataSpaceId struct {
}

func (o IsapDataSpaceId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapDataSpaceId struct{}"
	}

	return strings.Join([]string{"IsapDataSpaceId", string(data)}, " ")
}
