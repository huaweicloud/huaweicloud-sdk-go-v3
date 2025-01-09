package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuerySessionTypesReq struct {

	// 资源规格编码。
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 磁盘类型。
	SessionType *string `json:"session_type,omitempty"`

	// 资源类型字段。
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源所属云服务类型编码。
	CloudServiceType *string `json:"cloud_service_type,omitempty"`
}

func (o QuerySessionTypesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuerySessionTypesReq struct{}"
	}

	return strings.Join([]string{"QuerySessionTypesReq", string(data)}, " ")
}
