package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/csms/v1/model"
)

type CsmsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCsmsClient(hcClient *http_client.HcHttpClient) *CsmsClient {
	return &CsmsClient{HcClient: hcClient}
}

func CsmsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//- 功能介绍：批量添加或删除凭据标签。
func (c *CsmsClient) BatchCreateOrDeleteTags(request *model.BatchCreateOrDeleteTagsRequest) (*model.BatchCreateOrDeleteTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateOrDeleteTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateOrDeleteTagsResponse), nil
	}
}

//创建新的凭据，并将凭据值存入凭据的初始版本。  凭据管理服务将凭据值加密后，存储在凭据对象下的版本中。每个版本可与多个凭据版本状态相关联，凭据版本状态用于标识凭据版本处于的阶段，没有版本状态标记的版本视为已弃用，可用凭据管理服务自动删除。  初始版本的状态被标记为SYSCURRENT。
func (c *CsmsClient) CreateSecret(request *model.CreateSecretRequest) (*model.CreateSecretResponse, error) {
	requestDef := GenReqDefForCreateSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecretResponse), nil
	}
}

//- 功能介绍：添加凭据标签。
func (c *CsmsClient) CreateSecretTag(request *model.CreateSecretTagRequest) (*model.CreateSecretTagResponse, error) {
	requestDef := GenReqDefForCreateSecretTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecretTagResponse), nil
	}
}

//在指定的凭据中，创建一个新的凭据版本，用于加密保管新的凭据值。默认情况下，新创建的凭据版本被标记为SYSCURRENT状态，而SYSCURRENT标记的前一个凭据版本被标记为SYSPREVIOUS状态。您可以通过指定VersionStage参数来覆盖默认行为。
func (c *CsmsClient) CreateSecretVersion(request *model.CreateSecretVersionRequest) (*model.CreateSecretVersionResponse, error) {
	requestDef := GenReqDefForCreateSecretVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecretVersionResponse), nil
	}
}

//立即删除指定的凭据，且无法恢复。
func (c *CsmsClient) DeleteSecret(request *model.DeleteSecretRequest) (*model.DeleteSecretResponse, error) {
	requestDef := GenReqDefForDeleteSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecretResponse), nil
	}
}

//指定延迟删除时间，创建删除凭据的定时任务，可设置7~30天的的延迟删除时间。
func (c *CsmsClient) DeleteSecretForSchedule(request *model.DeleteSecretForScheduleRequest) (*model.DeleteSecretForScheduleResponse, error) {
	requestDef := GenReqDefForDeleteSecretForSchedule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecretForScheduleResponse), nil
	}
}

//删除指定的凭据版本状态。
func (c *CsmsClient) DeleteSecretStage(request *model.DeleteSecretStageRequest) (*model.DeleteSecretStageResponse, error) {
	requestDef := GenReqDefForDeleteSecretStage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecretStageResponse), nil
	}
}

//- 功能介绍：删除凭据标签。
func (c *CsmsClient) DeleteSecretTag(request *model.DeleteSecretTagRequest) (*model.DeleteSecretTagResponse, error) {
	requestDef := GenReqDefForDeleteSecretTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecretTagResponse), nil
	}
}

//- 功能介绍：查询用户在指定项目下的所有凭据标签集合。
func (c *CsmsClient) ListProjectSecretsTags(request *model.ListProjectSecretsTagsRequest) (*model.ListProjectSecretsTagsResponse, error) {
	requestDef := GenReqDefForListProjectSecretsTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectSecretsTagsResponse), nil
	}
}

//- 功能介绍：查询凭据实例。通过标签过滤，筛选用户凭据,返回凭据列表。
func (c *CsmsClient) ListResourceInstances(request *model.ListResourceInstancesRequest) (*model.ListResourceInstancesResponse, error) {
	requestDef := GenReqDefForListResourceInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceInstancesResponse), nil
	}
}

//- 功能介绍：查询凭据标签。
func (c *CsmsClient) ListSecretTags(request *model.ListSecretTagsRequest) (*model.ListSecretTagsResponse, error) {
	requestDef := GenReqDefForListSecretTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecretTagsResponse), nil
	}
}

//查询指定凭据下的版本列表信息。
func (c *CsmsClient) ListSecretVersions(request *model.ListSecretVersionsRequest) (*model.ListSecretVersionsResponse, error) {
	requestDef := GenReqDefForListSecretVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecretVersionsResponse), nil
	}
}

//查询当前用户在本项目下创建的所有凭据。
func (c *CsmsClient) ListSecrets(request *model.ListSecretsRequest) (*model.ListSecretsResponse, error) {
	requestDef := GenReqDefForListSecrets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecretsResponse), nil
	}
}

//取消凭据的定时删除任务，凭据对象恢复可使用状态。
func (c *CsmsClient) RestoreSecret(request *model.RestoreSecretRequest) (*model.RestoreSecretResponse, error) {
	requestDef := GenReqDefForRestoreSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreSecretResponse), nil
	}
}

//查询指定凭据的信息。
func (c *CsmsClient) ShowSecret(request *model.ShowSecretRequest) (*model.ShowSecretResponse, error) {
	requestDef := GenReqDefForShowSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecretResponse), nil
	}
}

//查询指定凭据版本状态标记的版本信息。
func (c *CsmsClient) ShowSecretStage(request *model.ShowSecretStageRequest) (*model.ShowSecretStageResponse, error) {
	requestDef := GenReqDefForShowSecretStage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecretStageResponse), nil
	}
}

//查询指定凭据版本的信息和版本中的明文凭据值，只能查询ENABLED状态的凭据。 通过/v1/{project_id}/secrets/{secret_id}/versions/latest可访问凭据最新版本的凭据值。
func (c *CsmsClient) ShowSecretVersion(request *model.ShowSecretVersionRequest) (*model.ShowSecretVersionResponse, error) {
	requestDef := GenReqDefForShowSecretVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecretVersionResponse), nil
	}
}

//更新指定凭据的元数据信息。
func (c *CsmsClient) UpdateSecret(request *model.UpdateSecretRequest) (*model.UpdateSecretResponse, error) {
	requestDef := GenReqDefForUpdateSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecretResponse), nil
	}
}

//更新凭据的版本状态。
func (c *CsmsClient) UpdateSecretStage(request *model.UpdateSecretStageRequest) (*model.UpdateSecretStageResponse, error) {
	requestDef := GenReqDefForUpdateSecretStage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecretStageResponse), nil
	}
}
