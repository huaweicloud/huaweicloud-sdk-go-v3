package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDbCacheMappingRequestBody struct {

	// 内存加速源实例ID。当前支持云数据库RDS for MySQL和GaussDB(for MySQL)实例。
	SourceInstanceId string `json:"source_instance_id"`

	// 内存加速目标实例ID。当前仅支持GeminiDB Redis主备类型实例。
	TargetInstanceId string `json:"target_instance_id"`
}

func (o CreateDbCacheMappingRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDbCacheMappingRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDbCacheMappingRequestBody", string(data)}, " ")
}
