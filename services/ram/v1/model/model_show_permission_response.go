package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPermissionResponse Response Object
type ShowPermissionResponse struct {
	Permission     *Permission `json:"permission,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowPermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPermissionResponse struct{}"
	}

	return strings.Join([]string{"ShowPermissionResponse", string(data)}, " ")
}
