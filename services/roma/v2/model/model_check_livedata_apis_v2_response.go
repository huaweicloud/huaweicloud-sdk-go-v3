package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckLivedataApisV2Response Response Object
type CheckLivedataApisV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o CheckLivedataApisV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckLivedataApisV2Response struct{}"
	}

	return strings.Join([]string{"CheckLivedataApisV2Response", string(data)}, " ")
}
