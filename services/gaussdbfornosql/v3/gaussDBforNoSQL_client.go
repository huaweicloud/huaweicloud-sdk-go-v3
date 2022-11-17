package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/gaussdbfornosql/v3/model"
)

type GaussDBforNoSQLClient struct {
	HcClient *http_client.HcHttpClient
}

func NewGaussDBforNoSQLClient(hcClient *http_client.HcHttpClient) *GaussDBforNoSQLClient {
	return &GaussDBforNoSQLClient{HcClient: hcClient}
}

func GaussDBforNoSQLClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// ApplyConfiguration 应用参数模板
//
// 将参数模板应用到实例，可以指定一个或多个实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ApplyConfiguration(request *model.ApplyConfigurationRequest) (*model.ApplyConfigurationResponse, error) {
	requestDef := GenReqDefForApplyConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ApplyConfigurationResponse), nil
	}
}

// ApplyConfigurationInvoker 应用参数模板
func (c *GaussDBforNoSQLClient) ApplyConfigurationInvoker(request *model.ApplyConfigurationRequest) *ApplyConfigurationInvoker {
	requestDef := GenReqDefForApplyConfiguration()
	return &ApplyConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchTagAction 批量添加或删除资源标签
//
// 批量添加或删除指定数据库实例的标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) BatchTagAction(request *model.BatchTagActionRequest) (*model.BatchTagActionResponse, error) {
	requestDef := GenReqDefForBatchTagAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchTagActionResponse), nil
	}
}

// BatchTagActionInvoker 批量添加或删除资源标签
func (c *GaussDBforNoSQLClient) BatchTagActionInvoker(request *model.BatchTagActionRequest) *BatchTagActionInvoker {
	requestDef := GenReqDefForBatchTagAction()
	return &BatchTagActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckDisasterRecoveryOperation 校验实例是否可以与指定实例建立/解除容灾关系
//
// 校验实例是否可以与指定实例建立/解除容灾关系。若接口返回成功，表示可以与指定实例建立/解除容灾关系。
// 该接口需要对建立/解除容灾关系的两个实例各调用一次，2次调用都响应成功才能进行容灾关系的搭建/解除。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) CheckDisasterRecoveryOperation(request *model.CheckDisasterRecoveryOperationRequest) (*model.CheckDisasterRecoveryOperationResponse, error) {
	requestDef := GenReqDefForCheckDisasterRecoveryOperation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckDisasterRecoveryOperationResponse), nil
	}
}

// CheckDisasterRecoveryOperationInvoker 校验实例是否可以与指定实例建立/解除容灾关系
func (c *GaussDBforNoSQLClient) CheckDisasterRecoveryOperationInvoker(request *model.CheckDisasterRecoveryOperationRequest) *CheckDisasterRecoveryOperationInvoker {
	requestDef := GenReqDefForCheckDisasterRecoveryOperation()
	return &CheckDisasterRecoveryOperationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConfiguration 创建参数模板
//
// 创建参数模板。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) CreateConfiguration(request *model.CreateConfigurationRequest) (*model.CreateConfigurationResponse, error) {
	requestDef := GenReqDefForCreateConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConfigurationResponse), nil
	}
}

// CreateConfigurationInvoker 创建参数模板
func (c *GaussDBforNoSQLClient) CreateConfigurationInvoker(request *model.CreateConfigurationRequest) *CreateConfigurationInvoker {
	requestDef := GenReqDefForCreateConfiguration()
	return &CreateConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDisasterRecovery 搭建实例与特定实例的容灾关系
//
// 搭建实例与特定实例的容灾关系。 该接口需要对搭建容灾关系的两个实例分别各调用一次，2次接口都调用成功才能成功搭建容灾关系。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) CreateDisasterRecovery(request *model.CreateDisasterRecoveryRequest) (*model.CreateDisasterRecoveryResponse, error) {
	requestDef := GenReqDefForCreateDisasterRecovery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDisasterRecoveryResponse), nil
	}
}

