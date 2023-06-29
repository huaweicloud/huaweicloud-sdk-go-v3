package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GesMetaData GES 元数据
type GesMetaData struct {

	// Label数据结构集合。
	Labels []Label `json:"labels"`
}

func (o GesMetaData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GesMetaData struct{}"
	}

	return strings.Join([]string{"GesMetaData", string(data)}, " ")
}
