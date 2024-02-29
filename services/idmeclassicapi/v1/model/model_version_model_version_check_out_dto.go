package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelVersionCheckOutDto struct {

	// 创建人。
	Creator *string `json:"creator,omitempty"`

	// 关系实体名称集合，与workCopyType的值CUSTOM配合使用。
	CustomLinkSet *[]string `json:"customLinkSet,omitempty"`

	// 主对象ID。
	MasterId int64 `json:"masterId"`

	// 更新者。
	Modifier *string `json:"modifier,omitempty"`

	// 关系COPY类型。 - BOTH:以其为源或目标的均需要复制。 - CUSTOM:自定义复制。 - NONE:不复制。 - SOURCE:仅复制以此为源的。 - TARGET:仅复制以此为目标的。
	WorkCopyType *string `json:"workCopyType,omitempty"`
}

func (o VersionModelVersionCheckOutDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelVersionCheckOutDto struct{}"
	}

	return strings.Join([]string{"VersionModelVersionCheckOutDto", string(data)}, " ")
}
