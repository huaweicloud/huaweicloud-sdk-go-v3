package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DimValueResp **参数解释**： 指标维度值 **取值范围**： 最小长度1，最大长度256
type DimValueResp struct {
}

func (o DimValueResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DimValueResp struct{}"
	}

	return strings.Join([]string{"DimValueResp", string(data)}, " ")
}
