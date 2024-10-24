package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListModelsResponse Response Object
type ListModelsResponse struct {

	// 符合条件的总数
	Total *int32 `json:"total,omitempty"`

	// 列表信息
	Models *[]ModelInfo `json:"models,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListModelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListModelsResponse struct{}"
	}

	return strings.Join([]string{"ListModelsResponse", string(data)}, " ")
}
