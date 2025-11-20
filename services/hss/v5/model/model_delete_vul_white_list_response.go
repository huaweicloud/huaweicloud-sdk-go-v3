package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVulWhiteListResponse Response Object
type DeleteVulWhiteListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVulWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVulWhiteListResponse struct{}"
	}

	return strings.Join([]string{"DeleteVulWhiteListResponse", string(data)}, " ")
}
