package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LabelRsp struct {

	// 标签id
	Id *string `json:"id,omitempty"`

	// 标签名称
	Name *string `json:"name,omitempty"`

	// 标签描述
	Description *string `json:"description,omitempty"`

	// 标签创建者
	Creator *string `json:"creator,omitempty"`

	// 标签创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 标签更新时间
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o LabelRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LabelRsp struct{}"
	}

	return strings.Join([]string{"LabelRsp", string(data)}, " ")
}
