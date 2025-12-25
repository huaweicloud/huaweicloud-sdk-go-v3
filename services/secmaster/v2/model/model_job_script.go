package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobScript Job Script 作业脚本
type JobScript struct {
}

func (o JobScript) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobScript struct{}"
	}

	return strings.Join([]string{"JobScript", string(data)}, " ")
}
