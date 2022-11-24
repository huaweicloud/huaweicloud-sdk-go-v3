package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CheckProductHealthyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CheckProductHealthyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckProductHealthyResponse struct{}"
	}

	return strings.Join([]string{"CheckProductHealthyResponse", string(data)}, " ")
}
