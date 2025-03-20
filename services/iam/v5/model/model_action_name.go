package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActionName 三段式的授权项名称，例如\"iam:policies:createV5\"。
type ActionName struct {
}

func (o ActionName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionName struct{}"
	}

	return strings.Join([]string{"ActionName", string(data)}, " ")
}
