package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulHostProcessResponse Response Object
type ListVulHostProcessResponse struct {

	// **参数解释**: 受影响服务器进程列表 **取值范围**: 最小值0，最大值10000000
	DataList *[]VulHostProcessResponseInfoDataList `json:"data_list,omitempty"`

	// **参数解释**： 受影响服务器进程数 **取值范围**： 最小值0，最大值10000000总数
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListVulHostProcessResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulHostProcessResponse struct{}"
	}

	return strings.Join([]string{"ListVulHostProcessResponse", string(data)}, " ")
}
