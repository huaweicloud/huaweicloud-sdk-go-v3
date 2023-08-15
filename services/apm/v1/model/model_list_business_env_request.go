package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBusinessEnvRequest Request Object
type ListBusinessEnvRequest struct {

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`

	Body *BusinessEnvRequest `json:"body,omitempty"`
}

func (o ListBusinessEnvRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBusinessEnvRequest struct{}"
	}

	return strings.Join([]string{"ListBusinessEnvRequest", string(data)}, " ")
}
