package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddVulWhiteListResponse Response Object
type AddVulWhiteListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddVulWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddVulWhiteListResponse struct{}"
	}

	return strings.Join([]string{"AddVulWhiteListResponse", string(data)}, " ")
}
