package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2xEdgeAppResponse Response Object
type ListV2xEdgeAppResponse struct {

	// **参数说明**：满足查询条件的记录数。
	Count *int64 `json:"count,omitempty"`

	// **参数说明**：已部署的边缘应用列表。
	Apps           *[]V2XEdgeAppResponseDto `json:"apps,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListV2xEdgeAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2xEdgeAppResponse struct{}"
	}

	return strings.Join([]string{"ListV2xEdgeAppResponse", string(data)}, " ")
}
