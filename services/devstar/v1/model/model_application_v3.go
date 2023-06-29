package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApplicationV3 应用详细信息
type ApplicationV3 struct {

	// 应用id
	Id *string `json:"id,omitempty"`

	// 应用名称
	Name string `json:"name"`

	// 应用描述
	Description *string `json:"description,omitempty"`

	// 区域id
	RegionId string `json:"region_id"`

	// 区域名称
	RegionName string `json:"region_name"`

	// 所属项目id
	ProjectId string `json:"project_id"`

	// 项目名称
	ProjectName string `json:"project_name"`

	// 应用图标
	Icon *string `json:"icon,omitempty"`

	// 流水线列表
	PipelineCreationResult *[]PipelineCreationResult `json:"pipeline_creation_result,omitempty"`

	RepositoryCreationResult *RepositoryCreationResult `json:"repository_creation_result,omitempty"`

	// 环境信息
	EnvironmentCreationResult *[]string `json:"environment_creation_result,omitempty"`

	// 模板类型
	TemplateTypes *[]TemplateType `json:"template_types,omitempty"`

	// 模板部署信息
	TemplateDeployment *string `json:"template_deployment,omitempty"`

	// 部署类型, function:函数部署,cci:cci容器部署,ServiceStage(Jar):ServiceStage jar包部署,ServiceStage(Docker):ServiceStage Docker容器部署,none不支持部署
	DeployType *ApplicationV3DeployType `json:"deploy_type,omitempty"`

	// 创建者名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 应用版本号
	Version *string `json:"version,omitempty"`
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
	NONE                 ApplicationV3DeployType
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
		NONE: ApplicationV3DeployType{
			value: "none",
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
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
