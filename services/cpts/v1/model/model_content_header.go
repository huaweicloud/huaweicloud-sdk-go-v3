package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ContentHeader struct {

	// key
	Key *string `json:"key,omitempty" xml:"key"`

	// value
	Value *string `json:"value,omitempty" xml:"value"`
}

func (o ContentHeader) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentHeader struct{}"
	}

	return strings.Join([]string{"ContentHeader", string(data)}, " ")
}
