package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAntivirusHandleHistoryResponse Response Object
type ListAntivirusHandleHistoryResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 病毒查杀历史处置记录
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
