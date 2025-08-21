package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCesInstanceRequest Request Object
type ListCesInstanceRequest struct {
	Body *ListCesInstanceRequestBody `json:"body,omitempty"`
}

func (o ListCesInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCesInstanceRequest struct{}"
	}

	return strings.Join([]string{"ListCesInstanceRequest", string(data)}, " ")
}
