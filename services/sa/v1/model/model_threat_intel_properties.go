package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ThreatIntelProperties struct {

	// 恶意软件Md5。
	FileMd5 *string `json:"file_md5,omitempty" xml:"file_md5"`

	// 恶意软件Sha1。
	FileSha1 *string `json:"file_sha1,omitempty" xml:"file_sha1"`

	// 恶意软件Sha256值。
	FileSha256 *string `json:"file_sha256,omitempty" xml:"file_sha256"`

	// 文件名称。
	FileName *string `json:"file_name,omitempty" xml:"file_name"`

	// 编译时间，格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。时区信息为事件发生时区，无法解析时区的时间，默认时区填东八区。
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 文件类别，TEXT|XCODE。
	FileClass *string `json:"file_class,omitempty" xml:"file_class"`

	// 家族，例如：wannacry（勒索软件）。
	FileFamily *string `json:"file_family,omitempty" xml:"file_family"`

	// 类别，例如：trojan（特洛伊）。
	FileMaltype *string `json:"file_maltype,omitempty" xml:"file_maltype"`

	// mac地址。
	IpResolvesToRefs *string `json:"ip_resolves_to_refs,omitempty" xml:"ip_resolves_to_refs"`

	// IP AS 自治系统。
	BelongsToRefs *string `json:"belongs_to_refs,omitempty" xml:"belongs_to_refs"`

	// 地区 格式：country/provice/city/lngwgs/latwgs。
	IpLocation *string `json:"ip_location,omitempty" xml:"ip_location"`

	// 例如：banjori|iodine。
	DomainFamily *string `json:"domain_family,omitempty" xml:"domain_family"`

	// 解析的IP地址。
	DomainResolvesToRefs *string `json:"domain_resolves_to_refs,omitempty" xml:"domain_resolves_to_refs"`

	// DNS类别。A|NS|CNAME|TXT。
	DomainDnsType *string `json:"domain_dns_type,omitempty" xml:"domain_dns_type"`

	// 例：3ms.huawei.com。
	UrlHost *string `json:"url_host,omitempty" xml:"url_host"`

	// IP地址。
	UrlResolvesToRefs *string `json:"url_resolves_to_refs,omitempty" xml:"url_resolves_to_refs"`

	// 显示名称。
	DisplayName *string `json:"display_name,omitempty" xml:"display_name"`

	// 邮箱账户，@之前部分。
	UrlBelongsToRef *string `json:"url_belongs_to_ref,omitempty" xml:"url_belongs_to_ref"`
}

func (o ThreatIntelProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThreatIntelProperties struct{}"
	}

	return strings.Join([]string{"ThreatIntelProperties", string(data)}, " ")
}
