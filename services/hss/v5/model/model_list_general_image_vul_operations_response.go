package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGeneralImageVulOperationsResponse Response Object
type ListGeneralImageVulOperationsResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 漏洞处置记录详情
	DataList       *[]GeneralImageVulOperationsResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                      `json:"-"`
}

func (o ListGeneralImageVulOperationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGeneralImageVulOperationsResponse struct{}"
	}

	return strings.Join([]string{"ListGeneralImageVulOperationsResponse", string(data)}, " ")
}
