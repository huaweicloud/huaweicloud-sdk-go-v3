package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUrgentVulnerabilitiesResponse Response Object
type ListUrgentVulnerabilitiesResponse struct {

	// **参数解释**： 总数 **取值范围**： 字符长度0-2147483647位
	TotalNum *int64 `json:"total_num,omitempty"`

	// **参数解释**： 应急漏洞列表
	DataList       *[]UrgentVulInfo `json:"data_list,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListUrgentVulnerabilitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUrgentVulnerabilitiesResponse struct{}"
	}

	return strings.Join([]string{"ListUrgentVulnerabilitiesResponse", string(data)}, " ")
}
