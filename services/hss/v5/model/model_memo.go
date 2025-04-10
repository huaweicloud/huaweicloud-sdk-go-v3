package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Memo 备注信息
type Memo struct {
}

func (o Memo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Memo struct{}"
	}

	return strings.Join([]string{"Memo", string(data)}, " ")
}
