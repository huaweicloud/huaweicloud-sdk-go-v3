package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type GlobalConnectionBandwidthSpecCode struct {

	// 实例ID。
	Id string `json:"id"`

	// 功能说明：本端接入点，配合remote_area信息描述带宽实例应用的范围。 取值范围：1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点），站点编码通过接口获取，带宽类型为Region可不传，其他类型必传
	LocalArea *string `json:"local_area,omitempty"`

	// 功能说明：远端接入点，配合local_area信息描述带宽实例应用的范围。 取值范围：1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点），站点编码通过接口获取，带宽类型为Region可不传，其他类型必传
	RemoteArea *string `json:"remote_area,omitempty"`

	// 实例创建时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 实例更新时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 功能说明：线路规格中文名。
	NameZh *string `json:"name_zh,omitempty"`

	// 功能说明：线路规格英文名。
	NameEn *string `json:"name_en,omitempty"`

	// 支持的线路等级： - Pt: 铂金 - Au: 金 - Ag: 银
	Level *GlobalConnectionBandwidthSpecCodeLevel `json:"level,omitempty"`

	// 功能描述：GCB特定线路规格产品编码。
	Sku *string `json:"sku,omitempty"`

	// 功能描述：带宽起售值，单位Mbps。
	Size *int32 `json:"size,omitempty"`
}

func (o GlobalConnectionBandwidthSpecCode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlobalConnectionBandwidthSpecCode struct{}"
	}

	return strings.Join([]string{"GlobalConnectionBandwidthSpecCode", string(data)}, " ")
}

type GlobalConnectionBandwidthSpecCodeLevel struct {
	value string
}

type GlobalConnectionBandwidthSpecCodeLevelEnum struct {
	PT GlobalConnectionBandwidthSpecCodeLevel
	AU GlobalConnectionBandwidthSpecCodeLevel
	AG GlobalConnectionBandwidthSpecCodeLevel
}

func GetGlobalConnectionBandwidthSpecCodeLevelEnum() GlobalConnectionBandwidthSpecCodeLevelEnum {
	return GlobalConnectionBandwidthSpecCodeLevelEnum{
		PT: GlobalConnectionBandwidthSpecCodeLevel{
			value: "Pt",
		},
		AU: GlobalConnectionBandwidthSpecCodeLevel{
			value: "Au",
		},
		AG: GlobalConnectionBandwidthSpecCodeLevel{
			value: "Ag",
		},
	}
}

func (c GlobalConnectionBandwidthSpecCodeLevel) Value() string {
	return c.value
}

func (c GlobalConnectionBandwidthSpecCodeLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlobalConnectionBandwidthSpecCodeLevel) UnmarshalJSON(b []byte) error {
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
