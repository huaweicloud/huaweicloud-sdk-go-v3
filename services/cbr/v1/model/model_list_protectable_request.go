package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListProtectableRequest Request Object
type ListProtectableRequest struct {

	// 每页显示的条目数量，每页最多支持50条
	Limit *int32 `json:"limit,omitempty"`

	// 上一次查询最后一条的ID
	Marker *string `json:"marker,omitempty"`

	// 按名称过滤
	Name *string `json:"name,omitempty"`

	// 偏移值
	Offset *int32 `json:"offset,omitempty"`

	// 对象类型
	ProtectableType ListProtectableRequestProtectableType `json:"protectable_type"`

	// 资源的状态，如available，error 等
	Status *string `json:"status,omitempty"`

	// 根据资源id过滤
	Id *string `json:"id,omitempty"`

	// 根据该id过滤属于该服务器的所有磁盘，支持企业多项目的用户才能传入此参数
	ServerId *string `json:"server_id,omitempty"`

	// **参数解释：** 企业项目ID，获取企业项目ID请参考[查询企业项目列表](https://support.huaweicloud.com/api-em/zh-cn_topic_0121230880.html)(tag:hws)[查询企业项目列表](https://support.huaweicloud.com/intl/zh-cn/api-em/zh-cn_topic_0121230880.html)(tag:hws_hk)[《企业管理API参考》“查询企业项目列表”章节](tag:ctc,dt,tm,sbc,hws_eu,fcs_vm,hcso_dt,srg) **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListProtectableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectableRequest struct{}"
	}

	return strings.Join([]string{"ListProtectableRequest", string(data)}, " ")
}

type ListProtectableRequestProtectableType struct {
	value string
}

type ListProtectableRequestProtectableTypeEnum struct {
	SERVER       ListProtectableRequestProtectableType
	DISK         ListProtectableRequestProtectableType
	TURBO        ListProtectableRequestProtectableType
	WORKSPACE    ListProtectableRequestProtectableType
	WORKSPACE_V2 ListProtectableRequestProtectableType
}

func GetListProtectableRequestProtectableTypeEnum() ListProtectableRequestProtectableTypeEnum {
	return ListProtectableRequestProtectableTypeEnum{
		SERVER: ListProtectableRequestProtectableType{
			value: "server",
		},
		DISK: ListProtectableRequestProtectableType{
			value: "disk",
		},
		TURBO: ListProtectableRequestProtectableType{
			value: "turbo",
		},
		WORKSPACE: ListProtectableRequestProtectableType{
			value: "workspace",
		},
		WORKSPACE_V2: ListProtectableRequestProtectableType{
			value: "workspace_v2",
		},
	}
}

func (c ListProtectableRequestProtectableType) Value() string {
	return c.value
}

func (c ListProtectableRequestProtectableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProtectableRequestProtectableType) UnmarshalJSON(b []byte) error {
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
