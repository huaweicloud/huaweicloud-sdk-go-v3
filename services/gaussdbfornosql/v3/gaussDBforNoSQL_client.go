package v3

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/gaussdbfornosql/v3/model"
)

type GaussDBforNoSQLClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewGaussDBforNoSQLClient(hcClient *httpclient.HcHttpClient) *GaussDBforNoSQLClient {
	return &GaussDBforNoSQLClient{HcClient: hcClient}
}

func GaussDBforNoSQLClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// ApplyConfiguration 应用参数模板
//
// 将参数模板应用到实例，可以指定一个或多个实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// CheckWeekPassword 判断弱密码
//
// 判断弱密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) CheckWeekPassword(request *model.CheckWeekPasswordRequest) (*model.CheckWeekPasswordResponse, error) {
	requestDef := GenReqDefForCheckWeekPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckWeekPasswordResponse), nil
	}
}

// CheckWeekPasswordInvoker 判断弱密码
func (c *GaussDBforNoSQLClient) CheckWeekPasswordInvoker(request *model.CheckWeekPasswordRequest) *CheckWeekPasswordInvoker {
	requestDef := GenReqDefForCheckWeekPassword()
	return &CheckWeekPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CompareConfiguration 参数模板比较
//
// 比较两个参数模板之间的差异
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) CompareConfiguration(request *model.CompareConfigurationRequest) (*model.CompareConfigurationResponse, error) {
	requestDef := GenReqDefForCompareConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CompareConfigurationResponse), nil
	}
}

// CompareConfigurationInvoker 参数模板比较
func (c *GaussDBforNoSQLClient) CompareConfigurationInvoker(request *model.CompareConfigurationRequest) *CompareConfigurationInvoker {
	requestDef := GenReqDefForCompareConfiguration()
	return &CompareConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyConfiguration 复制参数模板
//
// 复制参数模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) CopyConfiguration(request *model.CopyConfigurationRequest) (*model.CopyConfigurationResponse, error) {
	requestDef := GenReqDefForCopyConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyConfigurationResponse), nil
	}
}

// CopyConfigurationInvoker 复制参数模板
func (c *GaussDBforNoSQLClient) CopyConfigurationInvoker(request *model.CopyConfigurationRequest) *CopyConfigurationInvoker {
	requestDef := GenReqDefForCopyConfiguration()
	return &CopyConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateBack 创建手动备份
//
// 创建手动备份。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) CreateBack(request *model.CreateBackRequest) (*model.CreateBackResponse, error) {
	requestDef := GenReqDefForCreateBack()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBackResponse), nil
	}
}

// CreateBackInvoker 创建手动备份
func (c *GaussDBforNoSQLClient) CreateBackInvoker(request *model.CreateBackRequest) *CreateBackInvoker {
	requestDef := GenReqDefForCreateBack()
	return &CreateBackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateColdVolume ‘创建冷数据存储’
//
// ‘创建冷数据存储’
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) CreateColdVolume(request *model.CreateColdVolumeRequest) (*model.CreateColdVolumeResponse, error) {
	requestDef := GenReqDefForCreateColdVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateColdVolumeResponse), nil
	}
}

// CreateColdVolumeInvoker ‘创建冷数据存储’
func (c *GaussDBforNoSQLClient) CreateColdVolumeInvoker(request *model.CreateColdVolumeRequest) *CreateColdVolumeInvoker {
	requestDef := GenReqDefForCreateColdVolume()
	return &CreateColdVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConfiguration 创建参数模板
//
// 创建参数模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// CreateDbUser 创建Redis数据库账号
//
// 在Redis实例中创建数据库帐号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) CreateDbUser(request *model.CreateDbUserRequest) (*model.CreateDbUserResponse, error) {
	requestDef := GenReqDefForCreateDbUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDbUserResponse), nil
	}
}

// CreateDbUserInvoker 创建Redis数据库账号
func (c *GaussDBforNoSQLClient) CreateDbUserInvoker(request *model.CreateDbUserRequest) *CreateDbUserInvoker {
	requestDef := GenReqDefForCreateDbUser()
	return &CreateDbUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDisasterRecovery 搭建实例与特定实例的容灾关系
//
// 搭建实例与特定实例的容灾关系。 该接口需要对搭建容灾关系的两个实例分别各调用一次，2次接口都调用成功才能成功搭建容灾关系。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// DeleteDbUser 删除Redis数据库账号
//
// 删除Redis实例的数据库账号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) DeleteDbUser(request *model.DeleteDbUserRequest) (*model.DeleteDbUserResponse, error) {
	requestDef := GenReqDefForDeleteDbUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDbUserResponse), nil
	}
}

// DeleteDbUserInvoker 删除Redis数据库账号
func (c *GaussDBforNoSQLClient) DeleteDbUserInvoker(request *model.DeleteDbUserRequest) *DeleteDbUserInvoker {
	requestDef := GenReqDefForDeleteDbUser()
	return &DeleteDbUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDisasterRecovery 解除实例与特定实例的容灾关系
//
// 解除实例与特定实例的容灾关系。 该接口需要对搭建容灾关系的两个实例分别各调用一次，2次接口都调用成功才能成功解除容灾关系。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// DeleteEnlargeFailNode 删除扩容失败的节点
//
// 删除扩容失败的节点
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) DeleteEnlargeFailNode(request *model.DeleteEnlargeFailNodeRequest) (*model.DeleteEnlargeFailNodeResponse, error) {
	requestDef := GenReqDefForDeleteEnlargeFailNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEnlargeFailNodeResponse), nil
	}
}

// DeleteEnlargeFailNodeInvoker 删除扩容失败的节点
func (c *GaussDBforNoSQLClient) DeleteEnlargeFailNodeInvoker(request *model.DeleteEnlargeFailNodeRequest) *DeleteEnlargeFailNodeInvoker {
	requestDef := GenReqDefForDeleteEnlargeFailNode()
	return &DeleteEnlargeFailNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstance 删除实例
//
// 删除数据库实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// DeleteInstancesSession 关闭实例节点会话
//
// 关闭实例节点会话。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) DeleteInstancesSession(request *model.DeleteInstancesSessionRequest) (*model.DeleteInstancesSessionResponse, error) {
	requestDef := GenReqDefForDeleteInstancesSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstancesSessionResponse), nil
	}
}

