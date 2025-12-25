package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GroupIdResp **参数解释** 资源分组ID。 **取值范围** 以\"rg\"开头，后面跟着22个字母或数字
type GroupIdResp struct {
}

func (o GroupIdResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupIdResp struct{}"
	}

	return strings.Join([]string{"GroupIdResp", string(data)}, " ")
}
