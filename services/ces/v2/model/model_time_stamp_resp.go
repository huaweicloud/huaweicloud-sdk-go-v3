package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TimeStampResp **参数解释**： 指标采集时间 **取值范围**： 最小值为0
type TimeStampResp struct {
}

func (o TimeStampResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeStampResp struct{}"
	}

	return strings.Join([]string{"TimeStampResp", string(data)}, " ")
}
