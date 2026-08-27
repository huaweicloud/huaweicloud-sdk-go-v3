package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRemoteProviderModelsResponse Response Object
type ListRemoteProviderModelsResponse struct {

	// 远端模型总数。
	Total *int32 `json:"total,omitempty"`

	// 远程模型列表。
	RemoteModels   *[]BaseModeInfo `json:"remote_models,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListRemoteProviderModelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRemoteProviderModelsResponse struct{}"
	}

	return strings.Join([]string{"ListRemoteProviderModelsResponse", string(data)}, " ")
}
