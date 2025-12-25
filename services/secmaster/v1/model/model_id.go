package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Id 唯一UUID
type Id struct {
}

func (o Id) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Id struct{}"
	}

	return strings.Join([]string{"Id", string(data)}, " ")
}
