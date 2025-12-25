package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowShipperResponseBodyDataShipperSource 投递源
type ShowShipperResponseBodyDataShipperSource struct {

	// 数据类型
	DataType *int32 `json:"data_type,omitempty"`

	// 数据空间
	Dataspace *string `json:"dataspace,omitempty"`

	// 数据空间名称
	DataspaceName *string `json:"dataspace_name,omitempty"`

	// 源ID
	Id *int32 `json:"id,omitempty"`

	// 标识
	Identity *string `json:"identity,omitempty"`

	// 管道ID
	Pipe *string `json:"pipe,omitempty"`

	// 管道名称
	PipeName *string `json:"pipe_name,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	// 类型
	Type *int32 `json:"type,omitempty"`

	// 工作空间ID
	Workspace *string `json:"workspace,omitempty"`

	// 工作空间名称
	WorkspaceName *string `json:"workspace_name,omitempty"`
}

func (o ShowShipperResponseBodyDataShipperSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShipperResponseBodyDataShipperSource struct{}"
	}

	return strings.Join([]string{"ShowShipperResponseBodyDataShipperSource", string(data)}, " ")
}
