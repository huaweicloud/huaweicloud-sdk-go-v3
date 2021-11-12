package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListKeypairsRequest struct {
}

func (o ListKeypairsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeypairsRequest struct{}"
	}

	return strings.Join([]string{"ListKeypairsRequest", string(data)}, " ")
}
