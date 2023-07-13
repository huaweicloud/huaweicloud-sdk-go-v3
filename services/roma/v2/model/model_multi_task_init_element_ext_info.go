package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MultiTaskInitElementExtInfo 扩展属性，目前只支持server_id
type MultiTaskInitElementExtInfo struct {

	// Mysql的ServerID
	ServerId *string `json:"server_id,omitempty"`
}

func (o MultiTaskInitElementExtInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiTaskInitElementExtInfo struct{}"
	}

	return strings.Join([]string{"MultiTaskInitElementExtInfo", string(data)}, " ")
}