// CreateDisasterRecoveryInvoker 搭建实例与特定实例的容灾关系
func (c *GaussDBforNoSQLClient) CreateDisasterRecoveryInvoker(request *model.CreateDisasterRecoveryRequest) *CreateDisasterRecoveryInvoker {
	requestDef := GenReqDefForCreateDisasterRecovery()
	return &CreateDisasterRecoveryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstance 创建实例
//
// 创建数据库实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResponse), nil
	}
}

// CreateInstanceInvoker 创建实例
func (c *GaussDBforNoSQLClient) CreateInstanceInvoker(request *model.CreateInstanceRequest) *CreateInstanceInvoker {
	requestDef := GenReqDefForCreateInstance()
	return &CreateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBackup 删除手动备份
//
// 删除手动备份
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) DeleteBackup(request *model.DeleteBackupRequest) (*model.DeleteBackupResponse, error) {
	requestDef := GenReqDefForDeleteBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackupResponse), nil
	}
}

// DeleteBackupInvoker 删除手动备份
func (c *GaussDBforNoSQLClient) DeleteBackupInvoker(request *model.DeleteBackupRequest) *DeleteBackupInvoker {
	requestDef := GenReqDefForDeleteBackup()
	return &DeleteBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteConfiguration 删除参数模板
//
// 删除指定参数模板。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) DeleteConfiguration(request *model.DeleteConfigurationRequest) (*model.DeleteConfigurationResponse, error) {
	requestDef := GenReqDefForDeleteConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConfigurationResponse), nil
	}
}

// DeleteConfigurationInvoker 删除参数模板
func (c *GaussDBforNoSQLClient) DeleteConfigurationInvoker(request *model.DeleteConfigurationRequest) *DeleteConfigurationInvoker {
	requestDef := GenReqDefForDeleteConfiguration()
	return &DeleteConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDisasterRecovery 解除实例与特定实例的容灾关系
//
// 解除实例与特定实例的容灾关系。 该接口需要对搭建容灾关系的两个实例分别各调用一次，2次接口都调用成功才能成功解除容灾关系。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) DeleteDisasterRecovery(request *model.DeleteDisasterRecoveryRequest) (*model.DeleteDisasterRecoveryResponse, error) {
	requestDef := GenReqDefForDeleteDisasterRecovery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDisasterRecoveryResponse), nil
	}
}

// DeleteDisasterRecoveryInvoker 解除实例与特定实例的容灾关系
func (c *GaussDBforNoSQLClient) DeleteDisasterRecoveryInvoker(request *model.DeleteDisasterRecoveryRequest) *DeleteDisasterRecoveryInvoker {
	requestDef := GenReqDefForDeleteDisasterRecovery()
	return &DeleteDisasterRecoveryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstance 删除实例
//
// 删除数据库实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

// DeleteInstanceInvoker 删除实例
func (c *GaussDBforNoSQLClient) DeleteInstanceInvoker(request *model.DeleteInstanceRequest) *DeleteInstanceInvoker {
	requestDef := GenReqDefForDeleteInstance()
	return &DeleteInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExpandInstanceNode 扩容指定集群实例的节点数量
//
// 扩容指定集群实例的节点数量。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ExpandInstanceNode(request *model.ExpandInstanceNodeRequest) (*model.ExpandInstanceNodeResponse, error) {
	requestDef := GenReqDefForExpandInstanceNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandInstanceNodeResponse), nil
	}
}

// ExpandInstanceNodeInvoker 扩容指定集群实例的节点数量
func (c *GaussDBforNoSQLClient) ExpandInstanceNodeInvoker(request *model.ExpandInstanceNodeRequest) *ExpandInstanceNodeInvoker {
	requestDef := GenReqDefForExpandInstanceNode()
	return &ExpandInstanceNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigurationTemplates 获取参数模板列表
//
// 获取参数模板列表，包括所有数据库的默认参数模板和用户创建的参数模板。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListConfigurationTemplates(request *model.ListConfigurationTemplatesRequest) (*model.ListConfigurationTemplatesResponse, error) {
	requestDef := GenReqDefForListConfigurationTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigurationTemplatesResponse), nil
	}
}

