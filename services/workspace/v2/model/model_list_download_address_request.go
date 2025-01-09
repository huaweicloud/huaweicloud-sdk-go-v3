package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDownloadAddressRequest Request Object
type ListDownloadAddressRequest struct {
	Body *ListDownloadAddressRequestBody `json:"body,omitempty"`
}

func (o ListDownloadAddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDownloadAddressRequest struct{}"
	}

	return strings.Join([]string{"ListDownloadAddressRequest", string(data)}, " ")
}
