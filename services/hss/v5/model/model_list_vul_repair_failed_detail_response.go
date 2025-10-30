package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulRepairFailedDetailResponse Response Object
type ListVulRepairFailedDetailResponse struct {

	// **参数解释**: 漏洞修复失败原因列表 **取值范围**: 最少0条，最多50条
	DataList       *[]VulRepairFailedDetailInfo `json:"data_list,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListVulRepairFailedDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulRepairFailedDetailResponse struct{}"
	}

	return strings.Join([]string{"ListVulRepairFailedDetailResponse", string(data)}, " ")
}
