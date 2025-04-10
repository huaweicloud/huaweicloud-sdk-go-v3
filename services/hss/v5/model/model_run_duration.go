package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunDuration 运行时长，秒
type RunDuration struct {
}

func (o RunDuration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunDuration struct{}"
	}

	return strings.Join([]string{"RunDuration", string(data)}, " ")
}
