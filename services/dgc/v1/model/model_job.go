package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Job struct {
	Name *string `json:"name,omitempty" xml:"name"`

	Path *string `json:"path,omitempty" xml:"path"`

	Params *interface{} `json:"params,omitempty" xml:"params"`
}

func (o Job) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Job struct{}"
	}

	return strings.Join([]string{"Job", string(data)}, " ")
}
