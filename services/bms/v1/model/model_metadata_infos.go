package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetadataInfos metadata字段数据结构说明
type MetadataInfos struct {

	// metadata键、值。键、值长度均不大于255字节。
	Key *string `json:"key,omitempty"`
}

func (o MetadataInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetadataInfos struct{}"
	}

	return strings.Join([]string{"MetadataInfos", string(data)}, " ")
}
