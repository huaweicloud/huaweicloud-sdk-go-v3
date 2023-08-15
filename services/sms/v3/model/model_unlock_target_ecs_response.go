package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnlockTargetEcsResponse Response Object
type UnlockTargetEcsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UnlockTargetEcsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnlockTargetEcsResponse struct{}"
	}

	return strings.Join([]string{"UnlockTargetEcsResponse", string(data)}, " ")
}
