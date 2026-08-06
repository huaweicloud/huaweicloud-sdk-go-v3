package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDbNamesRequest Request Object
type ListDbNamesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 引擎类型
	EngineType *string `json:"engine_type,omitempty"`
}

func (o ListDbNamesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbNamesRequest struct{}"
	}

	return strings.Join([]string{"ListDbNamesRequest", string(data)}, " ")
}
