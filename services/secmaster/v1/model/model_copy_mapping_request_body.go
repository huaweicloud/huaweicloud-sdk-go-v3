package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CopyMappingRequestBody struct {

	// 名称
	Name string `json:"name"`
}

func (o CopyMappingRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyMappingRequestBody struct{}"
	}

	return strings.Join([]string{"CopyMappingRequestBody", string(data)}, " ")
}