// ListConfigurationTemplatesInvoker 获取参数模板列表
func (c *GaussDBforNoSQLClient) ListConfigurationTemplatesInvoker(request *model.ListConfigurationTemplatesRequest) *ListConfigurationTemplatesInvoker {
	requestDef := GenReqDefForListConfigurationTemplates()
	return &ListConfigurationTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigurations 获取参数模板列表
//
// 获取参数模板列表，包括所有数据库的默认参数模板和用户创建的参数模板。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListConfigurations(request *model.ListConfigurationsRequest) (*model.ListConfigurationsResponse, error) {
	requestDef := GenReqDefForListConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigurationsResponse), nil
	}
}

// ListConfigurationsInvoker 获取参数模板列表
func (c *GaussDBforNoSQLClient) ListConfigurationsInvoker(request *model.ListConfigurationsRequest) *ListConfigurationsInvoker {
	requestDef := GenReqDefForListConfigurations()
	return &ListConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatastores 查询指定实例类型的数据库版本信息
//
// 查询指定实例类型的数据库版本信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListDatastores(request *model.ListDatastoresRequest) (*model.ListDatastoresResponse, error) {
	requestDef := GenReqDefForListDatastores()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatastoresResponse), nil
	}
}

// ListDatastoresInvoker 查询指定实例类型的数据库版本信息
func (c *GaussDBforNoSQLClient) ListDatastoresInvoker(request *model.ListDatastoresRequest) *ListDatastoresInvoker {
	requestDef := GenReqDefForListDatastores()
	return &ListDatastoresInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDedicatedResources 查询专属资源列表
//
// 查询专属资源列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListDedicatedResources(request *model.ListDedicatedResourcesRequest) (*model.ListDedicatedResourcesResponse, error) {
	requestDef := GenReqDefForListDedicatedResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDedicatedResourcesResponse), nil
	}
}

// ListDedicatedResourcesInvoker 查询专属资源列表
func (c *GaussDBforNoSQLClient) ListDedicatedResourcesInvoker(request *model.ListDedicatedResourcesRequest) *ListDedicatedResourcesInvoker {
	requestDef := GenReqDefForListDedicatedResources()
	return &ListDedicatedResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlavorInfos 查询数据库规格
//
// 查询指定条件下的实例规格信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListFlavorInfos(request *model.ListFlavorInfosRequest) (*model.ListFlavorInfosResponse, error) {
	requestDef := GenReqDefForListFlavorInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorInfosResponse), nil
	}
}

// ListFlavorInfosInvoker 查询数据库规格
func (c *GaussDBforNoSQLClient) ListFlavorInfosInvoker(request *model.ListFlavorInfosRequest) *ListFlavorInfosInvoker {
	requestDef := GenReqDefForListFlavorInfos()
	return &ListFlavorInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlavors 查询指定条件下的所有实例规格信息
//
// 查询指定条件下的所有实例规格信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

// ListFlavorsInvoker 查询指定条件下的所有实例规格信息
func (c *GaussDBforNoSQLClient) ListFlavorsInvoker(request *model.ListFlavorsRequest) *ListFlavorsInvoker {
	requestDef := GenReqDefForListFlavors()
	return &ListFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceTags 查询资源标签
//
// 查询指定实例的标签信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListInstanceTags(request *model.ListInstanceTagsRequest) (*model.ListInstanceTagsResponse, error) {
	requestDef := GenReqDefForListInstanceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceTagsResponse), nil
	}
}

