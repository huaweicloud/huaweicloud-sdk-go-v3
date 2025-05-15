package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Index 连接编号。
type Index struct {
}

func (o Index) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Index struct{}"
	}

	return strings.Join([]string{"Index", string(data)}, " ")
}
