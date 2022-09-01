package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ThreatIntel struct {

	// 情报Id。
	Id string `json:"id" xml:"id"`

	// 威胁情报类型，Domain、Email_Address、Hash_MD5、Hash_SHA1、Hash_SHA256、 Hash_SHA512、IPv4_Address、IPv6_Address、URL。
	IndicatorType *string `json:"indicator_type,omitempty" xml:"indicator_type"`

	// 标签，如'矿池','外联'等，\"Directory Scan|Directory Traversal\"。
	Labels *string `json:"labels,omitempty" xml:"labels"`

	// 置信度，不同来源目前置信度分值定义不一样（分数）。
	Confidence *int32 `json:"confidence,omitempty" xml:"confidence"`

	// 威胁情报源，最大64个字符。
	InformationSource string `json:"information_source" xml:"information_source"`

	// 严重程度，不同渠道定义值不一样（分数）。
	Severity *int32 `json:"severity,omitempty" xml:"severity"`

	// 威胁情报描述。
	Description string `json:"description" xml:"description"`

	// 威胁情报的更新时间，格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。时区信息为事件发生时区，无法解析时区的时间，默认时区填东八区。
	Modified *string `json:"modified,omitempty" xml:"modified"`

	// 有效期开始（可读字符串）。
	ValidFrom *string `json:"valid_from,omitempty" xml:"valid_from"`

	// 有效期结束（可读字符串）。
	ValidUntil *string `json:"valid_until,omitempty" xml:"valid_until"`

	Properties *ThreatIntelProperties `json:"properties,omitempty" xml:"properties"`
}

func (o ThreatIntel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThreatIntel struct{}"
	}

	return strings.Join([]string{"ThreatIntel", string(data)}, " ")
}