// DeleteInstancesSessionInvoker 关闭实例节点会话
func (c *GaussDBforNoSQLClient) DeleteInstancesSessionInvoker(request *model.DeleteInstancesSessionRequest) *DeleteInstancesSessionInvoker {
	requestDef := GenReqDefForDeleteInstancesSession()
	return &DeleteInstancesSessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLtsConfigs 解除关联LTS日志流
//
// 将实例日志与LTS日志流解除关联，后台将取消上传实例日志到的LTS日志流里。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) DeleteLtsConfigs(request *model.DeleteLtsConfigsRequest) (*model.DeleteLtsConfigsResponse, error) {
	requestDef := GenReqDefForDeleteLtsConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLtsConfigsResponse), nil
	}
}

// DeleteLtsConfigsInvoker 解除关联LTS日志流
func (c *GaussDBforNoSQLClient) DeleteLtsConfigsInvoker(request *model.DeleteLtsConfigsRequest) *DeleteLtsConfigsInvoker {
	requestDef := GenReqDefForDeleteLtsConfigs()
	return &DeleteLtsConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExpandInstanceNode 扩容指定集群实例的节点数量
//
// 扩容指定集群实例的节点数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListAvailableFlavorInfos 查询实例可变更规格
//
// 查询实例可变更规格。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListAvailableFlavorInfos(request *model.ListAvailableFlavorInfosRequest) (*model.ListAvailableFlavorInfosResponse, error) {
	requestDef := GenReqDefForListAvailableFlavorInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableFlavorInfosResponse), nil
	}
}

// ListAvailableFlavorInfosInvoker 查询实例可变更规格
func (c *GaussDBforNoSQLClient) ListAvailableFlavorInfosInvoker(request *model.ListAvailableFlavorInfosRequest) *ListAvailableFlavorInfosInvoker {
	requestDef := GenReqDefForListAvailableFlavorInfos()
	return &ListAvailableFlavorInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCassandraSlowLogs 查询GeminiDB(for Cassandra)数据库慢日志
//
// 查询GeminiDB(for Cassandra)数据库慢日志信息，支持日志关键字搜索。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListCassandraSlowLogs(request *model.ListCassandraSlowLogsRequest) (*model.ListCassandraSlowLogsResponse, error) {
	requestDef := GenReqDefForListCassandraSlowLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCassandraSlowLogsResponse), nil
	}
}

// ListCassandraSlowLogsInvoker 查询GeminiDB(for Cassandra)数据库慢日志
func (c *GaussDBforNoSQLClient) ListCassandraSlowLogsInvoker(request *model.ListCassandraSlowLogsRequest) *ListCassandraSlowLogsInvoker {
	requestDef := GenReqDefForListCassandraSlowLogs()
	return &ListCassandraSlowLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigurationDatastores 查询支持参数模板的引擎信息
//
// 查询支持参数模板的引擎信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListConfigurationDatastores(request *model.ListConfigurationDatastoresRequest) (*model.ListConfigurationDatastoresResponse, error) {
	requestDef := GenReqDefForListConfigurationDatastores()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigurationDatastoresResponse), nil
	}
}

// ListConfigurationDatastoresInvoker 查询支持参数模板的引擎信息
func (c *GaussDBforNoSQLClient) ListConfigurationDatastoresInvoker(request *model.ListConfigurationDatastoresRequest) *ListConfigurationDatastoresInvoker {
	requestDef := GenReqDefForListConfigurationDatastores()
	return &ListConfigurationDatastoresInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigurationTemplates 获取参数模板列表
//
// 获取参数模板列表，包括所有数据库的默认参数模板和用户创建的参数模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListDbUsers 获取Redis数据库账号列表和详情
//
// 获取Redis数据库账号列表和详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListDbUsers(request *model.ListDbUsersRequest) (*model.ListDbUsersResponse, error) {
	requestDef := GenReqDefForListDbUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDbUsersResponse), nil
	}
}

// ListDbUsersInvoker 获取Redis数据库账号列表和详情
func (c *GaussDBforNoSQLClient) ListDbUsersInvoker(request *model.ListDbUsersRequest) *ListDbUsersInvoker {
	requestDef := GenReqDefForListDbUsers()
	return &ListDbUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDedicatedResources 查询专属资源列表
//
// 查询专属资源列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListEpsQuotas 查询企业项目配额
//
// 查询企业项目配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListEpsQuotas(request *model.ListEpsQuotasRequest) (*model.ListEpsQuotasResponse, error) {
	requestDef := GenReqDefForListEpsQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEpsQuotasResponse), nil
	}
}

// ListEpsQuotasInvoker 查询企业项目配额
func (c *GaussDBforNoSQLClient) ListEpsQuotasInvoker(request *model.ListEpsQuotasRequest) *ListEpsQuotasInvoker {
	requestDef := GenReqDefForListEpsQuotas()
	return &ListEpsQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlavorInfos 查询数据库规格
//
// 查询指定条件下的实例规格信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListInfluxdbSlowLogs 查询GeminiDB(for influxdb)数据库慢日志
//
// 查询GeminiDB(for influxdb)数据库慢日志信息，支持日志关键字搜索。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListInfluxdbSlowLogs(request *model.ListInfluxdbSlowLogsRequest) (*model.ListInfluxdbSlowLogsResponse, error) {
	requestDef := GenReqDefForListInfluxdbSlowLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInfluxdbSlowLogsResponse), nil
	}
}

