package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkSpaceId 工作空间ID。
type WorkSpaceId struct {
}

func (o WorkSpaceId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkSpaceId struct{}"
	}

	return strings.Join([]string{"WorkSpaceId", string(data)}, " ")
}
