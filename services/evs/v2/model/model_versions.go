package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 版本信息结构体
type Versions struct {

	// 接口版本的ID
	Id string `json:"id"`

	// 接口版本信息的URI描述信息
	Links []Link `json:"links"`

	// 接口版本的请求消息类型信息
	MediaTypes []MediaTypes `json:"media-types"`

	// 接口版本的最小版本号
	MinVersion *string `json:"min_version,omitempty"`

	// 接口版本的状态
	Status string `json:"status"`

	// 接口版本更新时间
	Updated string `json:"updated"`

	// 接口版本的版本号信息
	Version string `json:"version"`
}

func (o Versions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Versions struct{}"
	}

	return strings.Join([]string{"Versions", string(data)}, " ")
}
