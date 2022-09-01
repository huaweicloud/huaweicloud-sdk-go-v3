package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRecordSetWithLineReq struct {

	// 域名，后缀需以zone name结束且为FQDN（即以“.”号结束的完整主机名）。
	Name string `json:"name" xml:"name"`

	// 可选配置，对域名的描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// Record Set的类型。取值范围：A、AAAA、MX、CNAME、TXT、NS、SRV、CAA。
	Type string `json:"type" xml:"type"`

	// 解析记录的状态。默认值为ENABLE。
	Status *string `json:"status,omitempty" xml:"status"`

	// 解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。
	Ttl *int32 `json:"ttl,omitempty" xml:"ttl"`

	// 解析记录的值。不同类型解析记录对应的值的规则不同。
	Records *[]string `json:"records,omitempty" xml:"records"`

	// 解析线路ID
	Line *string `json:"line,omitempty" xml:"line"`

	//  资源标签
	Tags *[]Tag `json:"tags,omitempty" xml:"tags"`

	// 解析记录的权重。
	Weight *int32 `json:"weight,omitempty" xml:"weight"`

	AliasTarget *AliasTarget `json:"alias_target,omitempty" xml:"alias_target"`
}

func (o CreateRecordSetWithLineReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordSetWithLineReq struct{}"
	}

	return strings.Join([]string{"CreateRecordSetWithLineReq", string(data)}, " ")
}
