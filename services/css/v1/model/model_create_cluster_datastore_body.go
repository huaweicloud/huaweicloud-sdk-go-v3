package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据搜索引擎类型。
type CreateClusterDatastoreBody struct {
	// 引擎版本号，支持5.5.1、6.2.3、6.5.4、7.1.1、7.6.2和7.9.3，默认为5.5.1。

	Version string `json:"version"`
	// 引擎类型，默认为elasticsearch。目前只支持elasticsearch。

	Type string `json:"type"`
}

func (o CreateClusterDatastoreBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterDatastoreBody struct{}"
	}

	return strings.Join([]string{"CreateClusterDatastoreBody", string(data)}, " ")
}
