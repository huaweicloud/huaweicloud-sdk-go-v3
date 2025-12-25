package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FilterResp **参数解释**： 聚合方式。         **取值范围**： - average：平均值 - variance：方差 - min：最小值 - max：最大值 - sum：求和
type FilterResp struct {
}

func (o FilterResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FilterResp struct{}"
	}

	return strings.Join([]string{"FilterResp", string(data)}, " ")
}
