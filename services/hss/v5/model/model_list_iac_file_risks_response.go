package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIacFileRisksResponse Response Object
type ListIacFileRisksResponse struct {

	// IAC文件风险列表总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// IAC文件风险列表
	DataList       *[]ListIacFileRisksResponseInfoDataList `json:"data_list,omitempty"`
	HttpStatusCode int                                     `json:"-"`
}

func (o ListIacFileRisksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIacFileRisksResponse struct{}"
	}

	return strings.Join([]string{"ListIacFileRisksResponse", string(data)}, " ")
}
