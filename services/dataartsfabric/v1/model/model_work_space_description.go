package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkSpaceDescription 描述。用户输入的描述，最长为255个字符。
type WorkSpaceDescription struct {
}

func (o WorkSpaceDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkSpaceDescription struct{}"
	}

	return strings.Join([]string{"WorkSpaceDescription", string(data)}, " ")
}
