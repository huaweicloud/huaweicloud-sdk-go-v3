package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatabaseSmallVersion 版本信息对象
type DatabaseSmallVersion struct {

	// 参数解释： 数据库版本ID，该字段不会有重复。
	Id string `json:"id"`

	// 参数解释： 数据库版本号。
	Name string `json:"name"`

	// 参数解释： 是否优选版本。
	Favored bool `json:"favored"`
}

func (o DatabaseSmallVersion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseSmallVersion struct{}"
	}

	return strings.Join([]string{"DatabaseSmallVersion", string(data)}, " ")
}