// ListInfluxdbSlowLogsInvoker 查询GeminiDB(for influxdb)数据库慢日志
func (c *GaussDBforNoSQLClient) ListInfluxdbSlowLogsInvoker(request *model.ListInfluxdbSlowLogsRequest) *ListInfluxdbSlowLogsInvoker {
	requestDef := GenReqDefForListInfluxdbSlowLogs()
	return &ListInfluxdbSlowLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceDatabases 获取Redis实例数据库列表
//
// 获取Redis实例数据库列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListInstanceDatabases(request *model.ListInstanceDatabasesRequest) (*model.ListInstanceDatabasesResponse, error) {
	requestDef := GenReqDefForListInstanceDatabases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceDatabasesResponse), nil
	}
}

// ListInstanceDatabasesInvoker 获取Redis实例数据库列表
func (c *GaussDBforNoSQLClient) ListInstanceDatabasesInvoker(request *model.ListInstanceDatabasesRequest) *ListInstanceDatabasesInvoker {
	requestDef := GenReqDefForListInstanceDatabases()
	return &ListInstanceDatabasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceTags 查询资源标签
//
// 查询指定实例的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListInstancesSession 获取节点会话列表
//
// 获取节点会话列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListInstancesSession(request *model.ListInstancesSessionRequest) (*model.ListInstancesSessionResponse, error) {
	requestDef := GenReqDefForListInstancesSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesSessionResponse), nil
	}
}

// ListInstancesSessionInvoker 获取节点会话列表
func (c *GaussDBforNoSQLClient) ListInstancesSessionInvoker(request *model.ListInstancesSessionRequest) *ListInstancesSessionInvoker {
	requestDef := GenReqDefForListInstancesSession()
	return &ListInstancesSessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstancesSessionStatistics 查询实例节点会话统计信息
//
// 查询实例节点会话统计信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListInstancesSessionStatistics(request *model.ListInstancesSessionStatisticsRequest) (*model.ListInstancesSessionStatisticsResponse, error) {
	requestDef := GenReqDefForListInstancesSessionStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesSessionStatisticsResponse), nil
	}
}

// ListInstancesSessionStatisticsInvoker 查询实例节点会话统计信息
func (c *GaussDBforNoSQLClient) ListInstancesSessionStatisticsInvoker(request *model.ListInstancesSessionStatisticsRequest) *ListInstancesSessionStatisticsInvoker {
	requestDef := GenReqDefForListInstancesSessionStatistics()
	return &ListInstancesSessionStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJobs 查询任务列表和详情
//
// 查询任务列表和详情，默认查询任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListJobs(request *model.ListJobsRequest) (*model.ListJobsResponse, error) {
	requestDef := GenReqDefForListJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobsResponse), nil
	}
}

// ListJobsInvoker 查询任务列表和详情
func (c *GaussDBforNoSQLClient) ListJobsInvoker(request *model.ListJobsRequest) *ListJobsInvoker {
	requestDef := GenReqDefForListJobs()
	return &ListJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLtsConfigs 查询LTS日志配置信息
//
// 分页查询实例关联的LTS日志配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListLtsConfigs(request *model.ListLtsConfigsRequest) (*model.ListLtsConfigsResponse, error) {
	requestDef := GenReqDefForListLtsConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLtsConfigsResponse), nil
	}
}

// ListLtsConfigsInvoker 查询LTS日志配置信息
func (c *GaussDBforNoSQLClient) ListLtsConfigsInvoker(request *model.ListLtsConfigsRequest) *ListLtsConfigsInvoker {
	requestDef := GenReqDefForListLtsConfigs()
	return &ListLtsConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMongodbErrorLogs 查询GeminiDB(for Mongo)数据库错误日志
//
// 查询GeminiDB(for Mongo)数据库错误日志信息，支持日志关键字搜索。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListMongodbErrorLogs(request *model.ListMongodbErrorLogsRequest) (*model.ListMongodbErrorLogsResponse, error) {
	requestDef := GenReqDefForListMongodbErrorLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMongodbErrorLogsResponse), nil
	}
}

// ListMongodbErrorLogsInvoker 查询GeminiDB(for Mongo)数据库错误日志
func (c *GaussDBforNoSQLClient) ListMongodbErrorLogsInvoker(request *model.ListMongodbErrorLogsRequest) *ListMongodbErrorLogsInvoker {
	requestDef := GenReqDefForListMongodbErrorLogs()
	return &ListMongodbErrorLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMongodbSlowLogs 查询GeminiDB(for Mongo)数据库慢日志
//
// 查询GeminiDB(for Mongo)数据库慢日志信息，支持日志关键字搜索。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListMongodbSlowLogs(request *model.ListMongodbSlowLogsRequest) (*model.ListMongodbSlowLogsResponse, error) {
	requestDef := GenReqDefForListMongodbSlowLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMongodbSlowLogsResponse), nil
	}
}

// ListMongodbSlowLogsInvoker 查询GeminiDB(for Mongo)数据库慢日志
func (c *GaussDBforNoSQLClient) ListMongodbSlowLogsInvoker(request *model.ListMongodbSlowLogsRequest) *ListMongodbSlowLogsInvoker {
	requestDef := GenReqDefForListMongodbSlowLogs()
	return &ListMongodbSlowLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectTags 查询项目标签
//
// 查询指定项目的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListProjectTags(request *model.ListProjectTagsRequest) (*model.ListProjectTagsResponse, error) {
	requestDef := GenReqDefForListProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTagsResponse), nil
	}
}

// ListProjectTagsInvoker 查询项目标签
func (c *GaussDBforNoSQLClient) ListProjectTagsInvoker(request *model.ListProjectTagsRequest) *ListProjectTagsInvoker {
	requestDef := GenReqDefForListProjectTags()
	return &ListProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRecycleInstances 查询回收站实例列表
//
// 查询回收站所有引擎的实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListRecycleInstances(request *model.ListRecycleInstancesRequest) (*model.ListRecycleInstancesResponse, error) {
	requestDef := GenReqDefForListRecycleInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecycleInstancesResponse), nil
	}
}

