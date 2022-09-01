package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 允许来宾自己解除静音请求体。
type RestAllowUnMuteReqBody struct {

	// 是否允许自己解除静音（仅静音时有效），默认为允许。 - 0: 不允许。 - 1: 允许。
	AllowUnmuteByOneself int32 `json:"allowUnmuteByOneself" xml:"allowUnmuteByOneself"`
}

func (o RestAllowUnMuteReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestAllowUnMuteReqBody struct{}"
	}

	return strings.Join([]string{"RestAllowUnMuteReqBody", string(data)}, " ")
}
