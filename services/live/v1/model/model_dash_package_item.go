package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DashPackageItem DASH频道出流信息
type DashPackageItem struct {

	// 客户自定义的拉流地址，包括方法、域名、路径
	Url string `json:"url"`

	// 从全量流中过滤出一个码率在[min, max]区间的流。如果不需要码率过滤可不选。
	StreamSelection *[]StreamSelectionItem `json:"stream_selection,omitempty"`

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

	// 广告标识。  DASH取值：\"xml+bin\"。
	AdMarker *DashPackageItemAdMarker `json:"ad_marker,omitempty"`
}

func (o DashPackageItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DashPackageItem struct{}"
	}

	return strings.Join([]string{"DashPackageItem", string(data)}, " ")
}

type DashPackageItemAdMarker struct {
	value string
}

type DashPackageItemAdMarkerEnum struct {
	XMLBIN DashPackageItemAdMarker
	XML    DashPackageItemAdMarker
}

func GetDashPackageItemAdMarkerEnum() DashPackageItemAdMarkerEnum {
	return DashPackageItemAdMarkerEnum{
		XMLBIN: DashPackageItemAdMarker{
			value: "xml+bin",
		},
		XML: DashPackageItemAdMarker{
			value: "xml",
		},
	}
}

func (c DashPackageItemAdMarker) Value() string {
	return c.value
}

func (c DashPackageItemAdMarker) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DashPackageItemAdMarker) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
