package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UninstallResponse Response Object
type UninstallResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UninstallResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UninstallResponse struct{}"
	}

	return strings.Join([]string{"UninstallResponse", string(data)}, " ")
}
