package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShortGlobalEipSegment struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 资源名称
	Name *string `json:"name,omitempty"`
}

func (o ShortGlobalEipSegment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShortGlobalEipSegment struct{}"
	}

	return strings.Join([]string{"ShortGlobalEipSegment", string(data)}, " ")
}
