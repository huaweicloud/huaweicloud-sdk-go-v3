package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签信息
type EpResourceTag struct {

	// tms资源标签key，值最大长度36个unicode字符。 key不能为空。不能包含非打印字符ASCII(0-31)
	Key string `json:"key"`

	// tms资源标签value值，每个值最大长度43个unicode字符，删除时如果value有值按照key/value删除，如果value没值，则按照key删除，可以为空字符串
	Value string `json:"value"`
}

func (o EpResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EpResourceTag struct{}"
	}

	return strings.Join([]string{"EpResourceTag", string(data)}, " ")
}
