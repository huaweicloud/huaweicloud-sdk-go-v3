package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAntiVirusHostResponse Response Object
type ListAntiVirusHostResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// data list
	DataList       *[]AntiVirusHostResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListAntiVirusHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAntiVirusHostResponse struct{}"
	}

	return strings.Join([]string{"ListAntiVirusHostResponse", string(data)}, " ")
}
