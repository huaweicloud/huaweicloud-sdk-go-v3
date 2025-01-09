package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServerHdaUpgradeRecordsRequest Request Object
type ListServerHdaUpgradeRecordsRequest struct {

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 查询的数量，值区间[1-100]。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListServerHdaUpgradeRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerHdaUpgradeRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListServerHdaUpgradeRecordsRequest", string(data)}, " ")
}
