package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Area 支持的大区。
type Area struct {

	// 实例名称。
	Name string `json:"name"`

	// Area ID。
	Id *string `json:"id,omitempty"`

	// 英文 Area Name。
	EnName *string `json:"en_name,omitempty"`

	// 西语 Area Name。
	EsName *string `json:"es_name,omitempty"`

	// 葡语 Area Name。
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
