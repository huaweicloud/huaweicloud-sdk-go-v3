package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateKubernetesClusterCertResponse struct {

	// API类型，固定值“Config”，该值不可修改。
	Kind *string `json:"kind,omitempty" xml:"kind"`

	// API版本，固定值“v1”。
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion"`

	// 当前未使用该字段，当前默认为空。
	Preferences *interface{} `json:"preferences,omitempty" xml:"preferences"`

	// 集群列表。
	Clusters *[]Clusters `json:"clusters,omitempty" xml:"clusters"`

	// 存放了指定用户的一些证书信息和ClientKey信息。
	Users *[]Users `json:"users,omitempty" xml:"users"`

	// 上下文列表。
	Contexts *[]Contexts `json:"contexts,omitempty" xml:"contexts"`

	// 当前上下文，若存在publicIp（虚拟机弹性IP）时为 external; 若不存在publicIp为 internal。
	CurrentContext *string `json:"current-context,omitempty" xml:"current-context"`

	PortID         *string `json:"Port-ID,omitempty" xml:"Port-ID"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateKubernetesClusterCertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKubernetesClusterCertResponse struct{}"
	}

	return strings.Join([]string{"CreateKubernetesClusterCertResponse", string(data)}, " ")
}
