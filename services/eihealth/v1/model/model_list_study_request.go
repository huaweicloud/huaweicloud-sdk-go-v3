package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStudyRequest Request Object
type ListStudyRequest struct {
}

func (o ListStudyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStudyRequest struct{}"
	}

	return strings.Join([]string{"ListStudyRequest", string(data)}, " ")
}
