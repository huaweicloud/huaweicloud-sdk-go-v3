package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifierId 修改者id
type ModifierId struct {
}

func (o ModifierId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifierId struct{}"
	}

	return strings.Join([]string{"ModifierId", string(data)}, " ")
}
