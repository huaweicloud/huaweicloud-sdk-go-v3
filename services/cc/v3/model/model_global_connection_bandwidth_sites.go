package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// GlobalConnectionBandwidthSites 站点信息。
type GlobalConnectionBandwidthSites struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 实例创建时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 实例更新时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 功能说明：站点信息自定义的英文名字。 取值范围：1-255个字符
	NameEn *string `json:"name_en,omitempty"`

	// 功能说明：站点信息自定义的中文名字。 取值范围：1-64个字符。
	NameCn *string `json:"name_cn,omitempty"`

	// 功能说明：站点编码，格式为<area_code>[-<subarea_code>[-<region_code>]]。 取值范围：1-64个字符。
	SiteCode *string `json:"site_code,omitempty"`

	// 功能说明：站点类型，必须跟站点编码对应，一段编码为大区，两段编码为区域，三段编码为城域。 取值范围： - Area: 大区站点 - SubArea: 区域站点 - Region: 城域站点
	SiteType *GlobalConnectionBandwidthSitesSiteType `json:"site_type,omitempty"`

	// 功能说明：站点支持的服务列表，多个服务用\",\"分隔。 取值范围：0-255个字符
	ServiceList *string `json:"service_list,omitempty"`

	GroupList *[]SiteGroupReferenceInfo `json:"group_list,omitempty"`

	// 功能说明：对应华为云标准region的id，该站点继承自华为云region时才需要填写该字段。 取值范围：0-64个字符。
	RegionId *string `json:"region_id,omitempty"`

	// 功能说明：用于标记是中心还是边缘站点。中心：center 取值范围：0-255个字符。
	PublicBorderGroup *string `json:"public_border_group,omitempty"`
}

func (o GlobalConnectionBandwidthSites) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlobalConnectionBandwidthSites struct{}"
	}

	return strings.Join([]string{"GlobalConnectionBandwidthSites", string(data)}, " ")
}

type GlobalConnectionBandwidthSitesSiteType struct {
	value string
}

type GlobalConnectionBandwidthSitesSiteTypeEnum struct {
	AREA     GlobalConnectionBandwidthSitesSiteType
	SUB_AREA GlobalConnectionBandwidthSitesSiteType
	REGION   GlobalConnectionBandwidthSitesSiteType
}

func GetGlobalConnectionBandwidthSitesSiteTypeEnum() GlobalConnectionBandwidthSitesSiteTypeEnum {
	return GlobalConnectionBandwidthSitesSiteTypeEnum{
		AREA: GlobalConnectionBandwidthSitesSiteType{
			value: "Area",
		},
		SUB_AREA: GlobalConnectionBandwidthSitesSiteType{
			value: "SubArea",
		},
		REGION: GlobalConnectionBandwidthSitesSiteType{
			value: "Region",
		},
	}
}

func (c GlobalConnectionBandwidthSitesSiteType) Value() string {
	return c.value
}

func (c GlobalConnectionBandwidthSitesSiteType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlobalConnectionBandwidthSitesSiteType) UnmarshalJSON(b []byte) error {
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
