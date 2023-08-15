package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckApiGroupsV2Response Response Object
type CheckApiGroupsV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o CheckApiGroupsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckApiGroupsV2Response struct{}"
	}

	return strings.Join([]string{"CheckApiGroupsV2Response", string(data)}, " ")
}
