package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessPreviewId 访问预览的唯一标识符。
type AccessPreviewId struct {
}

func (o AccessPreviewId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessPreviewId struct{}"
	}

	return strings.Join([]string{"AccessPreviewId", string(data)}, " ")
}
