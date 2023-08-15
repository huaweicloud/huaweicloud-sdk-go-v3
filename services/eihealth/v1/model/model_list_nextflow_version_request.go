package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNextflowVersionRequest Request Object
type ListNextflowVersionRequest struct {
}

func (o ListNextflowVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNextflowVersionRequest struct{}"
	}

	return strings.Join([]string{"ListNextflowVersionRequest", string(data)}, " ")
}
