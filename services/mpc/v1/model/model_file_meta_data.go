package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FileMetaData struct {

	// metadata名称
	Name *string `json:"name,omitempty"`

	// metadata值
	Value *string `json:"value,omitempty"`
}

func (o FileMetaData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileMetaData struct{}"
	}

	return strings.Join([]string{"FileMetaData", string(data)}, " ")
}
