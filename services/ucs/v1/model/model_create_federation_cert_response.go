package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFederationCertResponse Response Object
type CreateFederationCertResponse struct {

	// API类型，固定值“Config”，该值不可修改
	Kind *string `json:"kind,omitempty"`

	// API版本，固定值“v1”，该值不可修改
	ApiVersion *string `json:"apiVersion,omitempty"`

	// 集群列表
	Clusters *[]NamedCluster `json:"clusters,omitempty"`

	// 存放了指定用户的一些证书信息和ClientKey信息
	Users *[]NamedAuthInfo `json:"users,omitempty"`

	// 上下文列表
	Contexts *[]NamedContext `json:"contexts,omitempty"`

	// 当前上下文
	CurrentContext *string `json:"current-context,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateFederationCertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFederationCertResponse struct{}"
	}

	return strings.Join([]string{"CreateFederationCertResponse", string(data)}, " ")
}
