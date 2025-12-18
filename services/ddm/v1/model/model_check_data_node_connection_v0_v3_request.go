package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckDataNodeConnectionV0V3Request Request Object
type CheckDataNodeConnectionV0V3Request struct {
	Body *EsdbCheckRdsConnectionsRequestV3 `json:"body,omitempty"`
}

func (o CheckDataNodeConnectionV0V3Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDataNodeConnectionV0V3Request struct{}"
	}

	return strings.Join([]string{"CheckDataNodeConnectionV0V3Request", string(data)}, " ")
}
