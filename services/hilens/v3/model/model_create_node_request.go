package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNodeRequest Request Object
type CreateNodeRequest struct {

	// 服务提供者：ief或hilens，选择设备纳管到不同的平台。不填默认为hilens平台
	Provider *string `json:"provider,omitempty"`

	Body *NodeRequest `json:"body,omitempty"`
}

func (o CreateNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNodeRequest struct{}"
	}

	return strings.Join([]string{"CreateNodeRequest", string(data)}, " ")
}
