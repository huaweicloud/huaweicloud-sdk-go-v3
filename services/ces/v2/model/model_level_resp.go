package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LevelResp **参数解释**： 告警级别。    **取值范围**： 只能为1、2、3、4。1为紧急，2为重要，3为次要，4为提示。
type LevelResp struct {
}

func (o LevelResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LevelResp struct{}"
	}

	return strings.Join([]string{"LevelResp", string(data)}, " ")
}
