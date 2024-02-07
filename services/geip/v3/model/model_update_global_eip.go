package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateGlobalEip struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 资源名称
	Name *string `json:"name,omitempty"`
}

func (o UpdateGlobalEip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalEip struct{}"
	}

	return strings.Join([]string{"UpdateGlobalEip", string(data)}, " ")
}
