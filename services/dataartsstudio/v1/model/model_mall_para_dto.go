package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MallParaDto 查询服务目录API列表请求参数
type MallParaDto struct {

	// 认证类型。
	AuthType *MallParaDtoAuthType `json:"auth_type,omitempty"`

	// API可见性，WORKSPACE：工作空间可见，PROJECT： 项目可见，DOMAIN：租户可见，SPECIFIC_PROJECT：指定项目可见。
	Visibility *MallParaDtoVisibility `json:"visibility,omitempty"`

	// 排序参数。
	MarketSortType *MallParaDtoMarketSortType `json:"market_sort_type,omitempty"`

	// 升序、降序。
	AscOrDesc *MallParaDtoAscOrDesc `json:"asc_or_desc,omitempty"`

	// 查询起始坐标。
	Offset *int32 `json:"offset,omitempty"`

	// 查询条数。
	Limit *int32 `json:"limit,omitempty"`

	// 是否显示拥有者。
	IsOwner *bool `json:"is_owner,omitempty"`

	// 是否显示已被授权。
	IsAuthorized *bool `json:"is_authorized,omitempty"`

	// 是否显示最近更新。
	IsUpdateRecently *bool `json:"is_update_recently,omitempty"`

	// 是否显示最近发布。
	IsReleaseRecently *bool `json:"is_release_recently,omitempty"`

	// 是否显示热销状态。
	IsHotRecently *bool `json:"is_hot_recently,omitempty"`

	// 是否显示7天内成功率与失败率。
	SuccessAndFailureRate *bool `json:"success_and_failure_rate,omitempty"`
}

func (o MallParaDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MallParaDto struct{}"
	}

	return strings.Join([]string{"MallParaDto", string(data)}, " ")
}

type MallParaDtoAuthType struct {
	value string
}

type MallParaDtoAuthTypeEnum struct {
	APP  MallParaDtoAuthType
	IAM  MallParaDtoAuthType
	NONE MallParaDtoAuthType
}

func GetMallParaDtoAuthTypeEnum() MallParaDtoAuthTypeEnum {
	return MallParaDtoAuthTypeEnum{
		APP: MallParaDtoAuthType{
			value: "APP",
		},
		IAM: MallParaDtoAuthType{
			value: "IAM",
		},
		NONE: MallParaDtoAuthType{
			value: "NONE",
		},
	}
}

func (c MallParaDtoAuthType) Value() string {
	return c.value
}

func (c MallParaDtoAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MallParaDtoAuthType) UnmarshalJSON(b []byte) error {
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

type MallParaDtoVisibility struct {
	value string
}

type MallParaDtoVisibilityEnum struct {
	WORKSPACE        MallParaDtoVisibility
	PROJECT          MallParaDtoVisibility
	DOMAIN           MallParaDtoVisibility
	SPECIFIC_PROJECT MallParaDtoVisibility
}

func GetMallParaDtoVisibilityEnum() MallParaDtoVisibilityEnum {
	return MallParaDtoVisibilityEnum{
		WORKSPACE: MallParaDtoVisibility{
			value: "WORKSPACE",
		},
		PROJECT: MallParaDtoVisibility{
			value: "PROJECT",
		},
		DOMAIN: MallParaDtoVisibility{
			value: "DOMAIN",
		},
		SPECIFIC_PROJECT: MallParaDtoVisibility{
			value: "SPECIFIC_PROJECT",
		},
	}
}

func (c MallParaDtoVisibility) Value() string {
	return c.value
}

func (c MallParaDtoVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MallParaDtoVisibility) UnmarshalJSON(b []byte) error {
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

type MallParaDtoMarketSortType struct {
	value string
}

type MallParaDtoMarketSortTypeEnum struct {
	UPDATE_TIME MallParaDtoMarketSortType
	CREATE_TIME MallParaDtoMarketSortType
}

func GetMallParaDtoMarketSortTypeEnum() MallParaDtoMarketSortTypeEnum {
	return MallParaDtoMarketSortTypeEnum{
		UPDATE_TIME: MallParaDtoMarketSortType{
			value: "updateTime",
		},
		CREATE_TIME: MallParaDtoMarketSortType{
			value: "createTime",
		},
	}
}

func (c MallParaDtoMarketSortType) Value() string {
	return c.value
}

func (c MallParaDtoMarketSortType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MallParaDtoMarketSortType) UnmarshalJSON(b []byte) error {
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

type MallParaDtoAscOrDesc struct {
	value string
}

type MallParaDtoAscOrDescEnum struct {
	ASC  MallParaDtoAscOrDesc
	DESC MallParaDtoAscOrDesc
}

func GetMallParaDtoAscOrDescEnum() MallParaDtoAscOrDescEnum {
	return MallParaDtoAscOrDescEnum{
		ASC: MallParaDtoAscOrDesc{
			value: "asc",
		},
		DESC: MallParaDtoAscOrDesc{
			value: "desc",
		},
	}
}

func (c MallParaDtoAscOrDesc) Value() string {
	return c.value
}

func (c MallParaDtoAscOrDesc) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MallParaDtoAscOrDesc) UnmarshalJSON(b []byte) error {
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
