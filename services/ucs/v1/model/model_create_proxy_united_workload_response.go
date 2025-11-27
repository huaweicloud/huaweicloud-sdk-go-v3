package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProxyUnitedWorkloadResponse Response Object
type CreateProxyUnitedWorkloadResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateProxyUnitedWorkloadResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProxyUnitedWorkloadResponse struct{}"
	}

	return strings.Join([]string{"CreateProxyUnitedWorkloadResponse", string(data)}, " ")
}
