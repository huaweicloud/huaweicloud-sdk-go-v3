package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagNameAndIdVo struct {

	// 标签Id
	TagId *string `json:"tag_id,omitempty"`

	// 标签名称
	TagName *string `json:"tag_name,omitempty"`
}

func (o TagNameAndIdVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagNameAndIdVo struct{}"
	}

	return strings.Join([]string{"TagNameAndIdVo", string(data)}, " ")
}
