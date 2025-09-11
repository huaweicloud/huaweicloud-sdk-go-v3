package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DimNameResp **参数解释**： 指标维度值 **取值范围**： 最小长度1，最大长度32
type DimNameResp struct {
}

func (o DimNameResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DimNameResp struct{}"
	}

	return strings.Join([]string{"DimNameResp", string(data)}, " ")
}
