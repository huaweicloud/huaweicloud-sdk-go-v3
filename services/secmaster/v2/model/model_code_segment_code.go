package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CodeSegmentCode 代码片段
type CodeSegmentCode struct {
}

func (o CodeSegmentCode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CodeSegmentCode struct{}"
	}

	return strings.Join([]string{"CodeSegmentCode", string(data)}, " ")
}
