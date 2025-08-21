package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCesInstanceRspBodyPipeline 通道
type ListCesInstanceRspBodyPipeline struct {

	// pipeline id
	Id *string `json:"id,omitempty"`

	// pipeline name
	Name *string `json:"name,omitempty"`
}

func (o ListCesInstanceRspBodyPipeline) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCesInstanceRspBodyPipeline struct{}"
	}

	return strings.Join([]string{"ListCesInstanceRspBodyPipeline", string(data)}, " ")
}
