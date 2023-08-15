package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteComponentRequest Request Object
type DeleteComponentRequest struct {

	// 组件id
	ComponentId string `json:"component_id"`
}

func (o DeleteComponentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteComponentRequest struct{}"
	}

	return strings.Join([]string{"DeleteComponentRequest", string(data)}, " ")
}
