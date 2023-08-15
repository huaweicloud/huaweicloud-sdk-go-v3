package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IoAccInfoDto struct {

	// io加速实例id
	Id *string `json:"id,omitempty"`

	// io加速实例类型
	Type *string `json:"type,omitempty"`

	// io加速实例总容量
	Space *int32 `json:"space,omitempty"`

	// io加速实例空闲容量
	FreeSpace *float64 `json:"free_space,omitempty"`
}

func (o IoAccInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IoAccInfoDto struct{}"
	}

	return strings.Join([]string{"IoAccInfoDto", string(data)}, " ")
}
