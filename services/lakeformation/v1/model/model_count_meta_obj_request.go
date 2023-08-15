package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountMetaObjRequest Request Object
type CountMetaObjRequest struct {

	// instance id
	InstanceId string `json:"instance_id"`
}

func (o CountMetaObjRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountMetaObjRequest struct{}"
	}

	return strings.Join([]string{"CountMetaObjRequest", string(data)}, " ")
}
