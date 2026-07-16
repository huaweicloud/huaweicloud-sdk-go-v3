package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSearchAlgorithmsRequest Request Object
type ShowSearchAlgorithmsRequest struct {
}

func (o ShowSearchAlgorithmsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSearchAlgorithmsRequest struct{}"
	}

	return strings.Join([]string{"ShowSearchAlgorithmsRequest", string(data)}, " ")
}