// ListRecycleInstancesInvoker 查询回收站实例列表
func (c *GaussDBforNoSQLClient) ListRecycleInstancesInvoker(request *model.ListRecycleInstancesRequest) *ListRecycleInstancesInvoker {
	requestDef := GenReqDefForListRecycleInstances()
	return &ListRecycleInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRedisSlowLogs 查询GeminiDB(for Redis)数据库慢日志
//
// 查询GeminiDB(for Redis)数据库慢日志信息，支持日志关键字搜索。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListRedisSlowLogs(request *model.ListRedisSlowLogsRequest) (*model.ListRedisSlowLogsResponse, error) {
	requestDef := GenReqDefForListRedisSlowLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRedisSlowLogsResponse), nil
	}
}

// ListRedisSlowLogsInvoker 查询GeminiDB(for Redis)数据库慢日志
func (c *GaussDBforNoSQLClient) ListRedisSlowLogsInvoker(request *model.ListRedisSlowLogsRequest) *ListRedisSlowLogsInvoker {
	requestDef := GenReqDefForListRedisSlowLogs()
	return &ListRedisSlowLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRestoreDatabases 获取GeminiDB(for Cassandra)实例表级恢复的数据库信息
//
// 获取GeminiDB(for Cassandra)实例表级恢复的数据库信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListRestoreDatabases(request *model.ListRestoreDatabasesRequest) (*model.ListRestoreDatabasesResponse, error) {
	requestDef := GenReqDefForListRestoreDatabases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRestoreDatabasesResponse), nil
	}
}

// ListRestoreDatabasesInvoker 获取GeminiDB(for Cassandra)实例表级恢复的数据库信息
func (c *GaussDBforNoSQLClient) ListRestoreDatabasesInvoker(request *model.ListRestoreDatabasesRequest) *ListRestoreDatabasesInvoker {
	requestDef := GenReqDefForListRestoreDatabases()
	return &ListRestoreDatabasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRestoreTables 获取GeminiDB(for Cassandra)实例表级恢复的表信息
//
// 获取GeminiDB(for Cassandra)实例表级恢复的表信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ListRestoreTables(request *model.ListRestoreTablesRequest) (*model.ListRestoreTablesResponse, error) {
	requestDef := GenReqDefForListRestoreTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRestoreTablesResponse), nil
	}
}

// ListRestoreTablesInvoker 获取GeminiDB(for Cassandra)实例表级恢复的表信息
func (c *GaussDBforNoSQLClient) ListRestoreTablesInvoker(request *model.ListRestoreTablesRequest) *ListRestoreTablesInvoker {
	requestDef := GenReqDefForListRestoreTables()
	return &ListRestoreTablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRestoreTime 查询实例可恢复的时间段
//
// 查询实例可恢复的时间段
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ModifyDbUserPrivilege 修改Redis数据库帐号权限
//
// 修改Redis数据库帐号权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ModifyDbUserPrivilege(request *model.ModifyDbUserPrivilegeRequest) (*model.ModifyDbUserPrivilegeResponse, error) {
	requestDef := GenReqDefForModifyDbUserPrivilege()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyDbUserPrivilegeResponse), nil
	}
}

// ModifyDbUserPrivilegeInvoker 修改Redis数据库帐号权限
func (c *GaussDBforNoSQLClient) ModifyDbUserPrivilegeInvoker(request *model.ModifyDbUserPrivilegeRequest) *ModifyDbUserPrivilegeInvoker {
	requestDef := GenReqDefForModifyDbUserPrivilege()
	return &ModifyDbUserPrivilegeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyEpsQuotas 修改企业项目配额
//
// 修改企业项目配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ModifyEpsQuotas(request *model.ModifyEpsQuotasRequest) (*model.ModifyEpsQuotasResponse, error) {
	requestDef := GenReqDefForModifyEpsQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyEpsQuotasResponse), nil
	}
}

// ModifyEpsQuotasInvoker 修改企业项目配额
func (c *GaussDBforNoSQLClient) ModifyEpsQuotasInvoker(request *model.ModifyEpsQuotasRequest) *ModifyEpsQuotasInvoker {
	requestDef := GenReqDefForModifyEpsQuotas()
	return &ModifyEpsQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyPort 修改数据库端口
//
// 修改数据库实例的端口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ModifyPort(request *model.ModifyPortRequest) (*model.ModifyPortResponse, error) {
	requestDef := GenReqDefForModifyPort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyPortResponse), nil
	}
}

// ModifyPortInvoker 修改数据库端口
func (c *GaussDBforNoSQLClient) ModifyPortInvoker(request *model.ModifyPortRequest) *ModifyPortInvoker {
	requestDef := GenReqDefForModifyPort()
	return &ModifyPortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyPublicIp 绑定/解绑弹性公网IP
//
// 实例下的节点绑定弹性公网IP/解绑弹性公网IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ModifyPublicIp(request *model.ModifyPublicIpRequest) (*model.ModifyPublicIpResponse, error) {
	requestDef := GenReqDefForModifyPublicIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyPublicIpResponse), nil
	}
}

// ModifyPublicIpInvoker 绑定/解绑弹性公网IP
func (c *GaussDBforNoSQLClient) ModifyPublicIpInvoker(request *model.ModifyPublicIpRequest) *ModifyPublicIpInvoker {
	requestDef := GenReqDefForModifyPublicIp()
	return &ModifyPublicIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyVolume 变更实例存储容量
//
// 变更实例的存储容量大小
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ModifyVolume(request *model.ModifyVolumeRequest) (*model.ModifyVolumeResponse, error) {
	requestDef := GenReqDefForModifyVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyVolumeResponse), nil
	}
}

// ModifyVolumeInvoker 变更实例存储容量
func (c *GaussDBforNoSQLClient) ModifyVolumeInvoker(request *model.ModifyVolumeRequest) *ModifyVolumeInvoker {
	requestDef := GenReqDefForModifyVolume()
	return &ModifyVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// OfflineNodes 支持节点的开关机
//
// 当底层故障导致节点无法正常工作时，可以对该节点执行关机操作，关机后会由其他节点接管业务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) OfflineNodes(request *model.OfflineNodesRequest) (*model.OfflineNodesResponse, error) {
	requestDef := GenReqDefForOfflineNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.OfflineNodesResponse), nil
	}
}

