package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddBlackWhiteIpListResponse Response Object
type AddBlackWhiteIpListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddBlackWhiteIpListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddBlackWhiteIpListResponse struct{}"
	}

	return strings.Join([]string{"AddBlackWhiteIpListResponse", string(data)}, " ")
}
