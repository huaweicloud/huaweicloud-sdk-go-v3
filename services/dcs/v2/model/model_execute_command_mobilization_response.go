package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteCommandMobilizationResponse Response Object
type ExecuteCommandMobilizationResponse struct {

	// 返回信息
	Response       *string `json:"response,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteCommandMobilizationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteCommandMobilizationResponse struct{}"
	}

	return strings.Join([]string{"ExecuteCommandMobilizationResponse", string(data)}, " ")
}
