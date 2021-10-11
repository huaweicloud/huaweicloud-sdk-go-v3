package model

import (
	"encoding/json"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"

	"strings"
)

type ApiVersionResp struct {
	// API历史版本的ID

	VersionId *string `json:"version_id,omitempty"`
	// API的版本号

	VersionNo *string `json:"version_no,omitempty"`
	// API编号

	ApiId *string `json:"api_id,omitempty"`
	// 发布的环境编号

	EnvId *string `json:"env_id,omitempty"`
	// 发布的环境名称

	EnvName *string `json:"env_name,omitempty"`
	// 发布描述

	Remark *string `json:"remark,omitempty"`
	// 发布时间

	PublishTime *sdktime.SdkTime `json:"publish_time,omitempty"`
	// 版本状态 - 1：当前生效中的版本 - 2：未生效的版本

	Status *int32 `json:"status,omitempty"`
}

func (o ApiVersionResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ApiVersionResp struct{}"
	}

	return strings.Join([]string{"ApiVersionResp", string(data)}, " ")
}
