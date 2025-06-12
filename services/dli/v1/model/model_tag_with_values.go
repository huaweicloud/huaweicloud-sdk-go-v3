package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagWithValues Resource Label
type TagWithValues struct {

	// Key
	Key string `json:"key"`

	// Values
	Values []string `json:"values"`
}

func (o TagWithValues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagWithValues struct{}"
	}

	return strings.Join([]string{"TagWithValues", string(data)}, " ")
}
