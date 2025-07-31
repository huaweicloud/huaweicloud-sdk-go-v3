package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIacFileRiskPathsResponse Response Object
type ListIacFileRiskPathsResponse struct {

	// 文件风险路径数量
	TotalNum *int32 `json:"total_num,omitempty"`

	// 文件风险路径列表
	DataList       *[]ListIacFileRiskPathsResponseInfoDataList `json:"data_list,omitempty"`
	HttpStatusCode int                                         `json:"-"`
}

func (o ListIacFileRiskPathsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIacFileRiskPathsResponse struct{}"
	}

	return strings.Join([]string{"ListIacFileRiskPathsResponse", string(data)}, " ")
}
