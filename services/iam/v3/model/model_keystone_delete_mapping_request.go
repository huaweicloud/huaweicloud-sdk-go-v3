package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type KeystoneDeleteMappingRequest struct {
	// 待删除的映射ID。

	Id string `json:"id"`
}

func (o KeystoneDeleteMappingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneDeleteMappingRequest struct{}"
	}

	return strings.Join([]string{"KeystoneDeleteMappingRequest", string(data)}, " ")
}
