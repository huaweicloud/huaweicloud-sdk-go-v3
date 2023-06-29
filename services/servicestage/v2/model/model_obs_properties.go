package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObsProperties 软件包的其他属性，只有在选择对象存储obs的时候才需要添加。
type ObsProperties struct {

	// obs的终端地址，比如：https://obs.region_name.external_domain_name.com。
	Endpoint *string `json:"endpoint,omitempty"`

	// 软件包在obs的桶名。
	Bucket *string `json:"bucket,omitempty"`

	// obs桶中的对象，一般是软件包名，有文件夹的话要加上文件夹的路径。比如test.jar或者demo/test.jar。
	Key *string `json:"key,omitempty"`
}

func (o ObsProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsProperties struct{}"
	}

	return strings.Join([]string{"ObsProperties", string(data)}, " ")
}
