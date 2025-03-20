package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TypeName 云服务资源类型名称。
type TypeName struct {
}

func (o TypeName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TypeName struct{}"
	}

	return strings.Join([]string{"TypeName", string(data)}, " ")
}
