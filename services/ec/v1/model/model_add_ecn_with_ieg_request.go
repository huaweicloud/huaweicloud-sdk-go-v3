package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddEcnWithIegRequest Request Object
type AddEcnWithIegRequest struct {

	// 企业连接网络ID
	EcnId string `json:"ecn_id"`

	Body *EcnWithIegRequest `json:"body,omitempty"`
}

func (o AddEcnWithIegRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddEcnWithIegRequest struct{}"
	}

	return strings.Join([]string{"AddEcnWithIegRequest", string(data)}, " ")
}
