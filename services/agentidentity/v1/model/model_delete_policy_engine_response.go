package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePolicyEngineResponse Response Object
type DeletePolicyEngineResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePolicyEngineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyEngineResponse struct{}"
	}

	return strings.Join([]string{"DeletePolicyEngineResponse", string(data)}, " ")
}
