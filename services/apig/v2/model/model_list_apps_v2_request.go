package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAppsV2Request Request Object
type ListAppsV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量，条目数量小于等于0时，自动转换为20，条目数量大于500时，自动转换为500
	Limit *int32 `json:"limit,omitempty"`

	// APP编号
	Id *string `json:"id,omitempty"`

	// APP名称
	Name *string `json:"name,omitempty"`

	// APP状态。 - 1：有效
	Status *ListAppsV2RequestStatus `json:"status,omitempty"`

	// APP凭据的key。
	AppKey *string `json:"app_key,omitempty"`

	// APP的创建者。 - USER：用户自行创建 - MARKET：[云商店分配](tag:hws)[暂未使用](tag:cmcc,ctc,DT,g42,hk_g42,hk_sbc,hk_tm,hws_eu,hws_ocb,OCB,sbc,tm,hws_hk,srg,ax)
	Creator *string `json:"creator,omitempty"`

	// 指定需要精确匹配查找的参数名称，目前仅支持name
	PreciseSearch *string `json:"precise_search,omitempty"`

	// 凭据关联的账号ID。
	RelatedDomainId *string `json:"related_domain_id,omitempty"`

	// 凭据关联的项目ID。
	RelatedProjectId *string `json:"related_project_id,omitempty"`
}

func (o ListAppsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppsV2Request struct{}"
	}

	return strings.Join([]string{"ListAppsV2Request", string(data)}, " ")
}

type ListAppsV2RequestStatus struct {
	value int32
}

type ListAppsV2RequestStatusEnum struct {
	E_1 ListAppsV2RequestStatus
}

func GetListAppsV2RequestStatusEnum() ListAppsV2RequestStatusEnum {
	return ListAppsV2RequestStatusEnum{
		E_1: ListAppsV2RequestStatus{
			value: 1,
		},
	}
}

func (c ListAppsV2RequestStatus) Value() int32 {
	return c.value
}

func (c ListAppsV2RequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAppsV2RequestStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
