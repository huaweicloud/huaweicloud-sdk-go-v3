package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddEcnWithErRequest Request Object
type AddEcnWithErRequest struct {

	// 企业连接网络ID
	EcnId string `json:"ecn_id"`

	Body *EcnWithErRequest `json:"body,omitempty"`
}

func (o AddEcnWithErRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddEcnWithErRequest struct{}"
	}

	return strings.Join([]string{"AddEcnWithErRequest", string(data)}, " ")
}
