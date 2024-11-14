package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateObsTargetAttributesResponse Response Object
type UpdateObsTargetAttributesResponse struct {

	// 绑定关系Id
	TargetId *string `json:"target_id,omitempty"`

	Attributes *ObsTargetAttributes `json:"attributes,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateObsTargetAttributesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateObsTargetAttributesResponse struct{}"
	}

	return strings.Join([]string{"UpdateObsTargetAttributesResponse", string(data)}, " ")
}
