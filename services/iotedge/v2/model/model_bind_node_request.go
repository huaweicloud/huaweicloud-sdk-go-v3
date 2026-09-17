package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindNodeRequest Request Object
type BindNodeRequest struct {

	// 资源id
	ResourceId string `json:"resource_id"`

	Body *AssociateNodeRequestBody `json:"body,omitempty"`
}

func (o BindNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindNodeRequest struct{}"
	}

	return strings.Join([]string{"BindNodeRequest", string(data)}, " ")
}
