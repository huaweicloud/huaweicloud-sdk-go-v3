package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceInstanceId Service Instance的ID
type ServiceInstanceId struct {
}

func (o ServiceInstanceId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceInstanceId struct{}"
	}

	return strings.Join([]string{"ServiceInstanceId", string(data)}, " ")
}
