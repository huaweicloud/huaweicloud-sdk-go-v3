package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStudyReq 创建study请求体
type CreateStudyReq struct {

	// study名称
	Name string `json:"name"`

	// study描述
	Description *string `json:"description,omitempty"`
}

func (o CreateStudyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStudyReq struct{}"
	}

	return strings.Join([]string{"CreateStudyReq", string(data)}, " ")
}
