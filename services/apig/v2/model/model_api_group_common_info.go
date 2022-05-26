package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ApiGroupCommonInfo struct {

	// 编号
	Id string `json:"id"`

	// API分组名称
	Name string `json:"name"`

	// 状态   - 1： 有效
	Status ApiGroupCommonInfoStatus `json:"status"`

	// 系统默认分配的子域名
	SlDomain string `json:"sl_domain"`

	// 创建时间
	RegisterTime *sdktime.SdkTime `json:"register_time"`

	// 最近修改时间
	UpdateTime *sdktime.SdkTime `json:"update_time"`

	// 是否已上架云市场： - 1：已上架 - 2：未上架 - 3：审核中
	OnSellStatus int32 `json:"on_sell_status"`

	// 分组上绑定的独立域名列表
	UrlDomains *[]UrlDomain `json:"url_domains,omitempty"`
}

func (o ApiGroupCommonInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiGroupCommonInfo struct{}"
	}

	return strings.Join([]string{"ApiGroupCommonInfo", string(data)}, " ")
}

type ApiGroupCommonInfoStatus struct {
	value int32
}

type ApiGroupCommonInfoStatusEnum struct {
	E_1 ApiGroupCommonInfoStatus
}

func GetApiGroupCommonInfoStatusEnum() ApiGroupCommonInfoStatusEnum {
	return ApiGroupCommonInfoStatusEnum{
		E_1: ApiGroupCommonInfoStatus{
			value: 1,
		},
	}
}

func (c ApiGroupCommonInfoStatus) Value() int32 {
	return c.value
}

func (c ApiGroupCommonInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiGroupCommonInfoStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