// OfflineNodesInvoker 支持节点的开关机
func (c *GaussDBforNoSQLClient) OfflineNodesInvoker(request *model.OfflineNodesRequest) *OfflineNodesInvoker {
	requestDef := GenReqDefForOfflineNodes()
	return &OfflineNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PauseResumeDataSynchronization 暂停/恢复具备容灾关系的实例数据同步
//
// 该接口用于暂停/恢复具备容灾关系的实例数据同步。
//
// 该接口需要对具备容灾关系的两个实例分别各调用一次，2次接口都调用成功才能成功暂停/恢复容灾实例数据同步。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) PauseResumeDataSynchronization(request *model.PauseResumeDataSynchronizationRequest) (*model.PauseResumeDataSynchronizationResponse, error) {
	requestDef := GenReqDefForPauseResumeDataSynchronization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PauseResumeDataSynchronizationResponse), nil
	}
}

// PauseResumeDataSynchronizationInvoker 暂停/恢复具备容灾关系的实例数据同步
func (c *GaussDBforNoSQLClient) PauseResumeDataSynchronizationInvoker(request *model.PauseResumeDataSynchronizationRequest) *PauseResumeDataSynchronizationInvoker {
	requestDef := GenReqDefForPauseResumeDataSynchronization()
	return &PauseResumeDataSynchronizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetDbUserPassword 重置Redis数据库账号密码
//
// 重置Redis数据库账号密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ResetDbUserPassword(request *model.ResetDbUserPasswordRequest) (*model.ResetDbUserPasswordResponse, error) {
	requestDef := GenReqDefForResetDbUserPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetDbUserPasswordResponse), nil
	}
}

// ResetDbUserPasswordInvoker 重置Redis数据库账号密码
func (c *GaussDBforNoSQLClient) ResetDbUserPasswordInvoker(request *model.ResetDbUserPasswordRequest) *ResetDbUserPasswordInvoker {
	requestDef := GenReqDefForResetDbUserPassword()
	return &ResetDbUserPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetParamGroupTemplate 重置自定义参数模板
//
// 重置自定义参数模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ResetParamGroupTemplate(request *model.ResetParamGroupTemplateRequest) (*model.ResetParamGroupTemplateResponse, error) {
	requestDef := GenReqDefForResetParamGroupTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetParamGroupTemplateResponse), nil
	}
}

// ResetParamGroupTemplateInvoker 重置自定义参数模板
func (c *GaussDBforNoSQLClient) ResetParamGroupTemplateInvoker(request *model.ResetParamGroupTemplateRequest) *ResetParamGroupTemplateInvoker {
	requestDef := GenReqDefForResetParamGroupTemplate()
	return &ResetParamGroupTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetPassword 修改实例的管理员密码
//
// 修改实例的管理员密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ResizeColdVolume 扩容冷数据存储
//
// 扩容冷数据存储。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ResizeColdVolume(request *model.ResizeColdVolumeRequest) (*model.ResizeColdVolumeResponse, error) {
	requestDef := GenReqDefForResizeColdVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeColdVolumeResponse), nil
	}
}

// ResizeColdVolumeInvoker 扩容冷数据存储
func (c *GaussDBforNoSQLClient) ResizeColdVolumeInvoker(request *model.ResizeColdVolumeRequest) *ResizeColdVolumeInvoker {
	requestDef := GenReqDefForResizeColdVolume()
	return &ResizeColdVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeInstance 变更实例规格
//
// 变更实例的规格。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// RestartInstance 重启实例的数据库服务
//
// 重启实例的数据库服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) RestartInstance(request *model.RestartInstanceRequest) (*model.RestartInstanceResponse, error) {
	requestDef := GenReqDefForRestartInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartInstanceResponse), nil
	}
}

// RestartInstanceInvoker 重启实例的数据库服务
func (c *GaussDBforNoSQLClient) RestartInstanceInvoker(request *model.RestartInstanceRequest) *RestartInstanceInvoker {
	requestDef := GenReqDefForRestartInstance()
	return &RestartInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreExistingInstance 恢复到已有实例
//
// 恢复到已有实例
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// SaveLtsConfigs 关联LTS日志流
//
// - 将实例日志与LTS日志流关联，后台将自动上传实例日志到关联的LTS日志流里。
// - 关联成功后，会产生一定费用，具体计费可参考云日志服务（LTS）的定价详情。
// - 系统会为当前选择的日志流创建对应日志类型的结构化配置，若该日志流已存在其他日志类型的结构化配置，系统会进行覆盖。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) SaveLtsConfigs(request *model.SaveLtsConfigsRequest) (*model.SaveLtsConfigsResponse, error) {
	requestDef := GenReqDefForSaveLtsConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SaveLtsConfigsResponse), nil
	}
}

// SaveLtsConfigsInvoker 关联LTS日志流
func (c *GaussDBforNoSQLClient) SaveLtsConfigsInvoker(request *model.SaveLtsConfigsRequest) *SaveLtsConfigsInvoker {
	requestDef := GenReqDefForSaveLtsConfigs()
	return &SaveLtsConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetAutoEnlargePolicy 设置磁盘自动扩容策略
//
// 设置磁盘自动扩容策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) SetAutoEnlargePolicy(request *model.SetAutoEnlargePolicyRequest) (*model.SetAutoEnlargePolicyResponse, error) {
	requestDef := GenReqDefForSetAutoEnlargePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetAutoEnlargePolicyResponse), nil
	}
}

// SetAutoEnlargePolicyInvoker 设置磁盘自动扩容策略
func (c *GaussDBforNoSQLClient) SetAutoEnlargePolicyInvoker(request *model.SetAutoEnlargePolicyRequest) *SetAutoEnlargePolicyInvoker {
	requestDef := GenReqDefForSetAutoEnlargePolicy()
	return &SetAutoEnlargePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetBackupPolicy 设置自动备份策略
//
// 设置自动备份策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// SetRecyclePolicy 设置回收策略
//
// 设置已删除实例保留天数，修改保留天数后删除的实例按照新的天数保留，修改之前已在回收站的实例保留天数不变。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) SetRecyclePolicy(request *model.SetRecyclePolicyRequest) (*model.SetRecyclePolicyResponse, error) {
	requestDef := GenReqDefForSetRecyclePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetRecyclePolicyResponse), nil
	}
}

