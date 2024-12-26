package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Mapstringstring struct {
}

func (o Mapstringstring) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Mapstringstring struct{}"
	}

	return strings.Join([]string{"Mapstringstring", string(data)}, " ")
}
