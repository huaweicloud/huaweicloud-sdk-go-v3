package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceOption 创建边缘实例请求参数
type CreateInstanceOption struct {

	// 边缘资源名称，与边缘实例一一对应。 取值范围：只能由中文字符、大小写英文字母、数字及中划线、下划线组成，且长度为[1-48]个字符。
	Name string `json:"name"`

	// 是否自动添加前缀。 - with_prefix为false时不拼接IEC前缀 - with_prefix不传或者传true时拼自动IEC前缀  以name为iec为例： 不添加前缀时实例名称为：iec 自动添加前缀实例名称为：IEC-ZS01-iec
	WithPrefix *bool `json:"with_prefix,omitempty"`

	// 边缘实例的系统镜像，需要指定已创建镜像的ID，ID格式为通用唯一识别码（Universally Unique Identifier，简称UUID）。 镜像的ID可以从控制台或者参考本文档的“查询边镜像列表”的章节获取。
	ImageRef string `json:"image_ref"`

	// 边缘实例的系统规格的ID。 已上线的规格请使用接口 “查询边缘实例规格列表“ 进行查询。
	FlavorRef string `json:"flavor_ref"`

	// 如果需要使用密码方式登录边缘实例，可使用adminPass字段指定边缘实例管理员帐户初始登录密码。其中，Linux管理员帐户为root，Windows管理员帐户为Administrator。 密码复杂度要求： 1. 长度为8-26位。 2. 密码至少必须包含大写字母、小写字母、数字和特殊字符（!@$%^-_=+[{}]:,./?）中的三种。 3. 密码不能包含用户名或用户名的逆序。 4. Windows系统密码不能包含用户名或用户名的逆序，不能包含用户名中超过两个连续字符的部分。 说明 目前边缘实例不支持创建后设置密码，不设置此参数会导致实例创建后无法登录。
	AdminPass *string `json:"admin_pass,omitempty"`

	// 密钥对名称。
	KeyName *string `json:"key_name,omitempty"`

	NetConfig *NetConfigInstance `json:"net_config"`

	Bandwidth *BandwidthConfigInstance `json:"bandwidth,omitempty"`

	RootVolume *RootVolume `json:"root_volume"`

	// 边缘实例对应数据盘相关配置。每一个数据结构代表一块待创建的数据盘。 约束：目前边缘实例最多可挂载2块数据盘
	DataVolumes *[]DataVolume `json:"data_volumes,omitempty"`

	// 边缘实例数量。
	Count int32 `json:"count"`

	// 边缘业务对应安全组信息。
	SecurityGroups *[]SecurityGroupOption `json:"security_groups,omitempty"`

	Coverage *CoverageInstance `json:"coverage"`

	// 创建边缘实例过程中注入用户数据。支持注入文本、文本文件或gzip文件。 更多关于待注入用户数据的信息，请参见《弹性云服务器用户指南 》的“[用户数据注入](https://support.huaweicloud.com/usermanual-ecs/zh-cn_topic_0032380449.html)”章节。
	UserData *string `json:"user_data,omitempty"`
}

func (o CreateInstanceOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceOption struct{}"
	}

	return strings.Join([]string{"CreateInstanceOption", string(data)}, " ")
}
