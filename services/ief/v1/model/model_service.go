package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Service 服务详情
type Service struct {
	Service *ServiceReqDetail `json:"service,omitempty"`
}

func (o Service) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Service struct{}"
	}

	return strings.Join([]string{"Service", string(data)}, " ")
}
