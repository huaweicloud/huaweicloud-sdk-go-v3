package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartType 启动类型，包含如下:   - now : 立即启动   - later : 稍后启动   - period : 周期启动
type StartType struct {
}

func (o StartType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartType struct{}"
	}

	return strings.Join([]string{"StartType", string(data)}, " ")
}
