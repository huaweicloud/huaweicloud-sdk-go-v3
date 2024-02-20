package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Urn 唯一的资源名称。
type Urn struct {
}

func (o Urn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Urn struct{}"
	}

	return strings.Join([]string{"Urn", string(data)}, " ")
}
