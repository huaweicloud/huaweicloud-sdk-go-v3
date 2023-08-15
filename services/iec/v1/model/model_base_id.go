package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BaseId ID对象
type BaseId struct {

	// 对象ID，uuid。
	Id string `json:"id"`
}

func (o BaseId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseId struct{}"
	}

	return strings.Join([]string{"BaseId", string(data)}, " ")
}
