package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListModelGroupProvidersResponse Response Object
type ListModelGroupProvidersResponse struct {

	// 总数。
	Total *int32 `json:"total,omitempty"`

	// 模型分组供应商关联详情列表项。
	Items          *[]ModelGroupProviderDetailResp `json:"items,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListModelGroupProvidersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListModelGroupProvidersResponse struct{}"
	}

	return strings.Join([]string{"ListModelGroupProvidersResponse", string(data)}, " ")
}
