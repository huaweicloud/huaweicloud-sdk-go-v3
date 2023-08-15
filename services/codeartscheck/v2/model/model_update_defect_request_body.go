package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDefectRequestBody struct {

	// 问题id,多个时英文逗号分隔
	DefectId *string `json:"defect_id,omitempty"`

	// 状态2：已忽略 1：已解决 0：未解决
	DefectStatus *string `json:"defect_status,omitempty"`
}

func (o UpdateDefectRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDefectRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDefectRequestBody", string(data)}, " ")
}
