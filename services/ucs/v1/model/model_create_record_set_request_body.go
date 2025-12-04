package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRecordSetRequestBody struct {

	// 域名，后缀需以zone name结束且为FQDN（即以“.”号结束的完整主机名）
	Name string `json:"name"`

	// 解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位
	Ttl int32 `json:"ttl"`

	// 解析记录的值，不同类型解析记录对应的值的规则不同
	Records *[]string `json:"records,omitempty"`

	// 解析线路ID
	Line string `json:"line"`

	// 解析记录的权重
	Weight *int32 `json:"weight,omitempty"`

	// Record Set的类型， 取值范围：A、CNAME
	Type *string `json:"type,omitempty"`
}

func (o CreateRecordSetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordSetRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRecordSetRequestBody", string(data)}, " ")
}
