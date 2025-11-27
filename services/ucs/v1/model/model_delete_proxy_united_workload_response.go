package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProxyUnitedWorkloadResponse Response Object
type DeleteProxyUnitedWorkloadResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteProxyUnitedWorkloadResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProxyUnitedWorkloadResponse struct{}"
	}

	return strings.Join([]string{"DeleteProxyUnitedWorkloadResponse", string(data)}, " ")
}
