package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublicPoolType struct {

	// - 功能说明：弹性公网IP池类型id
	Id *string `json:"id,omitempty"`

	// - 功能说明：弹性公网IP池类型
	Type *string `json:"type,omitempty"`
}

func (o PublicPoolType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicPoolType struct{}"
	}

	return strings.Join([]string{"PublicPoolType", string(data)}, " ")
}
