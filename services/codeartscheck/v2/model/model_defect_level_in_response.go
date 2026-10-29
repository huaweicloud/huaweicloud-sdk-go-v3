package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DefectLevelInResponse **参数解释：** 问题级别。 **取值范围：** - 0：致命。 - 1：严重。 - 2：一般。 - 3：提示。
type DefectLevelInResponse struct {
}

func (o DefectLevelInResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DefectLevelInResponse struct{}"
	}

	return strings.Join([]string{"DefectLevelInResponse", string(data)}, " ")
}
