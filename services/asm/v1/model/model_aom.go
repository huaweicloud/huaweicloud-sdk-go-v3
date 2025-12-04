package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Aom struct {

	// AOM实例ID。
	InstanceID string `json:"instanceID"`
}

func (o Aom) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Aom struct{}"
	}

	return strings.Join([]string{"Aom", string(data)}, " ")
}
