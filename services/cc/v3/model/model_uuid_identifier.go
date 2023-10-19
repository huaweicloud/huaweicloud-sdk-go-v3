package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UuidIdentifier struct {

	// 资源ID标识符。
	Id string `json:"id"`
}

func (o UuidIdentifier) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UuidIdentifier struct{}"
	}

	return strings.Join([]string{"UuidIdentifier", string(data)}, " ")
}
