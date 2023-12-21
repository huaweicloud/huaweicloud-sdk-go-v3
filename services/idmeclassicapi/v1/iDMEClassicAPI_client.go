package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/idmeclassicapi/v1/model"
)

type IDMEClassicAPIClient struct {
	HcClient *http_client.HcHttpClient
}

func NewIDMEClassicAPIClient(hcClient *http_client.HcHttpClient) *IDMEClassicAPIClient {
	return &IDMEClassicAPIClient{HcClient: hcClient}
}

func IDMEClassicAPIClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// BatchCreateUsingPost 批量创建实例
//
// 批量创建指定数据模型的数据实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchCreateUsingPost(request *model.BatchCreateUsingPostRequest) (*model.BatchCreateUsingPostResponse, error) {
	requestDef := GenReqDefForBatchCreateUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateUsingPostResponse), nil
	}
}

// BatchCreateUsingPostInvoker 批量创建实例
func (c *IDMEClassicAPIClient) BatchCreateUsingPostInvoker(request *model.BatchCreateUsingPostRequest) *BatchCreateUsingPostInvoker {
	requestDef := GenReqDefForBatchCreateUsingPost()
	return &BatchCreateUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteLogicalUsingPost 批量软删除实例
//
// 根据数据实例的唯一编码，批量软删除指定数据模型中的多个数据实例。
//
// 通过此接口进行删除操作时，系统会将当前删除的实例转存至XDM应用的XDMLogicDeleteData内置模型中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchDeleteLogicalUsingPost(request *model.BatchDeleteLogicalUsingPostRequest) (*model.BatchDeleteLogicalUsingPostResponse, error) {
	requestDef := GenReqDefForBatchDeleteLogicalUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteLogicalUsingPostResponse), nil
	}
}

// BatchDeleteLogicalUsingPostInvoker 批量软删除实例
func (c *IDMEClassicAPIClient) BatchDeleteLogicalUsingPostInvoker(request *model.BatchDeleteLogicalUsingPostRequest) *BatchDeleteLogicalUsingPostInvoker {
	requestDef := GenReqDefForBatchDeleteLogicalUsingPost()
	return &BatchDeleteLogicalUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteUsingPost 批量删除实例
//
// 根据数据实例的唯一编码，批量删除指定数据模型中的多个数据实例。
//
// 请您谨慎使用删除操作，实例删除后将无法恢复。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchDeleteUsingPost(request *model.BatchDeleteUsingPostRequest) (*model.BatchDeleteUsingPostResponse, error) {
	requestDef := GenReqDefForBatchDeleteUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteUsingPostResponse), nil
	}
}

// BatchDeleteUsingPostInvoker 批量删除实例
func (c *IDMEClassicAPIClient) BatchDeleteUsingPostInvoker(request *model.BatchDeleteUsingPostRequest) *BatchDeleteUsingPostInvoker {
	requestDef := GenReqDefForBatchDeleteUsingPost()
	return &BatchDeleteUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchShowGetUsingPost 批量查询实例
//
// 根据多个数据实例的唯一编码，批量查询实例的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchShowGetUsingPost(request *model.BatchShowGetUsingPostRequest) (*model.BatchShowGetUsingPostResponse, error) {
	requestDef := GenReqDefForBatchShowGetUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowGetUsingPostResponse), nil
	}
}

// BatchShowGetUsingPostInvoker 批量查询实例
func (c *IDMEClassicAPIClient) BatchShowGetUsingPostInvoker(request *model.BatchShowGetUsingPostRequest) *BatchShowGetUsingPostInvoker {
	requestDef := GenReqDefForBatchShowGetUsingPost()
	return &BatchShowGetUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateUsingPost 批量更新实例
//
// 批量更新指定数据模型中的多个实例数据。如果某个实例的唯一编码不存在，该实例不做任何更新操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchUpdateUsingPost(request *model.BatchUpdateUsingPostRequest) (*model.BatchUpdateUsingPostResponse, error) {
	requestDef := GenReqDefForBatchUpdateUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateUsingPostResponse), nil
	}
}

// BatchUpdateUsingPostInvoker 批量更新实例
func (c *IDMEClassicAPIClient) BatchUpdateUsingPostInvoker(request *model.BatchUpdateUsingPostRequest) *BatchUpdateUsingPostInvoker {
	requestDef := GenReqDefForBatchUpdateUsingPost()
	return &BatchUpdateUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountUsingPost 统计指定数据模型的实例总数
//
// 根据指定的查询条件，统计指定数据模型中的实例总数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) CountUsingPost(request *model.CountUsingPostRequest) (*model.CountUsingPostResponse, error) {
	requestDef := GenReqDefForCountUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountUsingPostResponse), nil
	}
}

