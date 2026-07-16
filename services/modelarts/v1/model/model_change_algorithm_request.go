package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAlgorithmRequest Request Object
type ChangeAlgorithmRequest struct {

	// 算法ID。
	AlgorithmId string `json:"algorithm_id"`

	Body *Algorithm `json:"body,omitempty"`
}

func (o ChangeAlgorithmRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAlgorithmRequest struct{}"
	}

	return strings.Join([]string{"ChangeAlgorithmRequest", string(data)}, " ")
}
