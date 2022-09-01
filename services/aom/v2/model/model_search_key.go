package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 日志过滤条件集合，不同日志来源所需字段不同。
type SearchKey struct {

	// 应用名称。
	AppName *string `json:"appName,omitempty" xml:"appName"`

	// CCE集群ID。
	ClusterId string `json:"clusterId" xml:"clusterId"`

	// 日志所在虚拟机IP。
	HostIP *string `json:"hostIP,omitempty" xml:"hostIP"`

	// CCE容器集群的命名空间。
	NameSpace *string `json:"nameSpace,omitempty" xml:"nameSpace"`

	// 日志文件名称。
	PathFile *string `json:"pathFile,omitempty" xml:"pathFile"`

	// 容器实例名称。
	PodName *string `json:"podName,omitempty" xml:"podName"`
}

func (o SearchKey) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchKey struct{}"
	}

	return strings.Join([]string{"SearchKey", string(data)}, " ")
}
