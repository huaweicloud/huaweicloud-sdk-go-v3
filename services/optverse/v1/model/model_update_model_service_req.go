package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateModelServiceReq 编辑推理服务请求体
type UpdateModelServiceReq struct {

	// 服务名称
	ServiceName string `json:"service_name"`

	// 服务描述
	ServiceDesc string `json:"service_desc"`
}

func (o UpdateModelServiceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateModelServiceReq struct{}"
	}

	return strings.Join([]string{"UpdateModelServiceReq", string(data)}, " ")
}
