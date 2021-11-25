package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// couchDB信息
type Couchdb struct {
	// couchDB用户名

	UserName string `json:"user_name"`
	// couchDB密码

	Password string `json:"password"`
}

func (o Couchdb) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Couchdb struct{}"
	}

	return strings.Join([]string{"Couchdb", string(data)}, " ")
}
