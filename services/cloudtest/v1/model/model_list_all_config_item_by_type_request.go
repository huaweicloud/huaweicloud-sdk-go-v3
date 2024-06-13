package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllConfigItemByTypeRequest Request Object
type ListAllConfigItemByTypeRequest struct {

	// 服务id
	ServiceId string `json:"service_id"`

	Body *ListAllConfigItemByTypeRequestBody `json:"body,omitempty"`
}

func (o ListAllConfigItemByTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllConfigItemByTypeRequest struct{}"
	}

	return strings.Join([]string{"ListAllConfigItemByTypeRequest", string(data)}, " ")
}
