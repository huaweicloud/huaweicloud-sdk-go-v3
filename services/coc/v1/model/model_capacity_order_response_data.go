package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CapacityOrderResponseData struct {

	// **参数解释：** 容量的种类。 **取值范围：** 根据云服务容量数据显示具体容量类型，一般种类 - cpu：CPU核数。 - mem：内存。 - size：云盘大小。
	Type *string `json:"type,omitempty"`

	RankList *CapacityOrderResponseRankList `json:"rank_list,omitempty"`
}

func (o CapacityOrderResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CapacityOrderResponseData struct{}"
	}

	return strings.Join([]string{"CapacityOrderResponseData", string(data)}, " ")
}
