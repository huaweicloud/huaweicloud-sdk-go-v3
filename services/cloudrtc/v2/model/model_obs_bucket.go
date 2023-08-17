package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObsBucket struct {

	// OBS的bucket名称
	Bucket string `json:"bucket"`

	// OBS的bucket所在region
	Location string `json:"location"`

	// 创建时间，格式如：2020/07/16 15:11:55 GMT+08:00
	CreationDate string `json:"creation_date"`

	// 鉴权
	IsAuthorized bool `json:"is_authorized"`
}

func (o ObsBucket) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsBucket struct{}"
	}

	return strings.Join([]string{"ObsBucket", string(data)}, " ")
}
