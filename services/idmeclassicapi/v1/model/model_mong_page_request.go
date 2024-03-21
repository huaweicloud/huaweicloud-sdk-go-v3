package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MongPageRequest struct {

	// 查询条件：最后修改时间，结束时间范围。
	EndLastModifiedTime *string `json:"endLastModifiedTime,omitempty"`

	// 唯一标识。
	Id string `json:"id"`

	// 版本号。
	RdmVersion *int32 `json:"rdmVersion,omitempty"`

	// 关系实体源端ID。
	SourceId *string `json:"sourceId,omitempty"`

	// 关系实体源端系统版本。
	SourceRdmVersion *int32 `json:"sourceRdmVersion,omitempty"`

	// 查询条件：最后修改时间,开始时间范围。
	StartLastModifiedTime *string `json:"startLastModifiedTime,omitempty"`

	// 关系实体目标端ID。
	TargetId *string `json:"targetId,omitempty"`

	// 关系实体目标端系统版本。
	TargetRdmVersion *int32 `json:"targetRdmVersion,omitempty"`

	// 单边不确定关系的目标端类型。
	TargetType *string `json:"targetType,omitempty"`
}

func (o MongPageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MongPageRequest struct{}"
	}

	return strings.Join([]string{"MongPageRequest", string(data)}, " ")
}
