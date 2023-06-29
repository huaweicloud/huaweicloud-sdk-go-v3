package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckApisV2Response Response Object
type CheckApisV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o CheckApisV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckApisV2Response struct{}"
	}

	return strings.Join([]string{"CheckApisV2Response", string(data)}, " ")
}
