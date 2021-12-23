package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 边缘资源对象
type Resource struct {
	// 边缘实例名称。 取值范围： 只能由中文字符、英文字母 (大小写)、数字及“_”、“-”、“.” 组成。 创建的边缘实例数量（count字段对应的值）大于1时，为区分不同边缘实例，创建过程中系统会自动在名称后加“-000x”的类似标记。  >  华为云边缘实例内部主机名 （hostname）命名规则遵循RFC 952和RFC 1123命名规范，建议使用a-zA-z或0-9以及中划线'-'组成的名称命名，' _' 将在边缘实例内部默认转化为'-'。

	Name string `json:"name"`
	// 是否自动添加名称前缀。 - with_prefix为false时不拼接IEC前缀 - with_prefix不传或者传true时拼自动IEC前缀  以name为iec为例： 不添加前缀时实例名称为：iec-0001 自动添加前缀实例名称为：IEC-ZS01-iec-0001 0001为创建边缘业务时根据实例个数自动添加的编号

	WithPrefix *bool `json:"with_prefix,omitempty"`
	// 待发放边缘实例的系统镜像，需要指定已创建镜像的ID。 > 镜像的ID可以从控制台或者参考本文档的“查询边镜像列表”的章节获取。

	ImageRef string `json:"image_ref"`
	// 边缘实例的系统规格的ID。

	FlavorRef string `json:"flavor_ref"`
	// 如果需要使用密码方式登录边缘实例，可使用admin_pass字段指定边缘实例管理员帐户初始登录密码。其中，Linux管理员帐户为root，Windows管理员帐户为Administrator。  密码复杂度要求： - 长度为8-26位。 - 密码至少必须包含大写字母、小写字母、数字和特殊字符（!@$%^-_=+[{}]:,./?）中的三种。 - 密码不能包含用户名或用户名的逆序。 - Windows系统密码不能包含用户名或用户名的逆序，不能包含用户名中超过两个连续字符的部分。  >   目前边缘实例不支持创建后设置密码，不设置此参数会导致实例无法登录。

	AdminPass *string `json:"admin_pass,omitempty"`
	// 密钥对名称。

	KeyName *string `json:"key_name,omitempty"`

	NetConfig *NetConfig `json:"net_config"`

	Bandwidth *BandwidthConfig `json:"bandwidth,omitempty"`

	RootVolume *RootVolume `json:"root_volume"`
	// 边缘实例对应数据盘相关配置。每一个数据结构代表一块待创建的数据盘。  约束：目前边缘实例最多可挂载2块数据盘

	DataVolumes *[]DataVolume `json:"data_volumes,omitempty"`
	// 边缘实例数量。 不传该字段时默认取值为1。

	Count *int32 `json:"count,omitempty"`
	// 边缘业务对应安全组信息。

	SecurityGroups *[]SecurityGroupOption `json:"security_groups,omitempty"`
	// 创建边缘实例过程中注入用户数据。支持注入文本、文本文件或gzip文件。 更多关于待注入用户数据的信息，请参见《弹性云服务器用户指南 》的“[用户数据注入](https://support.huaweicloud.com/usermanual-ecs/zh-cn_topic_0032380449.html)”章节。

	UserData *string `json:"user_data,omitempty"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
