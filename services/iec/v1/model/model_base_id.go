package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ID对象
type BaseId struct {

	// 对象ID，uuid。
	Id string `json:"id" xml:"id"`
}

func (o BaseId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseId struct{}"
	}

	return strings.Join([]string{"BaseId", string(data)}, " ")
}
