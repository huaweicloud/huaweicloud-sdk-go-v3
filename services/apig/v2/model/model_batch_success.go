package model

import (
	"encoding/json"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"

	"strings"
)

type BatchSuccess struct {
	// 发布记录的ID

	PublishId *string `json:"publish_id,omitempty"`
	// 发布或下线成功的API ID

	ApiId *string `json:"api_id,omitempty"`
	// 发布或下线成功的APi名称

	ApiName *string `json:"api_name,omitempty"`
	// 发布环境的ID

	EnvId *string `json:"env_id,omitempty"`
	// 发布描述信息

	Remark *string `json:"remark,omitempty"`
	// 发布时间

	PublishTime *sdktime.SdkTime `json:"publish_time,omitempty"`
	// 版本号

	VersionId *string `json:"version_id,omitempty"`
}

func (o BatchSuccess) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchSuccess struct{}"
	}

	return strings.Join([]string{"BatchSuccess", string(data)}, " ")
}
