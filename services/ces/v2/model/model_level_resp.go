package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LevelResp **参数解释**： 告警级别。    **取值范围**： 只能为1、2、3、4。 - 1：紧急 - 2：重要 - 3：次要 - 4：提示
type LevelResp struct {
}

func (o LevelResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LevelResp struct{}"
	}

	return strings.Join([]string{"LevelResp", string(data)}, " ")
}
