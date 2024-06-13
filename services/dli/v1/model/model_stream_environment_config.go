package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StreamEnvironmentConfig 流作业的上下文配置。
type StreamEnvironmentConfig struct {

	// 自定义委托名，定义作业级访问权限。
	ExecutionAgencyUrn *string `json:"execution_agency_urn,omitempty"`

	// 表示用户作业使用的镜像类型。basic：表示使用 DLI 提供的基础镜像； custom：表示使用用户自定义的镜像。
	ImageFeature *string `json:"image_feature,omitempty"`

	// 自定义镜像。当前只支持 SWR，格式为：组织名/镜像名:镜像版本。当用户设置“image_feature”为“custom”时，该参数生效。用户可通过与“image_feature”参数配合使用，指定作业运行使用自定义的镜像。关于如何使用自定义镜像，请参考《数据湖探索用户指南》。
	ImageUri *string `json:"image_uri,omitempty"`

	// 队列名称。长度限制：1-128个字符。
	QueueName *string `json:"queue_name,omitempty"`
}

func (o StreamEnvironmentConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StreamEnvironmentConfig struct{}"
	}

	return strings.Join([]string{"StreamEnvironmentConfig", string(data)}, " ")
}
