package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SiteInfo 站点信息。
type SiteInfo struct {

	// 站点id。
	SiteId *string `json:"site_id,omitempty"`

	// 站点名字。
	SiteName *string `json:"site_name,omitempty"`

	// 配置状态。 - CENTER： 中心初始化 - IES： 边缘初始化
	SiteType *SiteInfoSiteType `json:"site_type,omitempty"`

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 站点状态。
	Status *string `json:"status,omitempty"`

	// 创建时间。
	CreateTime *string `json:"create_time,omitempty"`
}

func (o SiteInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SiteInfo struct{}"
	}

	return strings.Join([]string{"SiteInfo", string(data)}, " ")
}

type SiteInfoSiteType struct {
	value string
}

type SiteInfoSiteTypeEnum struct {
	CENTER SiteInfoSiteType
	IES    SiteInfoSiteType
}

func GetSiteInfoSiteTypeEnum() SiteInfoSiteTypeEnum {
	return SiteInfoSiteTypeEnum{
		CENTER: SiteInfoSiteType{
			value: "CENTER",
		},
		IES: SiteInfoSiteType{
			value: "IES",
		},
	}
}

func (c SiteInfoSiteType) Value() string {
	return c.value
}

func (c SiteInfoSiteType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SiteInfoSiteType) UnmarshalJSON(b []byte) error {
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
