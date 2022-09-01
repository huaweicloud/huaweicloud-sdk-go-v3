package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type DataConfig struct {
	Offline *Offline `json:"offline" xml:"offline"`

	Nearline *Nearline `json:"nearline,omitempty" xml:"nearline"`
}

func (o DataConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataConfig struct{}"
	}

	return strings.Join([]string{"DataConfig", string(data)}, " ")
}
