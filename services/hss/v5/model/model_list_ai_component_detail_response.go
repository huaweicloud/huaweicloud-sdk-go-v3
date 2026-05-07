package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiComponentDetailResponse Response Object
type ListAiComponentDetailResponse struct {

	// **参数解释**： AI组件详情列表总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**： AI组件列表
	DataList       *[]AiDetailInfoResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListAiComponentDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiComponentDetailResponse struct{}"
	}

	return strings.Join([]string{"ListAiComponentDetailResponse", string(data)}, " ")
}
