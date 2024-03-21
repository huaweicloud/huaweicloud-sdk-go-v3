package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDependenciesRequest Request Object
type ListDependenciesRequest struct {

	// 上一次查询依赖包的最后记录位置，默认为\"0\"。
	Marker *string `json:"marker,omitempty"`

	// 单次查询最大条数
	Maxitems *string `json:"maxitems,omitempty"`

	// 是否为公共依赖包
	Ispublic *string `json:"ispublic,omitempty"`

	// 依赖包类型public：公开,private:私有，all：全部。缺省时查询全量
	DependencyType *string `json:"dependency_type,omitempty"`

	// FunctionGraph函数的执行环境 Python2.7: Python语言2.7版本。 Python3.6: Pyton语言3.6版本。 Python3.9: Python语言3.9版本。 Go1.8: Go语言1.8版本。 Go1.x: Go语言1.x版本。 Java8: Java语言8版本。 Java11: Java语言11版本。 Node.js6.10: Nodejs语言6.10版本。 Node.js8.10: Nodejs语言8.10版本。 Node.js10.16: Nodejs语言10.16版本。 Node.js12.13: Nodejs语言12.13版本。 Node.js14.18: Nodejs语言14.18版本。 C#(.NET Core 2.0): C#语言2.0版本。 C#(.NET Core 2.1): C#语言2.1版本。 C#(.NET Core 3.1): C#语言3.1版本。 Custom: 自定义运行时。 PHP7.3: Php语言7.3版本。 http: HTTP函数。
	Runtime *string `json:"runtime,omitempty"`

	// 依赖包名称。
	Name *string `json:"name,omitempty"`

	// 本次查询可获取的依赖包的最大数目，默认为\"400\"。
	Limit *string `json:"limit,omitempty"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o ListDependenciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDependenciesRequest struct{}"
	}

	return strings.Join([]string{"ListDependenciesRequest", string(data)}, " ")
}