// CountUsingPostInvoker 统计指定数据模型的实例总数
func (c *IDMEClassicAPIClient) CountUsingPostInvoker(request *model.CountUsingPostRequest) *CountUsingPostInvoker {
	requestDef := GenReqDefForCountUsingPost()
	return &CountUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateUsingPost 创建实例
//
// 创建指定数据模型的数据实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) CreateUsingPost(request *model.CreateUsingPostRequest) (*model.CreateUsingPostResponse, error) {
	requestDef := GenReqDefForCreateUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateUsingPostResponse), nil
	}
}

// CreateUsingPostInvoker 创建实例
func (c *IDMEClassicAPIClient) CreateUsingPostInvoker(request *model.CreateUsingPostRequest) *CreateUsingPostInvoker {
	requestDef := GenReqDefForCreateUsingPost()
	return &CreateUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteByConditionUsingPost 根据指定条件删除实例
//
// 通过此接口，删除指定条件查询返回的实例。
//
// 请您谨慎使用删除操作，实例删除后将无法恢复。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) DeleteByConditionUsingPost(request *model.DeleteByConditionUsingPostRequest) (*model.DeleteByConditionUsingPostResponse, error) {
	requestDef := GenReqDefForDeleteByConditionUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteByConditionUsingPostResponse), nil
	}
}

// DeleteByConditionUsingPostInvoker 根据指定条件删除实例
func (c *IDMEClassicAPIClient) DeleteByConditionUsingPostInvoker(request *model.DeleteByConditionUsingPostRequest) *DeleteByConditionUsingPostInvoker {
	requestDef := GenReqDefForDeleteByConditionUsingPost()
	return &DeleteByConditionUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteUsingPost 删除实例
//
// 根据数据实例的唯一编码，删除指定数据模型中的一个数据实例。
//
// 请您谨慎使用删除操作，实例删除后将无法恢复。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) DeleteUsingPost(request *model.DeleteUsingPostRequest) (*model.DeleteUsingPostResponse, error) {
	requestDef := GenReqDefForDeleteUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteUsingPostResponse), nil
	}
}

// DeleteUsingPostInvoker 删除实例
func (c *IDMEClassicAPIClient) DeleteUsingPostInvoker(request *model.DeleteUsingPostRequest) *DeleteUsingPostInvoker {
	requestDef := GenReqDefForDeleteUsingPost()
	return &DeleteUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQueryUsingPost 根据“列表属性”为“是”的属性查询实例
//
// 当数据模型中存在“列表属性”为“是”的属性时，可通过此接口查询数据模型中的实例数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ListQueryUsingPost(request *model.ListQueryUsingPostRequest) (*model.ListQueryUsingPostResponse, error) {
	requestDef := GenReqDefForListQueryUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueryUsingPostResponse), nil
	}
}

// ListQueryUsingPostInvoker 根据“列表属性”为“是”的属性查询实例
func (c *IDMEClassicAPIClient) ListQueryUsingPostInvoker(request *model.ListQueryUsingPostRequest) *ListQueryUsingPostInvoker {
	requestDef := GenReqDefForListQueryUsingPost()
	return &ListQueryUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSelectUsingPost 查询实例的指定属性
//
// 根据查询条件及指定属性分页返回（不支持扩展属性作为选定属性列)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ListSelectUsingPost(request *model.ListSelectUsingPostRequest) (*model.ListSelectUsingPostResponse, error) {
	requestDef := GenReqDefForListSelectUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSelectUsingPostResponse), nil
	}
}

// ListSelectUsingPostInvoker 查询实例的指定属性
func (c *IDMEClassicAPIClient) ListSelectUsingPostInvoker(request *model.ListSelectUsingPostRequest) *ListSelectUsingPostInvoker {
	requestDef := GenReqDefForListSelectUsingPost()
	return &ListSelectUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUsingPost 查询实例的基础属性
//
// 根据查询条件分页返回模型基本属性信息且不级联查询（不支持扩展属性作为查询条件）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ListUsingPost(request *model.ListUsingPostRequest) (*model.ListUsingPostResponse, error) {
	requestDef := GenReqDefForListUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsingPostResponse), nil
	}
}

