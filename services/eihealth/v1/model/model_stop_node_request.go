package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopNodeRequest Request Object
type StopNodeRequest struct {

	// 计算资源id
	Id string `json:"id"`

	// 是否强制关闭，默认为false
	Force *bool `json:"force,omitempty"`
}

func (o StopNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopNodeRequest struct{}"
	}

	return strings.Join([]string{"StopNodeRequest", string(data)}, " ")
}
