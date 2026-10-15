package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkMarkerPageInfo **参数解释**：分页游标信息，用于标识分页查询的位置和数量。
type SparkMarkerPageInfo struct {

	// **参数解释**：下一页游标，用于标识下一页数据的起始位置。为空表示没有更多数据。 **取值范围**：采用UUID格式，长度为36个字符，例如：6c98db52-cac2-4ff1-9a91-b7793e95557d。
	NextMarker *string `json:"next_marker,omitempty"`

	// **参数解释**：上一页游标，用于标识上一页数据的起始位置。为空表示当前页是第一页。 **取值范围**：采用UUID格式，长度为36个字符。
	PreviousMarker *string `json:"previous_marker,omitempty"`

	// **参数解释**：当前页数据条数，用于标识当前返回的数据数量。 **取值范围**：大于等于0的整数，最大不超过100。
	CurrentCount *int32 `json:"current_count,omitempty"`
}

func (o SparkMarkerPageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkMarkerPageInfo struct{}"
	}

	return strings.Join([]string{"SparkMarkerPageInfo", string(data)}, " ")
}
