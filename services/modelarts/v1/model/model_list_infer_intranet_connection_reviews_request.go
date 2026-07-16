package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInferIntranetConnectionReviewsRequest Request Object
type ListInferIntranetConnectionReviewsRequest struct {

	// **参数解释：** 内网访问场景。 **约束限制：** 不涉及。 **取值范围：** - POOL：用户资源池接入场景 - VPC：用户VPC网络接入场景 **默认取值：** 不涉及。
	Scene *string `json:"scene,omitempty"`

	// **参数解释：** 内网接入id。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 申请方domain ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ApplicantDomainId *string `json:"applicant_domain_id,omitempty"`

	// **参数解释：** 服务ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ServiceId *string `json:"service_id,omitempty"`

	// **参数解释：** 申请方用户名。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ApplicantUserName *string `json:"applicant_user_name,omitempty"`

	// **参数解释：** 服务名称。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ServiceName *string `json:"service_name,omitempty"`

	// **参数解释：** VPC名称。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	VpcName *string `json:"vpc_name,omitempty"`

	// **参数解释：** VPC ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释：** 资源池ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	PoolId *string `json:"pool_id,omitempty"`

	// **参数解释：** 排序字段，支持update_at、create_at，默认值update_at。 **约束限制：** 不涉及。 **取值范围：** - update_at：按更新时间排序。 - create_at：按创建时间排序。 **默认取值：** update_at。
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释：** 指定返回的最大条目数。 **约束限制：** 不涉及。 **取值范围：** [1,500] **默认取值：** 10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 分页列表查询的偏移量。 **约束限制：** offset必须是limit的整数倍。 **取值范围：** 不涉及。 **默认取值：** 0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 申请状态。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释：** 内网申请类型。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Type *string `json:"type,omitempty"`
}

func (o ListInferIntranetConnectionReviewsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInferIntranetConnectionReviewsRequest struct{}"
	}

	return strings.Join([]string{"ListInferIntranetConnectionReviewsRequest", string(data)}, " ")
}