// SetRecyclePolicyInvoker 设置回收策略
func (c *GaussDBforNoSQLClient) SetRecyclePolicyInvoker(request *model.SetRecyclePolicyRequest) *SetRecyclePolicyInvoker {
	requestDef := GenReqDefForSetRecyclePolicy()
	return &SetRecyclePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAllInstancesBackups 查询备份列表
//
// 根据指定条件查询备份列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShowAllInstancesBackups(request *model.ShowAllInstancesBackupsRequest) (*model.ShowAllInstancesBackupsResponse, error) {
	requestDef := GenReqDefForShowAllInstancesBackups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAllInstancesBackupsResponse), nil
	}
}

// ShowAllInstancesBackupsInvoker 查询备份列表
func (c *GaussDBforNoSQLClient) ShowAllInstancesBackupsInvoker(request *model.ShowAllInstancesBackupsRequest) *ShowAllInstancesBackupsInvoker {
	requestDef := GenReqDefForShowAllInstancesBackups()
	return &ShowAllInstancesBackupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAllInstancesBackupsNew 查询备份列表（推荐）
//
// 根据指定条件查询备份列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShowAllInstancesBackupsNew(request *model.ShowAllInstancesBackupsNewRequest) (*model.ShowAllInstancesBackupsNewResponse, error) {
	requestDef := GenReqDefForShowAllInstancesBackupsNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAllInstancesBackupsNewResponse), nil
	}
}

// ShowAllInstancesBackupsNewInvoker 查询备份列表（推荐）
func (c *GaussDBforNoSQLClient) ShowAllInstancesBackupsNewInvoker(request *model.ShowAllInstancesBackupsNewRequest) *ShowAllInstancesBackupsNewInvoker {
	requestDef := GenReqDefForShowAllInstancesBackupsNew()
	return &ShowAllInstancesBackupsNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApplicableInstances 查询参数模板可应用的实例列表
//
// 查询参数模板可应用的实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShowApplicableInstances(request *model.ShowApplicableInstancesRequest) (*model.ShowApplicableInstancesResponse, error) {
	requestDef := GenReqDefForShowApplicableInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApplicableInstancesResponse), nil
	}
}

// ShowApplicableInstancesInvoker 查询参数模板可应用的实例列表
func (c *GaussDBforNoSQLClient) ShowApplicableInstancesInvoker(request *model.ShowApplicableInstancesRequest) *ShowApplicableInstancesInvoker {
	requestDef := GenReqDefForShowApplicableInstances()
	return &ShowApplicableInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApplyHistory 查询参数模板应用历史
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShowApplyHistory(request *model.ShowApplyHistoryRequest) (*model.ShowApplyHistoryResponse, error) {
	requestDef := GenReqDefForShowApplyHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApplyHistoryResponse), nil
	}
}

// ShowApplyHistoryInvoker 查询参数模板应用历史
func (c *GaussDBforNoSQLClient) ShowApplyHistoryInvoker(request *model.ShowApplyHistoryRequest) *ShowApplyHistoryInvoker {
	requestDef := GenReqDefForShowApplyHistory()
	return &ShowApplyHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAutoEnlargePolicy 查询磁盘自动扩容策略
//
// 查询磁盘自动扩容策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShowAutoEnlargePolicy(request *model.ShowAutoEnlargePolicyRequest) (*model.ShowAutoEnlargePolicyResponse, error) {
	requestDef := GenReqDefForShowAutoEnlargePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAutoEnlargePolicyResponse), nil
	}
}

// ShowAutoEnlargePolicyInvoker 查询磁盘自动扩容策略
func (c *GaussDBforNoSQLClient) ShowAutoEnlargePolicyInvoker(request *model.ShowAutoEnlargePolicyRequest) *ShowAutoEnlargePolicyInvoker {
	requestDef := GenReqDefForShowAutoEnlargePolicy()
	return &ShowAutoEnlargePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBackupPolicy 查询自动备份策略
//
// 查询自动备份策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowElbIpGroup 查询实例负载均衡的IP访问黑白名单
//
// 查询实例负载均衡的IP访问黑白名单。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShowElbIpGroup(request *model.ShowElbIpGroupRequest) (*model.ShowElbIpGroupResponse, error) {
	requestDef := GenReqDefForShowElbIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowElbIpGroupResponse), nil
	}
}

// ShowElbIpGroupInvoker 查询实例负载均衡的IP访问黑白名单
func (c *GaussDBforNoSQLClient) ShowElbIpGroupInvoker(request *model.ShowElbIpGroupRequest) *ShowElbIpGroupInvoker {
	requestDef := GenReqDefForShowElbIpGroup()
	return &ShowElbIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowErrorLog 查询数据库错误日志信息
//
// 查询数据库错误日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShowErrorLog(request *model.ShowErrorLogRequest) (*model.ShowErrorLogResponse, error) {
	requestDef := GenReqDefForShowErrorLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowErrorLogResponse), nil
	}
}

// ShowErrorLogInvoker 查询数据库错误日志信息
func (c *GaussDBforNoSQLClient) ShowErrorLogInvoker(request *model.ShowErrorLogRequest) *ShowErrorLogInvoker {
	requestDef := GenReqDefForShowErrorLog()
	return &ShowErrorLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHighRiskCommands 查询高危命令
//
// 查询Redis的高危命令
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShowHighRiskCommands(request *model.ShowHighRiskCommandsRequest) (*model.ShowHighRiskCommandsResponse, error) {
	requestDef := GenReqDefForShowHighRiskCommands()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHighRiskCommandsResponse), nil
	}
}

