package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProxyUnitedWorkloadResponse Response Object
type UpdateProxyUnitedWorkloadResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateProxyUnitedWorkloadResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProxyUnitedWorkloadResponse struct{}"
	}

	return strings.Join([]string{"UpdateProxyUnitedWorkloadResponse", string(data)}, " ")
}
