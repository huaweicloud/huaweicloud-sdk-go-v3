package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListGlobalConnectionBandwidthSitesRequest Request Object
type ListGlobalConnectionBandwidthSitesRequest struct {

	// 每页返回的个数。 取值范围：1~1000。
	Limit *int32 `json:"limit,omitempty"`

	// 翻页信息，从上次API调用返回的翻页数据中获取，可填写前一页marker或者后一页marker，填入前一页previous_marker就向前翻页，后一页next_marker就向翻页。 翻页过程中，查询条件不能修改，包括过滤条件，排序条件，limit。
	Marker *string `json:"marker,omitempty"`

	// 根据id查询，可查询多个id。
	Id *[]string `json:"id,omitempty"`

	// 站点信息自定义英文名称。
	NameEn *string `json:"name_en,omitempty"`

	// 站点信息自定义中文名称。
	NameCn *string `json:"name_cn,omitempty"`

	// 站点编码。
	SiteCode *string `json:"site_code,omitempty"`

	// 站点类型： - Area: 大区 - SubArea: 区域 - Region: 城域
	SiteType *ListGlobalConnectionBandwidthSitesRequestSiteType `json:"site_type,omitempty"`
}

func (o ListGlobalConnectionBandwidthSitesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalConnectionBandwidthSitesRequest struct{}"
	}

	return strings.Join([]string{"ListGlobalConnectionBandwidthSitesRequest", string(data)}, " ")
}

type ListGlobalConnectionBandwidthSitesRequestSiteType struct {
	value string
}

type ListGlobalConnectionBandwidthSitesRequestSiteTypeEnum struct {
	AREA     ListGlobalConnectionBandwidthSitesRequestSiteType
	SUB_AREA ListGlobalConnectionBandwidthSitesRequestSiteType
	REGION   ListGlobalConnectionBandwidthSitesRequestSiteType
}

func GetListGlobalConnectionBandwidthSitesRequestSiteTypeEnum() ListGlobalConnectionBandwidthSitesRequestSiteTypeEnum {
	return ListGlobalConnectionBandwidthSitesRequestSiteTypeEnum{
		AREA: ListGlobalConnectionBandwidthSitesRequestSiteType{
			value: "Area",
		},
		SUB_AREA: ListGlobalConnectionBandwidthSitesRequestSiteType{
			value: "SubArea",
		},
		REGION: ListGlobalConnectionBandwidthSitesRequestSiteType{
			value: "Region",
		},
	}
}

func (c ListGlobalConnectionBandwidthSitesRequestSiteType) Value() string {
	return c.value
}

func (c ListGlobalConnectionBandwidthSitesRequestSiteType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalConnectionBandwidthSitesRequestSiteType) UnmarshalJSON(b []byte) error {
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
