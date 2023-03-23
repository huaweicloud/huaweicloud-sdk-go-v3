package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdatePrivateSnatResponse struct {

	// 请求ID。
	RequestId *string `json:"request_id,omitempty"`

	SnatRule       *PrivateSnat `json:"snat_rule,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdatePrivateSnatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateSnatResponse struct{}"
	}

	return strings.Join([]string{"UpdatePrivateSnatResponse", string(data)}, " ")
}
