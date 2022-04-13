package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListMetadataAomPromGetRequest struct {
}

func (o ListMetadataAomPromGetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetadataAomPromGetRequest struct{}"
	}

	return strings.Join([]string{"ListMetadataAomPromGetRequest", string(data)}, " ")
}
