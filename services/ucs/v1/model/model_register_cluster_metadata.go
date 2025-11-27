package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterClusterMetadata 集群元数据信息。
type RegisterClusterMetadata struct {

	// 集群ID信息，仅在注册CCE导入集群时使用，其他类型集群无需填写。
	Uid *string `json:"uid,omitempty"`

	// CCE集群填写CCE集群名称，其他类型集群自定义
	Name string `json:"name"`

	// 标签信息。可为空，不为空时，必须满足kurbenetes label规范,最多可支持100个标签。
	Labels map[string]string `json:"labels,omitempty"`

	// 集群annotations信息。 附着集群必填一个kubeconfig字段，取值是kubeconfig文件的内容。kubeconfig文件获取方法请参见[获取KubeConfig文件](https://support.huaweicloud.com/usermanual-ucs/ucs_01_0138.html)
	Annotations map[string]string `json:"annotations,omitempty"`
}

func (o RegisterClusterMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterClusterMetadata struct{}"
	}

	return strings.Join([]string{"RegisterClusterMetadata", string(data)}, " ")
}
