package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateLabelReq struct {

	// 标签名称
	Name string `json:"name"`

	// 标签描述
	Description *string `json:"description,omitempty"`
}

func (o CreateLabelReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLabelReq struct{}"
	}

	return strings.Join([]string{"CreateLabelReq", string(data)}, " ")
}
