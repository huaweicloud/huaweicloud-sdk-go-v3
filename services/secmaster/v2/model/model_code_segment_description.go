package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CodeSegmentDescription 代码段描述信息
type CodeSegmentDescription struct {
}

func (o CodeSegmentDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CodeSegmentDescription struct{}"
	}

	return strings.Join([]string{"CodeSegmentDescription", string(data)}, " ")
}
