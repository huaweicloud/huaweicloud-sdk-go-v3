/*
 * BMS
 *
 * BMS Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// image字段数据结构说明
type Image struct {
	// 裸金属服务器镜像ID
	Id *string `json:"id,omitempty"`
	// 裸金属服务器镜像相关快捷链接信息
	Links *[]Links `json:"links,omitempty"`
}

func (o Image) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"Image", string(data)}, " ")
}
