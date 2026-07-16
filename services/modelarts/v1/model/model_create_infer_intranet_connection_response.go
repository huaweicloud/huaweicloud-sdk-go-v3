package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateInferIntranetConnectionResponse Response Object
type CreateInferIntranetConnectionResponse struct {

	// **参数解释：** 申请方用户名。 **取值范围：** 不涉及。
	ApplicantUserName *string `json:"applicant_user_name,omitempty"`

	// **参数解释：** 内网接入id。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 申请描述。 **取值范围：** 不涉及。
	Message *string `json:"message,omitempty"`

	// **参数解释：** 审核方domain name。  **取值范围：** 不涉及。
	OwnerDomainName *string `json:"owner_domain_name,omitempty"`

	// **参数解释：** 内网访问场景。 **约束限制：** 不涉及。 **取值范围：** - POOL：用户资源池接入场景 - VPC：用户VPC接入场景 **默认取值：** 不涉及。
	Scene *string `json:"scene,omitempty"`

	// **参数解释：** 服务ID。 **取值范围：** 不涉及。
	ServiceId *string `json:"service_id,omitempty"`

	// **参数解释：** 服务名。 **取值范围：** 不涉及。
	ServiceName *string `json:"service_name,omitempty"`

	// **参数解释：** 内网接入状态，支持列表查询。 **约束限制：** 不涉及。 **取值范围：** - APPROVING：审批中 - REJECTED：拒绝 - CONNECTING：接入中 - CONNECTED：已接入 - CANCELED：已取消 - FAILED：失败 - DELETING：删除中 **默认取值：** 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释：** 子网ID。 **取值范围：** 不涉及。
	SubnetId *string `json:"subnet_id,omitempty"`

	// **参数解释：** 访问地址列表。
	UrlList *[]string `json:"url_list,omitempty"`

	// **参数解释：** 访问地址列表。
	CustomUrlList *[]string `json:"custom_url_list,omitempty"`

	// **参数解释：** VPC ID。 **取值范围：** 不涉及。
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释：** 服务绑定的dispatcher组ID。 **取值范围：** 不涉及。
	DispatcherGroupId *string `json:"dispatcher_group_id,omitempty"`

	// **参数解释：** 接入粒度：SERVICE、GLOBAL **取值范围：** 不涉及。
	Type *CreateInferIntranetConnectionResponseType `json:"type,omitempty"`

	// **参数解释：** 资源池网络名称。 **取值范围：** 不涉及。
	MaosNetworkName *string `json:"maos_network_name,omitempty"`

	// **参数解释：** 服务类型。 **取值范围：** 不涉及。
	ServiceType *string `json:"service_type,omitempty"`

	// **参数解释：** 资源池ID。 **取值范围：** 不涉及。
	PoolId *string `json:"pool_id,omitempty"`

	// **参数解释：** 创建时间。 **取值范围：** 不涉及。
	CreateAt *string `json:"create_at,omitempty"`

	// **参数解释：** 修改时间。 **取值范围：** 不涉及。
	UpdateAt       *string `json:"update_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInferIntranetConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInferIntranetConnectionResponse struct{}"
	}

	return strings.Join([]string{"CreateInferIntranetConnectionResponse", string(data)}, " ")
}

type CreateInferIntranetConnectionResponseType struct {
	value string
}

type CreateInferIntranetConnectionResponseTypeEnum struct {
	SERVICE CreateInferIntranetConnectionResponseType
	GLOBAL  CreateInferIntranetConnectionResponseType
}

func GetCreateInferIntranetConnectionResponseTypeEnum() CreateInferIntranetConnectionResponseTypeEnum {
	return CreateInferIntranetConnectionResponseTypeEnum{
		SERVICE: CreateInferIntranetConnectionResponseType{
			value: "SERVICE",
		},
		GLOBAL: CreateInferIntranetConnectionResponseType{
			value: "GLOBAL",
		},
	}
}

func (c CreateInferIntranetConnectionResponseType) Value() string {
	return c.value
}

func (c CreateInferIntranetConnectionResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInferIntranetConnectionResponseType) UnmarshalJSON(b []byte) error {
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
