package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterClusterGroupRequest Request Object
type RegisterClusterGroupRequest struct {
	Body *RegisterClusterGroup `json:"body,omitempty" type:"multipart"`
}

func (o RegisterClusterGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterClusterGroupRequest struct{}"
	}

	return strings.Join([]string{"RegisterClusterGroupRequest", string(data)}, " ")
}
