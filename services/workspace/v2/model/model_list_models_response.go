package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListModelsResponse Response Object
type ListModelsResponse struct {

	// 总数。
	Total *int32 `json:"total,omitempty"`

	// 模型列表项。
	Items          *[]ModelInfoForListResp `json:"items,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListModelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListModelsResponse struct{}"
	}

	return strings.Join([]string{"ListModelsResponse", string(data)}, " ")
}
