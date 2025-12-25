package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateExtendAttributeResponse Response Object
type UpdateExtendAttributeResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	ExtendAttribute *VifExtendAttribute `json:"extend_attribute,omitempty"`
	HttpStatusCode  int                 `json:"-"`
}

func (o UpdateExtendAttributeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateExtendAttributeResponse struct{}"
	}

	return strings.Join([]string{"UpdateExtendAttributeResponse", string(data)}, " ")
}
