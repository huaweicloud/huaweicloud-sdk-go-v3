package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GroupEntity struct {
	// 消费组的名称。  长度不超过32位的字符串，仅包含a~z，A~Z，0~9、下划线（_）和中划线（-）。

	Name string `json:"name"`
}

func (o GroupEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupEntity struct{}"
	}

	return strings.Join([]string{"GroupEntity", string(data)}, " ")
}
