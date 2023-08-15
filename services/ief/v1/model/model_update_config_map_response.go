package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConfigMapResponse Response Object
type UpdateConfigMapResponse struct {
	Configmap      *ConfigMapResp `json:"configmap,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateConfigMapResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigMapResponse struct{}"
	}

	return strings.Join([]string{"UpdateConfigMapResponse", string(data)}, " ")
}
