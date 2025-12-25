package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CodeSegmentName 代码段名称
type CodeSegmentName struct {
}

func (o CodeSegmentName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CodeSegmentName struct{}"
	}

	return strings.Join([]string{"CodeSegmentName", string(data)}, " ")
}
