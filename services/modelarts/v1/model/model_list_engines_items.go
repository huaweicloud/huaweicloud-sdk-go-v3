package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListEnginesItems struct {

	// 引擎规格的ID。如“caffe-1.0.0-python2.7”。
	EngineId *string `json:"engine_id,omitempty"`

	// 引擎规格的名称。如“Caffe”。
	EngineName *string `json:"engine_name,omitempty"`

	// 引擎规格的版本。对一个引擎名称，有多个版本的引擎，如使用python2.7的\"Caffe-1.0.0-python2.7\"等。
	EngineVersion *string `json:"engine_version,omitempty"`

	// 是否为v1兼容模式。
	V1Compatible *bool `json:"v1_compatible,omitempty"`

	// 引擎默认启动用户uid。
	RunUser *string `json:"run_user,omitempty"`

	ImageInfo *ListEnginesImageInfo `json:"image_info,omitempty"`
}

func (o ListEnginesItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnginesItems struct{}"
	}

	return strings.Join([]string{"ListEnginesItems", string(data)}, " ")
}
