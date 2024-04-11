package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MongPageRequest struct {

	// 结束时间。系统以数据实例的最后修改时间作为查询条件，您定义的开始时间和结束时间作为时间范围进行查询。
	EndLastModifiedTime *string `json:"endLastModifiedTime,omitempty"`

	// 数据实例ID。
	Id string `json:"id"`

	// 版本号。
	RdmVersion *int32 `json:"rdmVersion,omitempty"`

	// 关系实体源端ID。
	SourceId *string `json:"sourceId,omitempty"`

	// 关系实体源端系统版本。
	SourceRdmVersion *int32 `json:"sourceRdmVersion,omitempty"`

	// 开始时间。系统以数据实例的最后修改时间作为查询条件，您定义的开始时间和结束时间作为时间范围进行查询。
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
