package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeregisterScenes2Response Response Object
type DeregisterScenes2Response struct {

	// 取消订阅成功的SceneApplication。
	Success *[]string `json:"success,omitempty"`

	// 取消订阅失败的SceneApplication。
	Failure        *[]string `json:"failure,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DeregisterScenes2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeregisterScenes2Response struct{}"
	}

	return strings.Join([]string{"DeregisterScenes2Response", string(data)}, " ")
}
