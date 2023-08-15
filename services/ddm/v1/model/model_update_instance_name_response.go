package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceNameResponse Response Object
type UpdateInstanceNameResponse struct {

	// DDM实例名称
	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateInstanceNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceNameResponse", string(data)}, " ")
}