// ShowHighRiskCommandsInvoker 查询高危命令
func (c *GaussDBforNoSQLClient) ShowHighRiskCommandsInvoker(request *model.ShowHighRiskCommandsRequest) *ShowHighRiskCommandsInvoker {
	requestDef := GenReqDefForShowHighRiskCommands()
	return &ShowHighRiskCommandsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceBiactiveRegions 查询实例可搭建双活关系的Region
//
// 查询实例可搭建双活关系的Region。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShowInstanceBiactiveRegions(request *model.ShowInstanceBiactiveRegionsRequest) (*model.ShowInstanceBiactiveRegionsResponse, error) {
	requestDef := GenReqDefForShowInstanceBiactiveRegions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceBiactiveRegionsResponse), nil
	}
}

// ShowInstanceBiactiveRegionsInvoker 查询实例可搭建双活关系的Region
func (c *GaussDBforNoSQLClient) ShowInstanceBiactiveRegionsInvoker(request *model.ShowInstanceBiactiveRegionsRequest) *ShowInstanceBiactiveRegionsInvoker {
	requestDef := GenReqDefForShowInstanceBiactiveRegions()
	return &ShowInstanceBiactiveRegionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceConfiguration 获取指定实例的参数
//
// 获取指定实例的参数信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowIpNumRequirement 查询创建实例或扩容节点时需要的IP数量
//
// 查询创建实例或扩容节点时需要的IP数量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShowIpNumRequirement(request *model.ShowIpNumRequirementRequest) (*model.ShowIpNumRequirementResponse, error) {
	requestDef := GenReqDefForShowIpNumRequirement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIpNumRequirementResponse), nil
	}
}

// ShowIpNumRequirementInvoker 查询创建实例或扩容节点时需要的IP数量
func (c *GaussDBforNoSQLClient) ShowIpNumRequirementInvoker(request *model.ShowIpNumRequirementRequest) *ShowIpNumRequirementInvoker {
	requestDef := GenReqDefForShowIpNumRequirement()
	return &ShowIpNumRequirementInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowModifyHistory 查询实例参数的修改历史
//
// 查询实例参数的修改历史
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShowModifyHistory(request *model.ShowModifyHistoryRequest) (*model.ShowModifyHistoryResponse, error) {
	requestDef := GenReqDefForShowModifyHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowModifyHistoryResponse), nil
	}
}

// ShowModifyHistoryInvoker 查询实例参数的修改历史
func (c *GaussDBforNoSQLClient) ShowModifyHistoryInvoker(request *model.ShowModifyHistoryRequest) *ShowModifyHistoryInvoker {
	requestDef := GenReqDefForShowModifyHistory()
	return &ShowModifyHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPasswordlessConfig 获取GeminiDB Redis的免密配置
//
// 获取GeminiDB Redis的免密配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShowPasswordlessConfig(request *model.ShowPasswordlessConfigRequest) (*model.ShowPasswordlessConfigResponse, error) {
	requestDef := GenReqDefForShowPasswordlessConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPasswordlessConfigResponse), nil
	}
}

// ShowPasswordlessConfigInvoker 获取GeminiDB Redis的免密配置
func (c *GaussDBforNoSQLClient) ShowPasswordlessConfigInvoker(request *model.ShowPasswordlessConfigRequest) *ShowPasswordlessConfigInvoker {
	requestDef := GenReqDefForShowPasswordlessConfig()
	return &ShowPasswordlessConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPauseResumeStutus 获取容灾实例数据同步状态
//
// 获取容灾实例数据同步状态，主备实例id，数据同步指标值，以及倒换和切换场景下的RPO，RTO指标值。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShowPauseResumeStutus(request *model.ShowPauseResumeStutusRequest) (*model.ShowPauseResumeStutusResponse, error) {
	requestDef := GenReqDefForShowPauseResumeStutus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPauseResumeStutusResponse), nil
	}
}

// ShowPauseResumeStutusInvoker 获取容灾实例数据同步状态
func (c *GaussDBforNoSQLClient) ShowPauseResumeStutusInvoker(request *model.ShowPauseResumeStutusRequest) *ShowPauseResumeStutusInvoker {
	requestDef := GenReqDefForShowPauseResumeStutus()
	return &ShowPauseResumeStutusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuotas 查询配额
//
// 查询单租户在GeminiDB服务下的资源配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowRecyclePolicy 查询回收策略
//
// 查询回收策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShowRecyclePolicy(request *model.ShowRecyclePolicyRequest) (*model.ShowRecyclePolicyResponse, error) {
	requestDef := GenReqDefForShowRecyclePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecyclePolicyResponse), nil
	}
}

// ShowRecyclePolicyInvoker 查询回收策略
func (c *GaussDBforNoSQLClient) ShowRecyclePolicyInvoker(request *model.ShowRecyclePolicyRequest) *ShowRecyclePolicyInvoker {
	requestDef := GenReqDefForShowRecyclePolicy()
	return &ShowRecyclePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRedisBigKeys 查询Redis实例的大key
//
// 支持查询Redis实例的大key。value长度大于bigkeys-string-threshold参数的string类型的key或者元素数大于bigkeys-composite-threshold参数的hash/list/zset/set/stream类型key，会被判断为大key。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShowRedisBigKeys(request *model.ShowRedisBigKeysRequest) (*model.ShowRedisBigKeysResponse, error) {
	requestDef := GenReqDefForShowRedisBigKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRedisBigKeysResponse), nil
	}
}

// ShowRedisBigKeysInvoker 查询Redis实例的大key
func (c *GaussDBforNoSQLClient) ShowRedisBigKeysInvoker(request *model.ShowRedisBigKeysRequest) *ShowRedisBigKeysInvoker {
	requestDef := GenReqDefForShowRedisBigKeys()
	return &ShowRedisBigKeysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRestorableList 查询可恢复的实例列表
//
// 查询用户可恢复的实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowSlowLogDesensitization 查询慢日志脱敏状态
//
// 查询慢日志脱敏状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) ShowSlowLogDesensitization(request *model.ShowSlowLogDesensitizationRequest) (*model.ShowSlowLogDesensitizationResponse, error) {
	requestDef := GenReqDefForShowSlowLogDesensitization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSlowLogDesensitizationResponse), nil
	}
}

