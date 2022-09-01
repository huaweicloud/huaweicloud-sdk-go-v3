package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 应用详细信息
type ApplicationV3 struct {

	// 应用id
	Id *string `json:"id,omitempty" xml:"id"`

	// 应用名称
	Name string `json:"name" xml:"name"`

	// 应用描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 区域id
	RegionId string `json:"region_id" xml:"region_id"`

	// 区域名称
	RegionName string `json:"region_name" xml:"region_name"`

	// 所属项目id
	ProjectId string `json:"project_id" xml:"project_id"`

	// 项目名称
	ProjectName string `json:"project_name" xml:"project_name"`

	// 应用图标
	Icon *string `json:"icon,omitempty" xml:"icon"`

	// 流水线列表
	PipelineCreationResult *[]PipelineCreationResult `json:"pipeline_creation_result,omitempty" xml:"pipeline_creation_result"`

	RepositoryCreationResult *RepositoryCreationResult `json:"repository_creation_result,omitempty" xml:"repository_creation_result"`

	// 环境信息
	EnvironmentCreationResult *[]string `json:"environment_creation_result,omitempty" xml:"environment_creation_result"`

	// 模板类型
	TemplateTypes *[]TemplateType `json:"template_types,omitempty" xml:"template_types"`

	// 模板部署信息
	TemplateDeployment *string `json:"template_deployment,omitempty" xml:"template_deployment"`

	// 部署类型, function:函数部署,cci:cci容器部署,ServiceStage(Jar):ServiceStage jar包部署,ServiceStage(Docker):ServiceStage Docker容器部署
	DeployType *ApplicationV3DeployType `json:"deploy_type,omitempty" xml:"deploy_type"`

	// 创建者名称
	CreatorName *string `json:"creator_name,omitempty" xml:"creator_name"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at"`

	// 应用版本号
	Version *string `json:"version,omitempty" xml:"version"`
}

func (o ApplicationV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationV3 struct{}"
	}

	return strings.Join([]string{"ApplicationV3", string(data)}, " ")
}

type ApplicationV3DeployType struct {
	value string
}

type ApplicationV3DeployTypeEnum struct {
	FUNCTION             ApplicationV3DeployType
	CCI                  ApplicationV3DeployType
	SERVICE_STAGE_JAR    ApplicationV3DeployType
	SERVICE_STAGE_DOCKER ApplicationV3DeployType
}

func GetApplicationV3DeployTypeEnum() ApplicationV3DeployTypeEnum {
	return ApplicationV3DeployTypeEnum{
		FUNCTION: ApplicationV3DeployType{
			value: "function",
		},
		CCI: ApplicationV3DeployType{
			value: "cci",
		},
		SERVICE_STAGE_JAR: ApplicationV3DeployType{
			value: "ServiceStage(Jar)",
		},
		SERVICE_STAGE_DOCKER: ApplicationV3DeployType{
			value: "ServiceStage(Docker)",
		},
	}
}

func (c ApplicationV3DeployType) Value() string {
	return c.value
}

func (c ApplicationV3DeployType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApplicationV3DeployType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
