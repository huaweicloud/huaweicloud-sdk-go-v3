package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddBlackWhiteIpListRequest Request Object
type AddBlackWhiteIpListRequest struct {
	Body *BlackWhiteIpListRequest `json:"body,omitempty"`
}

func (o AddBlackWhiteIpListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddBlackWhiteIpListRequest struct{}"
	}

	return strings.Join([]string{"AddBlackWhiteIpListRequest", string(data)}, " ")
}
