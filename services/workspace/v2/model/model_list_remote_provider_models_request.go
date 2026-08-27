package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRemoteProviderModelsRequest Request Object
type ListRemoteProviderModelsRequest struct {
	Body *ListRemoteModelsReq `json:"body,omitempty"`
}

func (o ListRemoteProviderModelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRemoteProviderModelsRequest struct{}"
	}

	return strings.Join([]string{"ListRemoteProviderModelsRequest", string(data)}, " ")
}
