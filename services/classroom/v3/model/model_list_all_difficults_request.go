package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllDifficultsRequest Request Object
type ListAllDifficultsRequest struct {
}

func (o ListAllDifficultsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllDifficultsRequest struct{}"
	}

	return strings.Join([]string{"ListAllDifficultsRequest", string(data)}, " ")
}
