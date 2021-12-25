package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddTag struct {
	Tag *ResourceTag `json:"tag"`
}

func (o AddTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTag struct{}"
	}

	return strings.Join([]string{"AddTag", string(data)}, " ")
}
