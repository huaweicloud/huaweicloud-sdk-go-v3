package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Hash SHA256
type Hash struct {
}

func (o Hash) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Hash struct{}"
	}

	return strings.Join([]string{"Hash", string(data)}, " ")
}
