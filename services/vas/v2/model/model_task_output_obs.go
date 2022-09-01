package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 输出为obs类型时的配置信息
type TaskOutputObs struct {

	// OBS桶名，选用obs类型输出时必填。
	Bucket string `json:"bucket" xml:"bucket"`

	// OBS的路径，选用obs类型输出时必填。
	Path string `json:"path" xml:"path"`

	// 作业输出数据类别的列表，选填，仅部分服务需要。当输出类型下有这个列表时，表示希望这个输出类型下存放dataCategory列表内的数据。
	DataCategory *[]string `json:"data_category,omitempty" xml:"data_category"`
}

func (o TaskOutputObs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskOutputObs struct{}"
	}

	return strings.Join([]string{"TaskOutputObs", string(data)}, " ")
}
