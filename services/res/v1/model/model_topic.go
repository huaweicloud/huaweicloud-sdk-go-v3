package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Topic struct {

	// 名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// id。
	Id *string `json:"id,omitempty" xml:"id"`
}

func (o Topic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Topic struct{}"
	}

	return strings.Join([]string{"Topic", string(data)}, " ")
}
