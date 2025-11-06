package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LineContentDto struct {

	// **参数解释：** 文件行数
	LineNO *int32 `json:"lineNO,omitempty"`

	// **参数解释：** 该行对应内容
	Content *string `json:"content,omitempty"`
}

func (o LineContentDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LineContentDto struct{}"
	}

	return strings.Join([]string{"LineContentDto", string(data)}, " ")
}
