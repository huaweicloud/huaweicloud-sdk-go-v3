package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FaiureResource struct {

	// 部署失败的资源信息,部署WAF与ELB时，此字段为资源ID，部署CDN时，本字段为加速域名。
	Resource *string `json:"resource,omitempty"`

	// 失败原因，一般为目标服务返回的错误码信息。
	FailureInfo *string `json:"failure_info,omitempty"`
}

func (o FaiureResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FaiureResource struct{}"
	}

	return strings.Join([]string{"FaiureResource", string(data)}, " ")
}
