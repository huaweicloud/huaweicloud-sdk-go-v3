package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShipperDestination 投递目的
type ShipperDestination struct {

	// 数据参数，通常为JSON格式
	DataParam string `json:"data_param"`

	// 数据类型，具体含义根据业务定义
	DataType int32 `json:"data_type"`

	// 数据空间ID
	Dataspace string `json:"dataspace"`

	// 数据空间名称
	DataspaceName string `json:"dataspace_name"`

	// 目的详细信息
	DestinationInfo string `json:"destination_info"`

	// 唯一标识ID
	Id int32 `json:"id"`

	// 身份标识
	Identity string `json:"identity"`

	// 管道ID
	Pipe string `json:"pipe"`

	// 管道名称
	PipeName string `json:"pipe_name"`

	// 地域信息
	Region string `json:"region"`

	// 类型码，具体含义根据业务定义
	Type int32 `json:"type"`

	// 工作空间ID
	Workspace string `json:"workspace"`

	// 工作空间名称
	WorkspaceName string `json:"workspace_name"`
}

func (o ShipperDestination) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShipperDestination struct{}"
	}

	return strings.Join([]string{"ShipperDestination", string(data)}, " ")
}