// ListInstanceTagsInvoker 查询资源标签
func (c *GaussDBforNoSQLClient) ListInstanceTagsInvoker(request *model.ListInstanceTagsRequest) *ListInstanceTagsInvoker {
	requestDef := GenReqDefForListInstanceTags()
	return &ListInstanceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstances 查询实例列表和详情
//
// 根据指定条件查询数据库实例列表和详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// ListInstancesInvoker 查询实例列表和详情
func (c *GaussDBforNoSQLClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstancesByResourceTags 查询资源实例
//
// 根据标签查询指定的数据库实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListInstancesByResourceTags(request *model.ListInstancesByResourceTagsRequest) (*model.ListInstancesByResourceTagsResponse, error) {
	requestDef := GenReqDefForListInstancesByResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesByResourceTagsResponse), nil
	}
}

// ListInstancesByResourceTagsInvoker 查询资源实例
func (c *GaussDBforNoSQLClient) ListInstancesByResourceTagsInvoker(request *model.ListInstancesByResourceTagsRequest) *ListInstancesByResourceTagsInvoker {
	requestDef := GenReqDefForListInstancesByResourceTags()
	return &ListInstancesByResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstancesByTags 查询资源实例
//
// 根据标签查询指定的数据库实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListInstancesByTags(request *model.ListInstancesByTagsRequest) (*model.ListInstancesByTagsResponse, error) {
	requestDef := GenReqDefForListInstancesByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesByTagsResponse), nil
	}
}

// ListInstancesByTagsInvoker 查询资源实例
func (c *GaussDBforNoSQLClient) ListInstancesByTagsInvoker(request *model.ListInstancesByTagsRequest) *ListInstancesByTagsInvoker {
	requestDef := GenReqDefForListInstancesByTags()
	return &ListInstancesByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRestoreTime 查询实例可恢复的时间段
//
// 查询实例可恢复的时间段
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListRestoreTime(request *model.ListRestoreTimeRequest) (*model.ListRestoreTimeResponse, error) {
	requestDef := GenReqDefForListRestoreTime()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRestoreTimeResponse), nil
	}
}

// ListRestoreTimeInvoker 查询实例可恢复的时间段
func (c *GaussDBforNoSQLClient) ListRestoreTimeInvoker(request *model.ListRestoreTimeRequest) *ListRestoreTimeInvoker {
	requestDef := GenReqDefForListRestoreTime()
	return &ListRestoreTimeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSlowLogs 查询数据库慢日志
//
// 查询数据库慢日志信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListSlowLogs(request *model.ListSlowLogsRequest) (*model.ListSlowLogsResponse, error) {
	requestDef := GenReqDefForListSlowLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSlowLogsResponse), nil
	}
}

// ListSlowLogsInvoker 查询数据库慢日志
func (c *GaussDBforNoSQLClient) ListSlowLogsInvoker(request *model.ListSlowLogsRequest) *ListSlowLogsInvoker {
	requestDef := GenReqDefForListSlowLogs()
	return &ListSlowLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetPassword 修改实例的管理员密码
//
// 修改实例的管理员密码。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ResetPassword(request *model.ResetPasswordRequest) (*model.ResetPasswordResponse, error) {
	requestDef := GenReqDefForResetPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetPasswordResponse), nil
	}
}

// ResetPasswordInvoker 修改实例的管理员密码
func (c *GaussDBforNoSQLClient) ResetPasswordInvoker(request *model.ResetPasswordRequest) *ResetPasswordInvoker {
	requestDef := GenReqDefForResetPassword()
	return &ResetPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeInstance 变更实例规格
//
// 变更实例的规格。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ResizeInstance(request *model.ResizeInstanceRequest) (*model.ResizeInstanceResponse, error) {
	requestDef := GenReqDefForResizeInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeInstanceResponse), nil
	}
}

// ResizeInstanceInvoker 变更实例规格
func (c *GaussDBforNoSQLClient) ResizeInstanceInvoker(request *model.ResizeInstanceRequest) *ResizeInstanceInvoker {
	requestDef := GenReqDefForResizeInstance()
	return &ResizeInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeInstanceVolume 扩容实例存储容量
//
// 扩容实例的存储容量大小。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ResizeInstanceVolume(request *model.ResizeInstanceVolumeRequest) (*model.ResizeInstanceVolumeResponse, error) {
	requestDef := GenReqDefForResizeInstanceVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeInstanceVolumeResponse), nil
	}
}

