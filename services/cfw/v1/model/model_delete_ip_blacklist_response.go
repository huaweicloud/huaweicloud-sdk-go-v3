package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIpBlacklistResponse Response Object
type DeleteIpBlacklistResponse struct {
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteIpBlacklistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIpBlacklistResponse struct{}"
	}

	return strings.Join([]string{"DeleteIpBlacklistResponse", string(data)}, " ")
}
