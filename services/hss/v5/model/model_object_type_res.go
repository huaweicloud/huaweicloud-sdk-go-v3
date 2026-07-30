package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObjectTypeRes **参数解释**: 防护对象类型 **取值范围**: - 0：云服务 - 1：三方
type ObjectTypeRes struct {
}

func (o ObjectTypeRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectTypeRes struct{}"
	}

	return strings.Join([]string{"ObjectTypeRes", string(data)}, " ")
}
