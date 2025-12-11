package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAntiVirusResultResponse Response Object
type ListAntiVirusResultResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 病毒查杀结果详细信息列表 **取值范围**: 最小值0，最大值1000
	DataList       *[]AntiVirusResultResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListAntiVirusResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAntiVirusResultResponse struct{}"
	}

	return strings.Join([]string{"ListAntiVirusResultResponse", string(data)}, " ")
}
