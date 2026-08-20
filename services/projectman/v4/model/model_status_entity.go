package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StatusEntity 工作项状态对象
type StatusEntity struct {

	// 状态ID。
	Id *string `json:"id,omitempty"`

	// 工作项的状态属性。
	Belonging *string `json:"belonging,omitempty"`

	// 状态名。
	DisplayValue *string `json:"display_value,omitempty"`

	// 状态唯一标识。
	Code *string `json:"code,omitempty"`

	// 状态创建人。
	CreatedBy *string `json:"created_by,omitempty"`

	// 状态创建时间。
	CreatedTime *string `json:"created_time,omitempty"`

	// 状态修改人。
	ModifiedBy *string `json:"modified_by,omitempty"`

	// 状态最近修改时间。
	ModifiedTime *string `json:"modified_time,omitempty"`

	// 状态被哪些工作项使用。
	CategoryCode *[]string `json:"category_code,omitempty"`
}

func (o StatusEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatusEntity struct{}"
	}

	return strings.Join([]string{"StatusEntity", string(data)}, " ")
}
