package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityCheckHostResultResponse Response Object
type ListSecurityCheckHostResultResponse struct {

	// **参数解释**: 服务器总数 **取值范围**: 最小值0，最大值2000000
	TotalNum *int64 `json:"total_num,omitempty"`

	// **参数解释**: 数据列表 **取值范围**: 不涉及
	DataList       *[]SecurityCheckHostResultResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o ListSecurityCheckHostResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityCheckHostResultResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityCheckHostResultResponse", string(data)}, " ")
}
