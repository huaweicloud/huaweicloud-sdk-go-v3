package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFlowLogResponse Response Object
type DeleteFlowLogResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteFlowLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFlowLogResponse struct{}"
	}

	return strings.Join([]string{"DeleteFlowLogResponse", string(data)}, " ")
}
