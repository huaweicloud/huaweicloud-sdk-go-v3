package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowConfigMapResponse struct {
	Configmap      *ConfigMapResp `json:"configmap,omitempty" xml:"configmap"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowConfigMapResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigMapResponse struct{}"
	}

	return strings.Join([]string{"ShowConfigMapResponse", string(data)}, " ")
}
