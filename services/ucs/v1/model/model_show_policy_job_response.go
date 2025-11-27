package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPolicyJobResponse Response Object
type ShowPolicyJobResponse struct {

	// API类型
	Kind *string `json:"kind,omitempty"`

	// API版本
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *ObjectMeta `json:"metadata,omitempty"`

	Spec *UcsJobSpec `json:"spec,omitempty"`

	Status         *UcsJobStatus `json:"status,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowPolicyJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyJobResponse struct{}"
	}

	return strings.Join([]string{"ShowPolicyJobResponse", string(data)}, " ")
}
