package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDeploymentResponse Response Object
type DeleteDeploymentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDeploymentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeploymentResponse struct{}"
	}

	return strings.Join([]string{"DeleteDeploymentResponse", string(data)}, " ")
}
