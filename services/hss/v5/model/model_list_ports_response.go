package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPortsResponse Response Object
type ListPortsResponse struct {

	// **参数解释**: 开放端口总数 **取值范围**: 最小值0，最大值10000
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**： 端口信息列表 **取值范围**： 不涉及
	DataList       *[]PortResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListPortsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortsResponse struct{}"
	}

	return strings.Join([]string{"ListPortsResponse", string(data)}, " ")
}
