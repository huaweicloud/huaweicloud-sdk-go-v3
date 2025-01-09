package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchChangeDesktopNetworkRequest Request Object
type BatchChangeDesktopNetworkRequest struct {
	Body *BatchChangeDesktopNetworkReq `json:"body,omitempty"`
}

func (o BatchChangeDesktopNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchChangeDesktopNetworkRequest struct{}"
	}

	return strings.Join([]string{"BatchChangeDesktopNetworkRequest", string(data)}, " ")
}
