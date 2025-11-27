package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWebTamperProtectionConfigsResponse Response Object
type ListWebTamperProtectionConfigsResponse struct {

	// **参数解释**: 网页防篡改配置总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 网页防篡改配置列表 **取值范围**: 最少0条，最多200条
	DataList       *[]WebTamperProtectionConfigResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                      `json:"-"`
}

func (o ListWebTamperProtectionConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWebTamperProtectionConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListWebTamperProtectionConfigsResponse", string(data)}, " ")
}
