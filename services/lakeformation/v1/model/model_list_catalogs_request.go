package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCatalogsRequest Request Object
type ListCatalogsRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`
}

func (o ListCatalogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCatalogsRequest struct{}"
	}

	return strings.Join([]string{"ListCatalogsRequest", string(data)}, " ")
}
