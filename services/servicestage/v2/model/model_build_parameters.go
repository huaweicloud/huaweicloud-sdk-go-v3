package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BuildParameters 构建任务的环境变量。
type BuildParameters struct {

	// 编译命令。默认：  1、根目录存在build.sh：./build.sh  2、根据运行系统，示例如下：  Java和Tomcat：mvn clean package  Nodejs: npm build
	BuildCmd *string `json:"build_cmd,omitempty"`

	// dockerfile地址。默认是根目录./。
	DockerfilePath *string `json:"dockerfile_path,omitempty"`

	// 构建归档组织，默认cas_{project_id}。
	ArtifactNamespace *string `json:"artifact_namespace,omitempty"`

	// 指定构建集群的id。
	ClusterId *string `json:"cluster_id,omitempty"`

	// 指定构建集群的id。
	NodeLabelSelector *interface{} `json:"node_label_selector,omitempty"`
}

func (o BuildParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BuildParameters struct{}"
	}

	return strings.Join([]string{"BuildParameters", string(data)}, " ")
}
