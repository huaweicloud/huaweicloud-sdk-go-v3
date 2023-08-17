package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDataRestInfo 添加数据的相关信息。
type AddDataRestInfo struct {
	ImageInfo *AddDataRestInfoImageInfo `json:"image_info,omitempty"`
}

func (o AddDataRestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDataRestInfo struct{}"
	}

	return strings.Join([]string{"AddDataRestInfo", string(data)}, " ")
}
