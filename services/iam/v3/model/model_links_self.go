package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LinksSelf
type LinksSelf struct {

	// 资源链接地址。
	Self string `json:"self"`
}

func (o LinksSelf) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LinksSelf struct{}"
	}

	return strings.Join([]string{"LinksSelf", string(data)}, " ")
}
