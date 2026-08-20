package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopReferSummary TOP100 Referer数据明细
type TopReferSummary struct {

	// **参数解释：** referer值 **取值范围：** 不涉及
	Refer *string `json:"refer,omitempty"`

	// **参数解释：** 对应查询类型的值 **取值范围：** 若值为流量，流量单位：Byte
	Value *int64 `json:"value,omitempty"`
}

func (o TopReferSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopReferSummary struct{}"
	}

	return strings.Join([]string{"TopReferSummary", string(data)}, " ")
}
