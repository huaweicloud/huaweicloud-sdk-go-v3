package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapTimeStamp 毫秒时间戳
type IsapTimeStamp struct {
}

func (o IsapTimeStamp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapTimeStamp struct{}"
	}

	return strings.Join([]string{"IsapTimeStamp", string(data)}, " ")
}
