package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServerHdaUpgradeRecordsResponse Response Object
type ListServerHdaUpgradeRecordsResponse struct {

	// 总数。
	Count *int32 `json:"count,omitempty"`

	// 返回列表条目数量上限为分页的最大上限值。
	Items          *[]ServerHdaUpgradeRecord `json:"items,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListServerHdaUpgradeRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerHdaUpgradeRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListServerHdaUpgradeRecordsResponse", string(data)}, " ")
}
