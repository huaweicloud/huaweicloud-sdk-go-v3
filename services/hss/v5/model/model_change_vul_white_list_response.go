package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeVulWhiteListResponse Response Object
type ChangeVulWhiteListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeVulWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeVulWhiteListResponse struct{}"
	}

	return strings.Join([]string{"ChangeVulWhiteListResponse", string(data)}, " ")
}