// ListUsingPostInvoker 查询实例的基础属性
func (c *IDMEClassicAPIClient) ListUsingPostInvoker(request *model.ListUsingPostRequest) *ListUsingPostInvoker {
	requestDef := GenReqDefForListUsingPost()
	return &ListUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SaveAllUsingPost 根据唯一键为“是”的属性更新实例数据
//
// 当数据模型中存在“唯一键”为“是”的属性时，可根据该属性的英文名称更新该数据模型中实例的所有字段数据。如果更新的实例不存在，系统将自动创建该实例数据。
//
// 调用此接口时，建议传入该实例的所有字段信息。如果未传入某个字段，该字段的数据将更新为空值。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) SaveAllUsingPost(request *model.SaveAllUsingPostRequest) (*model.SaveAllUsingPostResponse, error) {
	requestDef := GenReqDefForSaveAllUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SaveAllUsingPostResponse), nil
	}
}

// SaveAllUsingPostInvoker 根据唯一键为“是”的属性更新实例数据
func (c *IDMEClassicAPIClient) SaveAllUsingPostInvoker(request *model.SaveAllUsingPostRequest) *SaveAllUsingPostInvoker {
	requestDef := GenReqDefForSaveAllUsingPost()
	return &SaveAllUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SaveAsUsingPost 另存版本对象的实例数据
//
// 版本对象的另存为接口（saveAs）用于创建一条与原版本对象实例数据相同的数据实例。该实例数据会完全复制原实例现有的数据，包括与其关联的主对象和分支对象，且新实例数据的版本号从初始值开始计算。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) SaveAsUsingPost(request *model.SaveAsUsingPostRequest) (*model.SaveAsUsingPostResponse, error) {
	requestDef := GenReqDefForSaveAsUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SaveAsUsingPostResponse), nil
	}
}

// SaveAsUsingPostInvoker 另存版本对象的实例数据
func (c *IDMEClassicAPIClient) SaveAsUsingPostInvoker(request *model.SaveAsUsingPostRequest) *SaveAsUsingPostInvoker {
	requestDef := GenReqDefForSaveAsUsingPost()
	return &SaveAsUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SaveUsingPost 根据唯一键为“是”的属性更新实例的指定字段
//
// 当数据模型中存在“唯一键”为“是”的属性时，可根据该属性的英文名称更新该数据模型中实例的指定字段数据。
//
// 如果更新的实例不存在，系统将自动创建该实例数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) SaveUsingPost(request *model.SaveUsingPostRequest) (*model.SaveUsingPostResponse, error) {
	requestDef := GenReqDefForSaveUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SaveUsingPostResponse), nil
	}
}

// SaveUsingPostInvoker 根据唯一键为“是”的属性更新实例的指定字段
func (c *IDMEClassicAPIClient) SaveUsingPostInvoker(request *model.SaveUsingPostRequest) *SaveUsingPostInvoker {
	requestDef := GenReqDefForSaveUsingPost()
	return &SaveUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFindUsingPost 分页查询实例
//
// 分页查询指定数据模型中的所有实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowFindUsingPost(request *model.ShowFindUsingPostRequest) (*model.ShowFindUsingPostResponse, error) {
	requestDef := GenReqDefForShowFindUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFindUsingPostResponse), nil
	}
}

// ShowFindUsingPostInvoker 分页查询实例
func (c *IDMEClassicAPIClient) ShowFindUsingPostInvoker(request *model.ShowFindUsingPostRequest) *ShowFindUsingPostInvoker {
	requestDef := GenReqDefForShowFindUsingPost()
	return &ShowFindUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGetByUniqueKey 根据唯一键为“是”的属性查询实例
//
// 当数据模型中存在“唯一键”为“是”的属性时，可根据该属性查询实例数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowGetByUniqueKey(request *model.ShowGetByUniqueKeyRequest) (*model.ShowGetByUniqueKeyResponse, error) {
	requestDef := GenReqDefForShowGetByUniqueKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGetByUniqueKeyResponse), nil
	}
}

// ShowGetByUniqueKeyInvoker 根据唯一键为“是”的属性查询实例
func (c *IDMEClassicAPIClient) ShowGetByUniqueKeyInvoker(request *model.ShowGetByUniqueKeyRequest) *ShowGetByUniqueKeyInvoker {
	requestDef := GenReqDefForShowGetByUniqueKey()
	return &ShowGetByUniqueKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGetUsingPost 查询实例
//
// 根据一个数据实例的唯一编码，查询实例的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowGetUsingPost(request *model.ShowGetUsingPostRequest) (*model.ShowGetUsingPostResponse, error) {
	requestDef := GenReqDefForShowGetUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGetUsingPostResponse), nil
	}
}

