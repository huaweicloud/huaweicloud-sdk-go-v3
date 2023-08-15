package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MultiTaskInitBodyExtInfo 扩展属性，目前支持server_id，允许为空
type MultiTaskInitBodyExtInfo struct {

	// Mysql的ServerID
	ServerId *string `json:"server_id,omitempty"`
}

func (o MultiTaskInitBodyExtInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiTaskInitBodyExtInfo struct{}"
	}

	return strings.Join([]string{"MultiTaskInitBodyExtInfo", string(data)}, " ")
}
