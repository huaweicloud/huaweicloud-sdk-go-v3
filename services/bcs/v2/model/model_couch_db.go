package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// couchDB信息
type CouchDb struct {

	// couchDB用户名
	UserName string `json:"user_name" xml:"user_name"`

	// couchDB密码
	Password string `json:"password" xml:"password"`
}

func (o CouchDb) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CouchDb struct{}"
	}

	return strings.Join([]string{"CouchDb", string(data)}, " ")
}
