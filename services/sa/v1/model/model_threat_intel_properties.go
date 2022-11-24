package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ThreatIntelProperties struct {

	// 恶意软件Md5。
	FileMd5 *string `json:"file_md5,omitempty"`

	// 恶意软件Sha1。
	FileSha1 *string `json:"file_sha1,omitempty"`

	// 恶意软件Sha256值。
	FileSha256 *string `json:"file_sha256,omitempty"`

	// 文件名称。
	FileName *string `json:"file_name,omitempty"`

	// 编译时间，格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。时区信息为事件发生时区，无法解析时区的时间，默认时区填东八区。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 文件类别，TEXT|XCODE。
	FileClass *string `json:"file_class,omitempty"`

	// 家族，例如：wannacry（勒索软件）。
	FileFamily *string `json:"file_family,omitempty"`

	// 类别，例如：trojan（特洛伊）。
	FileMaltype *string `json:"file_maltype,omitempty"`

	// mac地址。
	IpResolvesToRefs *string `json:"ip_resolves_to_refs,omitempty"`

	// IP AS 自治系统。
	BelongsToRefs *string `json:"belongs_to_refs,omitempty"`

	// 地区 格式：country/provice/city/lngwgs/latwgs。
	IpLocation *string `json:"ip_location,omitempty"`

	// 例如：banjori|iodine。
	DomainFamily *string `json:"domain_family,omitempty"`

	// 解析的IP地址。
	DomainResolvesToRefs *string `json:"domain_resolves_to_refs,omitempty"`

	// DNS类别。A|NS|CNAME|TXT。
	DomainDnsType *string `json:"domain_dns_type,omitempty"`

	// 例：3ms.huawei.com。
	UrlHost *string `json:"url_host,omitempty"`

	// IP地址。
	UrlResolvesToRefs *string `json:"url_resolves_to_refs,omitempty"`

	// 显示名称。
	DisplayName *string `json:"display_name,omitempty"`

	// 邮箱账户，@之前部分。
	UrlBelongsToRef *string `json:"url_belongs_to_ref,omitempty"`
}

func (o ThreatIntelProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThreatIntelProperties struct{}"
	}

	return strings.Join([]string{"ThreatIntelProperties", string(data)}, " ")
}
