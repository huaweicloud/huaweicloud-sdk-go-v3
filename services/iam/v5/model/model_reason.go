package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Reason 删除失败的原因。
type Reason struct {
}

func (o Reason) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Reason struct{}"
	}

	return strings.Join([]string{"Reason", string(data)}, " ")
}
