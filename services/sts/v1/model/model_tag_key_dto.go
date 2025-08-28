package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagKeyDto 标签键，仅包含key字段。
type TagKeyDto struct {
}

func (o TagKeyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagKeyDto struct{}"
	}

	return strings.Join([]string{"TagKeyDto", string(data)}, " ")
}
