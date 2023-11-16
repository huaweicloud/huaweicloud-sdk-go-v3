package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateFactoryJobNameRequestBody struct {

	// 作业名称
	Name string `json:"name"`
}

func (o UpdateFactoryJobNameRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFactoryJobNameRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFactoryJobNameRequestBody", string(data)}, " ")
}
