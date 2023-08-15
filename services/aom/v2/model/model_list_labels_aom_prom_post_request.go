package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLabelsAomPromPostRequest Request Object
type ListLabelsAomPromPostRequest struct {
}

func (o ListLabelsAomPromPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLabelsAomPromPostRequest struct{}"
	}

	return strings.Join([]string{"ListLabelsAomPromPostRequest", string(data)}, " ")
}
