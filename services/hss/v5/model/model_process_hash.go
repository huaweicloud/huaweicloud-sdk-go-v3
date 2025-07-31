package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessHash 进程hash
type ProcessHash struct {
}

func (o ProcessHash) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessHash struct{}"
	}

	return strings.Join([]string{"ProcessHash", string(data)}, " ")
}
