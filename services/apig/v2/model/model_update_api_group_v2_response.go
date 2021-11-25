package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Response Object
type UpdateApiGroupV2Response struct {
	// 编号

	Id string `json:"id"`
	// API分组名称

	Name string `json:"name"`
	// 状态   - 1： 有效

	Status UpdateApiGroupV2ResponseStatus `json:"status"`
	// 系统默认分配的子域名

	SlDomain string `json:"sl_domain"`
	// 创建时间

	RegisterTime *sdktime.SdkTime `json:"register_time"`
	// 最近修改时间

	UpdateTime *sdktime.SdkTime `json:"update_time"`
	// 是否已上架云市场： - 1：已上架 - 2：未上架 - 3：审核中  ROMAConnect暂未对接云市场，此字段默认返回2

	OnSellStatus int32 `json:"on_sell_status"`
	// 分组上绑定的独立域名列表

	UrlDomains     *[]UrlDomain `json:"url_domains,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateApiGroupV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApiGroupV2Response struct{}"
	}

	return strings.Join([]string{"UpdateApiGroupV2Response", string(data)}, " ")
}

type UpdateApiGroupV2ResponseStatus struct {
	value int32
}

type UpdateApiGroupV2ResponseStatusEnum struct {
	E_1 UpdateApiGroupV2ResponseStatus
}

func GetUpdateApiGroupV2ResponseStatusEnum() UpdateApiGroupV2ResponseStatusEnum {
	return UpdateApiGroupV2ResponseStatusEnum{
		E_1: UpdateApiGroupV2ResponseStatus{
			value: 1,
		},
	}
}

func (c UpdateApiGroupV2ResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateApiGroupV2ResponseStatus) UnmarshalJSON(b []byte) error {
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
