package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StarRocksCreateRequestEngine 引擎信息。
type StarRocksCreateRequestEngine struct {

	// 引擎类型。仅支持star-rocks。
	Type string `json:"type"`

	// 引擎大版本号。
	Version string `json:"version"`
}

func (o StarRocksCreateRequestEngine) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StarRocksCreateRequestEngine struct{}"
	}

	return strings.Join([]string{"StarRocksCreateRequestEngine", string(data)}, " ")
}
