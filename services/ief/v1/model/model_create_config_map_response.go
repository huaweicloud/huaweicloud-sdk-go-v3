package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConfigMapResponse Response Object
type CreateConfigMapResponse struct {
	Configmap      *ConfigMapResp `json:"configmap,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateConfigMapResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigMapResponse struct{}"
	}

	return strings.Join([]string{"CreateConfigMapResponse", string(data)}, " ")
}
