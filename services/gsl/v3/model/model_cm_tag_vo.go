package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CmTagVo struct {

	// 标签标识
	Id *int64 `json:"id,omitempty"`

	// 标签名称
	TagName *string `json:"tag_name,omitempty"`

	// 标签状态，0未使用，1使用中。
	Status *int32 `json:"status,omitempty"`
}

func (o CmTagVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CmTagVo struct{}"
	}

	return strings.Join([]string{"CmTagVo", string(data)}, " ")
}
