package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociationDto struct {
	Accessor *AccessorDto `json:"accessor"`

	InstanceId string `json:"instance_id"`

	ProfileId string `json:"profile_id"`
}

func (o AssociationDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociationDto struct{}"
	}

	return strings.Join([]string{"AssociationDto", string(data)}, " ")
}