// ResizeInstanceVolumeInvoker 扩容实例存储容量
func (c *GaussDBforNoSQLClient) ResizeInstanceVolumeInvoker(request *model.ResizeInstanceVolumeRequest) *ResizeInstanceVolumeInvoker {
	requestDef := GenReqDefForResizeInstanceVolume()
	return &ResizeInstanceVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreExistingInstance 恢复到已有实例
//
// 恢复到已有实例
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) RestoreExistingInstance(request *model.RestoreExistingInstanceRequest) (*model.RestoreExistingInstanceResponse, error) {
	requestDef := GenReqDefForRestoreExistingInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreExistingInstanceResponse), nil
	}
}

// RestoreExistingInstanceInvoker 恢复到已有实例
func (c *GaussDBforNoSQLClient) RestoreExistingInstanceInvoker(request *model.RestoreExistingInstanceRequest) *RestoreExistingInstanceInvoker {
	requestDef := GenReqDefForRestoreExistingInstance()
	return &RestoreExistingInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetBackupPolicy 设置自动备份策略
//
// 设置自动备份策略。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) SetBackupPolicy(request *model.SetBackupPolicyRequest) (*model.SetBackupPolicyResponse, error) {
	requestDef := GenReqDefForSetBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetBackupPolicyResponse), nil
	}
}

// SetBackupPolicyInvoker 设置自动备份策略
func (c *GaussDBforNoSQLClient) SetBackupPolicyInvoker(request *model.SetBackupPolicyRequest) *SetBackupPolicyInvoker {
	requestDef := GenReqDefForSetBackupPolicy()
	return &SetBackupPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBackupPolicy 查询自动备份策略
//
// 查询自动备份策略。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShowBackupPolicy(request *model.ShowBackupPolicyRequest) (*model.ShowBackupPolicyResponse, error) {
	requestDef := GenReqDefForShowBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackupPolicyResponse), nil
	}
}

// ShowBackupPolicyInvoker 查询自动备份策略
func (c *GaussDBforNoSQLClient) ShowBackupPolicyInvoker(request *model.ShowBackupPolicyRequest) *ShowBackupPolicyInvoker {
	requestDef := GenReqDefForShowBackupPolicy()
	return &ShowBackupPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConfigurationDetail 获取指定参数模板的参数
//
// 获取指定参数模板的详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShowConfigurationDetail(request *model.ShowConfigurationDetailRequest) (*model.ShowConfigurationDetailResponse, error) {
	requestDef := GenReqDefForShowConfigurationDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfigurationDetailResponse), nil
	}
}

// ShowConfigurationDetailInvoker 获取指定参数模板的参数
func (c *GaussDBforNoSQLClient) ShowConfigurationDetailInvoker(request *model.ShowConfigurationDetailRequest) *ShowConfigurationDetailInvoker {
	requestDef := GenReqDefForShowConfigurationDetail()
	return &ShowConfigurationDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceConfiguration 获取指定实例的参数
//
// 获取指定实例的参数信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShowInstanceConfiguration(request *model.ShowInstanceConfigurationRequest) (*model.ShowInstanceConfigurationResponse, error) {
	requestDef := GenReqDefForShowInstanceConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceConfigurationResponse), nil
	}
}

// ShowInstanceConfigurationInvoker 获取指定实例的参数
func (c *GaussDBforNoSQLClient) ShowInstanceConfigurationInvoker(request *model.ShowInstanceConfigurationRequest) *ShowInstanceConfigurationInvoker {
	requestDef := GenReqDefForShowInstanceConfiguration()
	return &ShowInstanceConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceRole 获取容灾实例主/备角色信息
//
// 该接口用于获取容灾实例主/备角色信息，以便后续容灾实例备升主和容灾实例主降备接口调用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShowInstanceRole(request *model.ShowInstanceRoleRequest) (*model.ShowInstanceRoleResponse, error) {
	requestDef := GenReqDefForShowInstanceRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceRoleResponse), nil
	}
}

