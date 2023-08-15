package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTemplateDetailResponse Response Object
type ShowTemplateDetailResponse struct {

	// 模板的id。
	Id *string `json:"id,omitempty"`

	// 模板的名称。
	Title *string `json:"title,omitempty"`

	// 模板的描述信息。
	Description *string `json:"description,omitempty"`

	// 模板关联的region host id。
	RegionId *string `json:"region_id,omitempty"`

	// 模板关联的repo id。
	RepostoryId *string `json:"repostory_id,omitempty"`

	// 模板https下载路径。
	CodeUrl *string `json:"code_url,omitempty"`

	// 模板ssh下载路径。
	SshUrl *string `json:"ssh_url,omitempty"`

	// 项目id。
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 模板状态。
	Status *int32 `json:"status,omitempty"`

	// 源数据信息： - key：元数据标识 - defaultValue：用户输入值的默认值 - isShow：前台界面组件是否展示该元数据 - isProjectName：是否使用作为项目名称 - label：前台界面组件展示名称 - type：前台界面组件类型 - helpText：前台界面组件帮助文本 - readOnly：前台界面组件是否可修改 - required：前台界面组件是否展示必填 - regType：该元数据进行正则校验类型；简化模板编码使用 - regPattern：该元数据对应js语法正则表达式 - regTip：该元数据正则校验提示信息 - visibleRule：该元数据可见规则 - isRequired：是否必填 - isReadOnly：是否只读 - options：option对象集合   - displayName：前台界面展示字符串   - value：该选项值 - eventOnchange：联动属性集合   - associatedProperty：被关联Property的key值   - associatedValue：被关联的value - fold：是否折叠 - show：是否展示该Property
	Properties *[]interface{} `json:"properties,omitempty"`

	// dependency信息： - id：依赖全局唯一标识 - name：依赖展示名称 - description：依赖展示描述 - recommended：是否推荐使用该依赖 - versionProperty：该依赖版本被关联Property的key值 - versionRange：该依赖版本适用范围 - groupName：分组名称 - items：分组列表
	Dependencies *[]interface{} `json:"dependencies,omitempty"`

	// dependency类型： - 0：分组 - 1：不分组 - null：无分组信息
	DependencyType *string `json:"dependency_type,omitempty"`

	// 部署信息： - param：参数对象   - build：构建类型   - runtime：函数运行时   - handler：函数执行入口   - outputFile：构建产物文件路径 - target：部署环境
	Deployment     *interface{} `json:"deployment,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowTemplateDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowTemplateDetailResponse", string(data)}, " ")
}
