package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulHandleHistoryResponse Response Object
type ListVulHandleHistoryResponse struct {

	// 列表
	DataList *[]VulhandleHistoryResponseInfoDataList `json:"data_list,omitempty"`

	// 总数
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListVulHandleHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulHandleHistoryResponse struct{}"
	}

	return strings.Join([]string{"ListVulHandleHistoryResponse", string(data)}, " ")
}
