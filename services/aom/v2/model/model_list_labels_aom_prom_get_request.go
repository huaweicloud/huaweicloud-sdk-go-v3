package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListLabelsAomPromGetRequest struct {
}

func (o ListLabelsAomPromGetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLabelsAomPromGetRequest struct{}"
	}

	return strings.Join([]string{"ListLabelsAomPromGetRequest", string(data)}, " ")
}
