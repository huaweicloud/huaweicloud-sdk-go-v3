package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UninstallNextflowResponse Response Object
type UninstallNextflowResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UninstallNextflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UninstallNextflowResponse struct{}"
	}

	return strings.Join([]string{"UninstallNextflowResponse", string(data)}, " ")
}
