package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppWhitelistHostStatusResponse Response Object
type ListAppWhitelistHostStatusResponse struct {

	// data list
	DataList *[]AppWhitelistHostResponseInfo `json:"data_list,omitempty"`

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAppWhitelistHostStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppWhitelistHostStatusResponse struct{}"
	}

	return strings.Join([]string{"ListAppWhitelistHostStatusResponse", string(data)}, " ")
}
