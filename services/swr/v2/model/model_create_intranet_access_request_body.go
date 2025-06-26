package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateIntranetAccessRequestBody struct {

	// vpc和子网所在项目的编号
	ProjectId string `json:"project_id"`

	// vpc编号ID
	VpcId string `json:"vpc_id"`

	// 子网编号ID
	SubnetId string `json:"subnet_id"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o CreateIntranetAccessRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIntranetAccessRequestBody struct{}"
	}

	return strings.Join([]string{"CreateIntranetAccessRequestBody", string(data)}, " ")
}
