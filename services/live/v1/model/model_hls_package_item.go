package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HlsPackageItem HLS频道出流信息
type HlsPackageItem struct {

	// 客户自定义的拉流地址，包括方法、域名、路径
	Url *string `json:"url,omitempty"`

	// 从全量流中过滤出一个码率在[min, max]区间的流。如果不需要码率过滤可不选。
	StreamSelection *[]StreamSelectionItem `json:"stream_selection,omitempty"`

	// HLS版本号
	HlsVersion *string `json:"hls_version,omitempty"`

	// 频道输出分片的时长，为必选项  单位：秒。取值范围：1-10 > 修改分片时长会影响已录制内容的时移和回看服务，请谨慎修改！
	SegmentDurationSeconds int32 `json:"segment_duration_seconds"`

	// 频道直播返回分片的窗口长度，为频道输出分片的时长乘以数量后得到的值。实际返回的分片数不小于3个。  单位：秒。取值范围：0 - 86400（24小时转化成秒后的取值）
	PlaylistWindowSeconds *int32 `json:"playlist_window_seconds,omitempty"`

	Encryption *Encryption `json:"encryption,omitempty"`

	// 广告配置
	Ads *interface{} `json:"ads,omitempty"`

	// 其他额外参数
	ExtArgs *interface{} `json:"ext_args,omitempty"`

	RequestArgs *PackageRequestArgs `json:"request_args,omitempty"`

	// 广告标识。  HLS取值：[\"ENHANCED_SCTE35\"]。
	AdMarker *[]string `json:"ad_marker,omitempty"`

	// 当频道mode是ONLY_OS类型时，允许本输出可以直接从源站拉流，默认：false true：允许output访问 false：禁止output访问
	EnableAccess *bool `json:"enable_access,omitempty"`

	// 是否放通所有的IP访问，默认：false true：允许所有的IP地址访问，ip_whitelist配置不生效 false：不允许所有的IP地址访问，ip_whitelist生效，仅在ip_whitelist配置的ip地址才能访问
	AllowAllIpAccess *bool `json:"allow_all_ip_access,omitempty"`

	// 当频道类型mode是ONLY_OS类型时，允许直接从源站拉流的IP白名单
	IpWhitelist *string `json:"ip_whitelist,omitempty"`

	CdnIdentifierHeader *HttpHeader `json:"cdn_identifier_header,omitempty"`

	// 源站分发域名-主region 跟CreateOttChannelInfoReq.region一致 满足正则：^(\\[a-zA-Z0-9]([a-zA-Z0-9\\-]{0,61}[a-zA-Z0-9])?\\.){2,}[a-zA-Z]{2,16}$ 最大长度255
	OriginDomainMaster *string `json:"origin_domain_master,omitempty"`

	// 源站分发域名-备region 满足正则：^(\\[a-zA-Z0-9]([a-zA-Z0-9\\-]{0,61}[a-zA-Z0-9])?\\.){2,}[a-zA-Z]{2,16}$ 最大长度255
	OriginDomainSlave *string `json:"origin_domain_slave,omitempty"`

	// output的索引文件名字 默认：index 长度：0-128 字符集：大小写字母、数字、\"-\"、\".\"、\"_\"，不能有/路径
	ManifestName *string `json:"manifest_name,omitempty"`

	// 客户自定义的拉流地址，包括方法、域名、路径
	SlaveUrl *string `json:"slave_url,omitempty"`
}

func (o HlsPackageItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HlsPackageItem struct{}"
	}

	return strings.Join([]string{"HlsPackageItem", string(data)}, " ")
}
