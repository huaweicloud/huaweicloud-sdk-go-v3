package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPolicyInstanceResponse Response Object
type ShowPolicyInstanceResponse struct {

	// API类型
	Kind *string `json:"kind,omitempty"`

	// API版本
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *ObjectMeta `json:"metadata,omitempty"`

	Spec *UcsConstraintSpec `json:"spec,omitempty"`

	Status         *UcsConstraintStatus `json:"status,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowPolicyInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowPolicyInstanceResponse", string(data)}, " ")
}
