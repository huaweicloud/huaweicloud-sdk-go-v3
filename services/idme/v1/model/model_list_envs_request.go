package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnvsRequest Request Object
type ListEnvsRequest struct {
}

func (o ListEnvsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvsRequest struct{}"
	}

	return strings.Join([]string{"ListEnvsRequest", string(data)}, " ")
}
