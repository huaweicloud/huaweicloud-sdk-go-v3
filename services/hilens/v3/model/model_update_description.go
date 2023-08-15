package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDescription struct {

	// 工作空间描述
	Description string `json:"description"`
}

func (o UpdateDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDescription struct{}"
	}

	return strings.Join([]string{"UpdateDescription", string(data)}, " ")
}
