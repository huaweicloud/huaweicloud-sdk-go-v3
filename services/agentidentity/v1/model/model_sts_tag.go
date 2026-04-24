package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StsTag A custom key-value pair for STS session tags
type StsTag struct {

	// The key identifier, or name, of the tag
	Key string `json:"key"`

	// The string value that's associated with the key of the tag
	Value string `json:"value"`
}

func (o StsTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StsTag struct{}"
	}

	return strings.Join([]string{"StsTag", string(data)}, " ")
}
