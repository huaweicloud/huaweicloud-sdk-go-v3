package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FileCtime **参数解释**： 文件创建时间 **取值范围**： 非负长整数，时间格式：毫秒级时间戳（UTC时区，从1970-01-01 00:00:00开始计算），单位：ms
type FileCtime struct {
}

func (o FileCtime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileCtime struct{}"
	}

	return strings.Join([]string{"FileCtime", string(data)}, " ")
}
