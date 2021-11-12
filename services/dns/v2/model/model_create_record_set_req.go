package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRecordSetReq struct {
	// 域名，后缀需以zone name结束且为FQDN（即以“.”号结束的完整主机名）。

	Name string `json:"name"`

	Description *string `json:"description,omitempty"`
	// Record Set的类型。  取值范围：A、AAAA、MX、CNAME、TXT、NS、SRV、CAA。

	Type string `json:"type"`

	Status *string `json:"status,omitempty"`

	Ttl *int32 `json:"ttl,omitempty"`
	// 解析记录的值。不同类型解析记录对应的值的规则不同。

	Records []string `json:"records"`
	// 资源标签。

	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateRecordSetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordSetReq struct{}"
	}

	return strings.Join([]string{"CreateRecordSetReq", string(data)}, " ")
}
