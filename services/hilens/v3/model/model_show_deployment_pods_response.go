package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeploymentPodsResponse Response Object
type ShowDeploymentPodsResponse struct {

	// pod总个数
	Count *int32 `json:"count,omitempty"`

	// pod 列表
	Pods           *[]Pod `json:"pods,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowDeploymentPodsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeploymentPodsResponse struct{}"
	}

	return strings.Join([]string{"ShowDeploymentPodsResponse", string(data)}, " ")
}