// ShowSlowLogDesensitizationInvoker 查询慢日志脱敏状态
func (c *GaussDBforNoSQLClient) ShowSlowLogDesensitizationInvoker(request *model.ShowSlowLogDesensitizationRequest) *ShowSlowLogDesensitizationInvoker {
	requestDef := GenReqDefForShowSlowLogDesensitization()
	return &ShowSlowLogDesensitizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShrinkInstanceNode 缩容指定集群实例的节点数量
//
// 缩容指定集群实例的节点数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// SwitchIpGroup 设置实例负载均衡的IP访问黑白名单
//
// 设置实例负载均衡的IP访问黑白名单，黑名单、白名单只能选一种，每次调用此接口覆盖之前的设置。关闭后不限制连接的源IP地址。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) SwitchIpGroup(request *model.SwitchIpGroupRequest) (*model.SwitchIpGroupResponse, error) {
	requestDef := GenReqDefForSwitchIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchIpGroupResponse), nil
	}
}

// SwitchIpGroupInvoker 设置实例负载均衡的IP访问黑白名单
func (c *GaussDBforNoSQLClient) SwitchIpGroupInvoker(request *model.SwitchIpGroupRequest) *SwitchIpGroupInvoker {
	requestDef := GenReqDefForSwitchIpGroup()
	return &SwitchIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchSlowlogDesensitization 设置慢日志脱敏状态
//
// 设置慢日志脱敏状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) SwitchSlowlogDesensitization(request *model.SwitchSlowlogDesensitizationRequest) (*model.SwitchSlowlogDesensitizationResponse, error) {
	requestDef := GenReqDefForSwitchSlowlogDesensitization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchSlowlogDesensitizationResponse), nil
	}
}

// SwitchSlowlogDesensitizationInvoker 设置慢日志脱敏状态
func (c *GaussDBforNoSQLClient) SwitchSlowlogDesensitizationInvoker(request *model.SwitchSlowlogDesensitizationRequest) *SwitchSlowlogDesensitizationInvoker {
	requestDef := GenReqDefForSwitchSlowlogDesensitization()
	return &SwitchSlowlogDesensitizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchSsl 切换实例SSL开关
//
// 切换实例SSL开关。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) SwitchSsl(request *model.SwitchSslRequest) (*model.SwitchSslResponse, error) {
	requestDef := GenReqDefForSwitchSsl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchSslResponse), nil
	}
}

// SwitchSslInvoker 切换实例SSL开关
func (c *GaussDBforNoSQLClient) SwitchSslInvoker(request *model.SwitchSslRequest) *SwitchSslInvoker {
	requestDef := GenReqDefForSwitchSsl()
	return &SwitchSslInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchToMaster 容灾实例备升主
//
// 该接口用于对已经搭建容灾关系的实例，将备实例升级为主实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// UpdateClientNetwork 修改副本集跨网段访问配置
//
// 修改副本集跨网段访问配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) UpdateClientNetwork(request *model.UpdateClientNetworkRequest) (*model.UpdateClientNetworkResponse, error) {
	requestDef := GenReqDefForUpdateClientNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClientNetworkResponse), nil
	}
}

// UpdateClientNetworkInvoker 修改副本集跨网段访问配置
func (c *GaussDBforNoSQLClient) UpdateClientNetworkInvoker(request *model.UpdateClientNetworkRequest) *UpdateClientNetworkInvoker {
	requestDef := GenReqDefForUpdateClientNetwork()
	return &UpdateClientNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateConfiguration 修改参数模板参数
//
// 修改参数模板参数。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// UpdateDatabases 操作GeminDB实例数据库
//
// 操作GeminDB实例数据库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) UpdateDatabases(request *model.UpdateDatabasesRequest) (*model.UpdateDatabasesResponse, error) {
	requestDef := GenReqDefForUpdateDatabases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDatabasesResponse), nil
	}
}

// UpdateDatabasesInvoker 操作GeminDB实例数据库
func (c *GaussDBforNoSQLClient) UpdateDatabasesInvoker(request *model.UpdateDatabasesRequest) *UpdateDatabasesInvoker {
	requestDef := GenReqDefForUpdateDatabases()
	return &UpdateDatabasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHighRiskCommands 修改高危命令
//
// 批量修改高危命令
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) UpdateHighRiskCommands(request *model.UpdateHighRiskCommandsRequest) (*model.UpdateHighRiskCommandsResponse, error) {
	requestDef := GenReqDefForUpdateHighRiskCommands()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHighRiskCommandsResponse), nil
	}
}

// UpdateHighRiskCommandsInvoker 修改高危命令
func (c *GaussDBforNoSQLClient) UpdateHighRiskCommandsInvoker(request *model.UpdateHighRiskCommandsRequest) *UpdateHighRiskCommandsInvoker {
	requestDef := GenReqDefForUpdateHighRiskCommands()
	return &UpdateHighRiskCommandsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceConfiguration 修改指定实例的参数
//
// 修改指定实例的参数。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// UpdatePasswordlessConfig 支持修改GeminiDB Redis的免密配置
//
// 支持修改GeminiDB Redis的免密配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforNoSQLClient) UpdatePasswordlessConfig(request *model.UpdatePasswordlessConfigRequest) (*model.UpdatePasswordlessConfigResponse, error) {
	requestDef := GenReqDefForUpdatePasswordlessConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePasswordlessConfigResponse), nil
	}
}

// UpdatePasswordlessConfigInvoker 支持修改GeminiDB Redis的免密配置
func (c *GaussDBforNoSQLClient) UpdatePasswordlessConfigInvoker(request *model.UpdatePasswordlessConfigRequest) *UpdatePasswordlessConfigInvoker {
	requestDef := GenReqDefForUpdatePasswordlessConfig()
	return &UpdatePasswordlessConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSecurityGroup 变更实例安全组
//
// 变更实例关联的安全组
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
