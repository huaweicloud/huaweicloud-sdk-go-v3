package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateInstanceBy3rdRequest struct {

	// 实例标签（不同的第三方需要和CodeArtsIDEOnline服务共同设定标签）。不传默认为classroom
	InstanceLabel *string `json:"instance_label,omitempty"`

	Body *InstanceEdgeParam `json:"body,omitempty"`
}

func (o CreateInstanceBy3rdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceBy3rdRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceBy3rdRequest", string(data)}, " ")
}
