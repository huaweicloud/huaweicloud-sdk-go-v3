package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCentralNetworkRequest Request Object
type CreateCentralNetworkRequest struct {
	Body *CreateCentralNetworkRequestBody `json:"body,omitempty"`
}

func (o CreateCentralNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCentralNetworkRequest struct{}"
	}

	return strings.Join([]string{"CreateCentralNetworkRequest", string(data)}, " ")
}