// ShowInstanceRoleInvoker 获取容灾实例主/备角色信息
func (c *GaussDBforNoSQLClient) ShowInstanceRoleInvoker(request *model.ShowInstanceRoleRequest) *ShowInstanceRoleInvoker {
	requestDef := GenReqDefForShowInstanceRole()
	return &ShowInstanceRoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuotas 查询配额
//
// 查询单租户在GaussDBforNoSQL服务下的资源配额。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShowQuotas(request *model.ShowQuotasRequest) (*model.ShowQuotasResponse, error) {
	requestDef := GenReqDefForShowQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotasResponse), nil
	}
}

// ShowQuotasInvoker 查询配额
func (c *GaussDBforNoSQLClient) ShowQuotasInvoker(request *model.ShowQuotasRequest) *ShowQuotasInvoker {
	requestDef := GenReqDefForShowQuotas()
	return &ShowQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRestorableList 查询可恢复的实例列表
//
// 查询用户可恢复的实例列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShowRestorableList(request *model.ShowRestorableListRequest) (*model.ShowRestorableListResponse, error) {
	requestDef := GenReqDefForShowRestorableList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRestorableListResponse), nil
	}
}

// ShowRestorableListInvoker 查询可恢复的实例列表
func (c *GaussDBforNoSQLClient) ShowRestorableListInvoker(request *model.ShowRestorableListRequest) *ShowRestorableListInvoker {
	requestDef := GenReqDefForShowRestorableList()
	return &ShowRestorableListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShrinkInstanceNode 缩容指定集群实例的节点数量
//
// 缩容指定集群实例的节点数量。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShrinkInstanceNode(request *model.ShrinkInstanceNodeRequest) (*model.ShrinkInstanceNodeResponse, error) {
	requestDef := GenReqDefForShrinkInstanceNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShrinkInstanceNodeResponse), nil
	}
}

// ShrinkInstanceNodeInvoker 缩容指定集群实例的节点数量
func (c *GaussDBforNoSQLClient) ShrinkInstanceNodeInvoker(request *model.ShrinkInstanceNodeRequest) *ShrinkInstanceNodeInvoker {
	requestDef := GenReqDefForShrinkInstanceNode()
	return &ShrinkInstanceNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchToMaster 容灾实例备升主
//
// 该接口用于对已经搭建容灾关系的实例，将备实例升级为主实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) SwitchToMaster(request *model.SwitchToMasterRequest) (*model.SwitchToMasterResponse, error) {
	requestDef := GenReqDefForSwitchToMaster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchToMasterResponse), nil
	}
}

// SwitchToMasterInvoker 容灾实例备升主
func (c *GaussDBforNoSQLClient) SwitchToMasterInvoker(request *model.SwitchToMasterRequest) *SwitchToMasterInvoker {
	requestDef := GenReqDefForSwitchToMaster()
	return &SwitchToMasterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchToSlave 容灾实例主降备
//
// 该接口用于对已经搭建容灾关系的实例，将主实例降级为备实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) SwitchToSlave(request *model.SwitchToSlaveRequest) (*model.SwitchToSlaveResponse, error) {
	requestDef := GenReqDefForSwitchToSlave()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchToSlaveResponse), nil
	}
}

// SwitchToSlaveInvoker 容灾实例主降备
func (c *GaussDBforNoSQLClient) SwitchToSlaveInvoker(request *model.SwitchToSlaveRequest) *SwitchToSlaveInvoker {
	requestDef := GenReqDefForSwitchToSlave()
	return &SwitchToSlaveInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateConfiguration 修改参数模板参数
//
// 修改参数模板参数。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) UpdateConfiguration(request *model.UpdateConfigurationRequest) (*model.UpdateConfigurationResponse, error) {
	requestDef := GenReqDefForUpdateConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConfigurationResponse), nil
	}
}

