package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RebootNodeRequest struct {

	// 计算资源id
	Id string `json:"id"`

	// 是否强制重启，默认为false
	Force *bool `json:"force,omitempty"`
}

func (o RebootNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebootNodeRequest struct{}"
	}

	return strings.Join([]string{"RebootNodeRequest", string(data)}, " ")
}
