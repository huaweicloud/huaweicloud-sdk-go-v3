package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据搜索引擎类型。
type CreateClusterDatastoreBody struct {

	// 引擎版本号。  elasticsearch支持5.5.1、6.2.3、6.5.4、7.1.1、7.6.2或7.9.3。  logstash支持5.6.16或7.10.0。
	Version string `json:"version"`

	// 引擎类型，支持elasticsearch和logstash。
	Type string `json:"type"`
}

func (o CreateClusterDatastoreBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterDatastoreBody struct{}"
	}

	return strings.Join([]string{"CreateClusterDatastoreBody", string(data)}, " ")
}
