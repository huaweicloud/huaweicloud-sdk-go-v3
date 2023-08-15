package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectRequest Request Object
type ListProjectRequest struct {
}

func (o ListProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectRequest struct{}"
	}

	return strings.Join([]string{"ListProjectRequest", string(data)}, " ")
}
