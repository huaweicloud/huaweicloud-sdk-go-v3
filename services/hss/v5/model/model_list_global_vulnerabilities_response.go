package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalVulnerabilitiesResponse Response Object
type ListGlobalVulnerabilitiesResponse struct {

	// **参数解释** 漏洞总数 **取值范围** 取值0-65535
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释** 漏洞对应的记录
	DataList       *[]GlobalVulInfo `json:"data_list,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListGlobalVulnerabilitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalVulnerabilitiesResponse struct{}"
	}

	return strings.Join([]string{"ListGlobalVulnerabilitiesResponse", string(data)}, " ")
}