// UpdateConfigurationInvoker 修改参数模板参数
func (c *GaussDBforNoSQLClient) UpdateConfigurationInvoker(request *model.UpdateConfigurationRequest) *UpdateConfigurationInvoker {
	requestDef := GenReqDefForUpdateConfiguration()
	return &UpdateConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceConfiguration 修改指定实例的参数
//
// 修改指定实例的参数。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) UpdateInstanceConfiguration(request *model.UpdateInstanceConfigurationRequest) (*model.UpdateInstanceConfigurationResponse, error) {
	requestDef := GenReqDefForUpdateInstanceConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceConfigurationResponse), nil
	}
}

// UpdateInstanceConfigurationInvoker 修改指定实例的参数
func (c *GaussDBforNoSQLClient) UpdateInstanceConfigurationInvoker(request *model.UpdateInstanceConfigurationRequest) *UpdateInstanceConfigurationInvoker {
	requestDef := GenReqDefForUpdateInstanceConfiguration()
	return &UpdateInstanceConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceName 修改实例名称
//
// 修改实例名称
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) UpdateInstanceName(request *model.UpdateInstanceNameRequest) (*model.UpdateInstanceNameResponse, error) {
	requestDef := GenReqDefForUpdateInstanceName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceNameResponse), nil
	}
}

// UpdateInstanceNameInvoker 修改实例名称
func (c *GaussDBforNoSQLClient) UpdateInstanceNameInvoker(request *model.UpdateInstanceNameRequest) *UpdateInstanceNameInvoker {
	requestDef := GenReqDefForUpdateInstanceName()
	return &UpdateInstanceNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSecurityGroup 变更实例安全组
//
// 变更实例关联的安全组
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) UpdateSecurityGroup(request *model.UpdateSecurityGroupRequest) (*model.UpdateSecurityGroupResponse, error) {
	requestDef := GenReqDefForUpdateSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecurityGroupResponse), nil
	}
}

// UpdateSecurityGroupInvoker 变更实例安全组
func (c *GaussDBforNoSQLClient) UpdateSecurityGroupInvoker(request *model.UpdateSecurityGroupRequest) *UpdateSecurityGroupInvoker {
	requestDef := GenReqDefForUpdateSecurityGroup()
	return &UpdateSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpgradeDbVersion 数据库补丁升级
//
// 升级数据库补丁版本
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) UpgradeDbVersion(request *model.UpgradeDbVersionRequest) (*model.UpgradeDbVersionResponse, error) {
	requestDef := GenReqDefForUpgradeDbVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpgradeDbVersionResponse), nil
	}
}

// UpgradeDbVersionInvoker 数据库补丁升级
func (c *GaussDBforNoSQLClient) UpgradeDbVersionInvoker(request *model.UpgradeDbVersionRequest) *UpgradeDbVersionInvoker {
	requestDef := GenReqDefForUpgradeDbVersion()
	return &UpgradeDbVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiVersion 查询当前支持的API版本信息列表
//
// 查询当前支持的API版本信息列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListApiVersion(request *model.ListApiVersionRequest) (*model.ListApiVersionResponse, error) {
	requestDef := GenReqDefForListApiVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionResponse), nil
	}
}

// ListApiVersionInvoker 查询当前支持的API版本信息列表
func (c *GaussDBforNoSQLClient) ListApiVersionInvoker(request *model.ListApiVersionRequest) *ListApiVersionInvoker {
	requestDef := GenReqDefForListApiVersion()
	return &ListApiVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApiVersion 查询指定API版本信息
//
// 查询指定API版本信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShowApiVersion(request *model.ShowApiVersionRequest) (*model.ShowApiVersionResponse, error) {
	requestDef := GenReqDefForShowApiVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApiVersionResponse), nil
	}
}

// ShowApiVersionInvoker 查询指定API版本信息
func (c *GaussDBforNoSQLClient) ShowApiVersionInvoker(request *model.ShowApiVersionRequest) *ShowApiVersionInvoker {
	requestDef := GenReqDefForShowApiVersion()
	return &ShowApiVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
