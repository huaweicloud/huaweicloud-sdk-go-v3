package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Metadata struct {

	// 属性信息。
	Annotations map[string]string `json:"annotations,omitempty"`

	// 创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 组件id。
	Id *string `json:"id,omitempty"`

	// 任务id。
	JodId *string `json:"jod_id,omitempty"`

	// 名称。
	Name string `json:"name"`

	// 状态。
	Status *string `json:"status,omitempty"`

	// 组件类型。
	Type *string `json:"type,omitempty"`

	// 更新时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o Metadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Metadata struct{}"
	}

	return strings.Join([]string{"Metadata", string(data)}, " ")
}