// ShowGetUsingPostInvoker 查询实例
func (c *IDMEClassicAPIClient) ShowGetUsingPostInvoker(request *model.ShowGetUsingPostRequest) *ShowGetUsingPostInvoker {
	requestDef := GenReqDefForShowGetUsingPost()
	return &ShowGetUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLogicalDeleteByConditionUsingPost 根据指定条件软删除实例
//
// 通过此接口，软删除指定条件查询返回的实例。
//
// 通过此接口进行删除操作时，系统会将当前删除的实例转存至XDM应用的XDMLogicDeleteData内置模型中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowLogicalDeleteByConditionUsingPost(request *model.ShowLogicalDeleteByConditionUsingPostRequest) (*model.ShowLogicalDeleteByConditionUsingPostResponse, error) {
	requestDef := GenReqDefForShowLogicalDeleteByConditionUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLogicalDeleteByConditionUsingPostResponse), nil
	}
}

// ShowLogicalDeleteByConditionUsingPostInvoker 根据指定条件软删除实例
func (c *IDMEClassicAPIClient) ShowLogicalDeleteByConditionUsingPostInvoker(request *model.ShowLogicalDeleteByConditionUsingPostRequest) *ShowLogicalDeleteByConditionUsingPostInvoker {
	requestDef := GenReqDefForShowLogicalDeleteByConditionUsingPost()
	return &ShowLogicalDeleteByConditionUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLogicalDeleteUsingPost 软删除实例
//
// 根据数据实例的唯一编码，软删除指定数据模型中的一个数据实例。
//
// 通过此接口进行删除操作时，系统会将当前删除的实例转存至XDM应用的XDMLogicDeleteData内置模型中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowLogicalDeleteUsingPost(request *model.ShowLogicalDeleteUsingPostRequest) (*model.ShowLogicalDeleteUsingPostResponse, error) {
	requestDef := GenReqDefForShowLogicalDeleteUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLogicalDeleteUsingPostResponse), nil
	}
}

// ShowLogicalDeleteUsingPostInvoker 软删除实例
func (c *IDMEClassicAPIClient) ShowLogicalDeleteUsingPostInvoker(request *model.ShowLogicalDeleteUsingPostRequest) *ShowLogicalDeleteUsingPostInvoker {
	requestDef := GenReqDefForShowLogicalDeleteUsingPost()
	return &ShowLogicalDeleteUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStaticsUsingPost 查询指定数据模型的实例统计信息
//
// 根据指定函数，统计指定数据模型的实例信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowStaticsUsingPost(request *model.ShowStaticsUsingPostRequest) (*model.ShowStaticsUsingPostResponse, error) {
	requestDef := GenReqDefForShowStaticsUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStaticsUsingPostResponse), nil
	}
}

// ShowStaticsUsingPostInvoker 查询指定数据模型的实例统计信息
func (c *IDMEClassicAPIClient) ShowStaticsUsingPostInvoker(request *model.ShowStaticsUsingPostRequest) *ShowStaticsUsingPostInvoker {
	requestDef := GenReqDefForShowStaticsUsingPost()
	return &ShowStaticsUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateByConditionUsingPost 根据指定条件更新实例
//
// 根据指定条件更新指定模型的实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) UpdateByConditionUsingPost(request *model.UpdateByConditionUsingPostRequest) (*model.UpdateByConditionUsingPostResponse, error) {
	requestDef := GenReqDefForUpdateByConditionUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateByConditionUsingPostResponse), nil
	}
}

// UpdateByConditionUsingPostInvoker 根据指定条件更新实例
func (c *IDMEClassicAPIClient) UpdateByConditionUsingPostInvoker(request *model.UpdateByConditionUsingPostRequest) *UpdateByConditionUsingPostInvoker {
	requestDef := GenReqDefForUpdateByConditionUsingPost()
	return &UpdateByConditionUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUsingPost 更新实例
//
// 更新指定数据模型中的一个实例数据。如果实例的唯一编码不存在，则不做任何更新操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) UpdateUsingPost(request *model.UpdateUsingPostRequest) (*model.UpdateUsingPostResponse, error) {
	requestDef := GenReqDefForUpdateUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUsingPostResponse), nil
	}
}

// UpdateUsingPostInvoker 更新实例
func (c *IDMEClassicAPIClient) UpdateUsingPostInvoker(request *model.UpdateUsingPostRequest) *UpdateUsingPostInvoker {
	requestDef := GenReqDefForUpdateUsingPost()
	return &UpdateUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
