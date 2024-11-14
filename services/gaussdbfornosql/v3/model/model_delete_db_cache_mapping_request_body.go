package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteDbCacheMappingRequestBody struct {

	// 内存加速映射ID。
	Id string `json:"id"`
}

func (o DeleteDbCacheMappingRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDbCacheMappingRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteDbCacheMappingRequestBody", string(data)}, " ")
}
