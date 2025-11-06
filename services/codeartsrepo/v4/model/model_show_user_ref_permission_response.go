package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserRefPermissionResponse Response Object
type ShowUserRefPermissionResponse struct {
	Read *UserRefPermissionBasicDto `json:"read,omitempty"`

	Review *UserRefPermissionBasicDto `json:"review,omitempty"`

	Approval *UserRefPermissionBasicDto `json:"approval,omitempty"`

	CreateChange *UserRefPermissionBasicDto `json:"create_change,omitempty"`

	Merge *UserRefPermissionBasicDto `json:"merge,omitempty"`

	CreateDelete *UserRefPermissionBasicDto `json:"create_delete,omitempty"`

	Push           *UserRefPermissionBasicDto `json:"push,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowUserRefPermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserRefPermissionResponse struct{}"
	}

	return strings.Join([]string{"ShowUserRefPermissionResponse", string(data)}, " ")
}
