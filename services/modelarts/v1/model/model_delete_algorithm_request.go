package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlgorithmRequest Request Object
type DeleteAlgorithmRequest struct {

	// 算法ID。
	AlgorithmId string `json:"algorithm_id"`
}

func (o DeleteAlgorithmRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlgorithmRequest struct{}"
	}

	return strings.Join([]string{"DeleteAlgorithmRequest", string(data)}, " ")
}
