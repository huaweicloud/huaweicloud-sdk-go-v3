package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RegisterScenes2Response struct {

	// 订阅scenes结果。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RegisterScenes2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterScenes2Response struct{}"
	}

	return strings.Join([]string{"RegisterScenes2Response", string(data)}, " ")
}
