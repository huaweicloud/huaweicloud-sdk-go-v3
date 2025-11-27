package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type KustomizationSpec struct {

	// kustomization.yaml文件的路径
	Path *string `json:"path,omitempty"`

	// 用于指定控制器执行 Kustomization同步与校验的时间间隔
	Interval *string `json:"interval,omitempty"`

	// 用于定义验证、应用和健康检查操作的超时
	Timeout *string `json:"timeout,omitempty"`

	SourceRef *SourceRef `json:"sourceRef,omitempty"`

	// 用于设置或覆盖kustomization.yaml文件中的命名空间
	TargetNamespace *string `json:"targetNamespace,omitempty"`

	// 是否启用垃圾回收功能
	Prune *bool `json:"prune,omitempty"`
}

func (o KustomizationSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KustomizationSpec struct{}"
	}

	return strings.Join([]string{"KustomizationSpec", string(data)}, " ")
}
