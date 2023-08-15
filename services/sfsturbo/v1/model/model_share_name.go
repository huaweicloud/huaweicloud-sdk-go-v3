package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShareName 需要修改的SFS Turbo文件系统
type ShareName struct {

	// 需要修改的SFS Turbo文件系统的名字
	Name string `json:"name"`
}

func (o ShareName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShareName struct{}"
	}

	return strings.Join([]string{"ShareName", string(data)}, " ")
}
