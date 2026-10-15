package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkSpaceId **参数解释**：工作空间ID。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。
type WorkSpaceId struct {
}

func (o WorkSpaceId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkSpaceId struct{}"
	}

	return strings.Join([]string{"WorkSpaceId", string(data)}, " ")
}
