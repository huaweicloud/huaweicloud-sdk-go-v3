package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBaseModelsResponse Response Object
type ListBaseModelsResponse struct {

	// 符合条件的总数
	Total *int32 `json:"total,omitempty"`

	// 列表信息
	Models *[]BaseModel `json:"models,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListBaseModelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBaseModelsResponse struct{}"
	}

	return strings.Join([]string{"ListBaseModelsResponse", string(data)}, " ")
}
