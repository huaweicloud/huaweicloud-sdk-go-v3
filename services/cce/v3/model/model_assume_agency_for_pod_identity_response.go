package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssumeAgencyForPodIdentityResponse Response Object
type AssumeAgencyForPodIdentityResponse struct {
	AssumedAgency *AssumedAgency `json:"assumedAgency,omitempty"`

	// **参数解释：** 凭据签发时传入的audience属性，通过pod-identity关联获取委托凭据的场景下，该值固定为 service.cce.pods。 该属性只在pod-identity关联绑定信任委托时返回 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Audience *string `json:"audience,omitempty"`

	Credentials *Credentials `json:"credentials,omitempty"`

	// **参数解释：** 委托凭据所属的pod-identity关联id。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PodIdentityAssociationId *string `json:"podIdentityAssociationId,omitempty"`

	Subject        *PodIdentitySubject `json:"subject,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o AssumeAgencyForPodIdentityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssumeAgencyForPodIdentityResponse struct{}"
	}

	return strings.Join([]string{"AssumeAgencyForPodIdentityResponse", string(data)}, " ")
}
