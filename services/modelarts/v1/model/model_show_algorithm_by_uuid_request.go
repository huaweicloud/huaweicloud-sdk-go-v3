package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlgorithmByUuidRequest Request Object
type ShowAlgorithmByUuidRequest struct {

	// 算法ID。
	AlgorithmId string `json:"algorithm_id"`
}

func (o ShowAlgorithmByUuidRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlgorithmByUuidRequest struct{}"
	}

	return strings.Join([]string{"ShowAlgorithmByUuidRequest", string(data)}, " ")
}
