package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// L2Id 主题域ID，只读，创建和更新时无需填写。
type L2Id struct {
}

func (o L2Id) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "L2Id struct{}"
	}

	return strings.Join([]string{"L2Id", string(data)}, " ")
}
