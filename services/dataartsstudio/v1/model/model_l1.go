package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// L1 主题域分组中文名，只读，创建和更新时无需填写。
type L1 struct {
}

func (o L1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "L1 struct{}"
	}

	return strings.Join([]string{"L1", string(data)}, " ")
}
