package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkSpaceName 工作空间名称。
type WorkSpaceName struct {
}

func (o WorkSpaceName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkSpaceName struct{}"
	}

	return strings.Join([]string{"WorkSpaceName", string(data)}, " ")
}
