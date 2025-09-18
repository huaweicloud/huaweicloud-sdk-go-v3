package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Area 支持的大区。
type Area struct {

	// 实例名称。
	Name string `json:"name"`

	// 大区ID。
	Id *string `json:"id,omitempty"`

	// 大区英文名称。
	EnName *string `json:"en_name,omitempty"`

	// 大区西语名称。
	EsName *string `json:"es_name,omitempty"`

	// 大区葡语名称。
	PtName *string `json:"pt_name,omitempty"`

	// 站点。
	Station *string `json:"station,omitempty"`
}

func (o Area) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Area struct{}"
	}

	return strings.Join([]string{"Area", string(data)}, " ")
}
