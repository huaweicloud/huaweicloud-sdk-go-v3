package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAntivirusHandleHistoryResponse Response Object
type ListAntivirusHandleHistoryResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 病毒查杀历史处置记录 **取值范围**: 最小值0，最大值100
	DataList       *[]AntiVirusHandleHistory `json:"data_list,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListAntivirusHandleHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAntivirusHandleHistoryResponse struct{}"
	}

	return strings.Join([]string{"ListAntivirusHandleHistoryResponse", string(data)}, " ")
}
