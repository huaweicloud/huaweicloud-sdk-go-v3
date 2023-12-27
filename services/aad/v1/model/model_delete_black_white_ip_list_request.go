package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBlackWhiteIpListRequest Request Object
type DeleteBlackWhiteIpListRequest struct {
	Body *BlackWhiteIpListRequest `json:"body,omitempty"`
}

func (o DeleteBlackWhiteIpListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBlackWhiteIpListRequest struct{}"
	}

	return strings.Join([]string{"DeleteBlackWhiteIpListRequest", string(data)}, " ")
}
