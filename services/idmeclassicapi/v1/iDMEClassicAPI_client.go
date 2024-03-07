package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/idmeclassicapi/v1/model"
)

type IDMEClassicAPIClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewIDMEClassicAPIClient(hcClient *httpclient.HcHttpClient) *IDMEClassicAPIClient {
	return &IDMEClassicAPIClient{HcClient: hcClient}
}

func IDMEClassicAPIClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// BatchCheckin XDM_批量检入VersionModel
//
// 根据主对象ID批量检入版本对象，小版本升版。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchCheckin(request *model.BatchCheckinRequest) (*model.BatchCheckinResponse, error) {
	requestDef := GenReqDefForBatchCheckin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCheckinResponse), nil
	}
}

// BatchCheckinInvoker XDM_批量检入VersionModel
func (c *IDMEClassicAPIClient) BatchCheckinInvoker(request *model.BatchCheckinRequest) *BatchCheckinInvoker {
	requestDef := GenReqDefForBatchCheckin()
	return &BatchCheckinInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCheckout XDM_批量检出VersionModel
//
// 根据主对象ID批量检出版本对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchCheckout(request *model.BatchCheckoutRequest) (*model.BatchCheckoutResponse, error) {
	requestDef := GenReqDefForBatchCheckout()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCheckoutResponse), nil
	}
}

// BatchCheckoutInvoker XDM_批量检出VersionModel
func (c *IDMEClassicAPIClient) BatchCheckoutInvoker(request *model.BatchCheckoutRequest) *BatchCheckoutInvoker {
	requestDef := GenReqDefForBatchCheckout()
	return &BatchCheckoutInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCheckoutAndUpdate XDM_批量检出并更新VersionModel
//
// 根据主对象ID批量检出对象并根据传入字段批量更新版本对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchCheckoutAndUpdate(request *model.BatchCheckoutAndUpdateRequest) (*model.BatchCheckoutAndUpdateResponse, error) {
	requestDef := GenReqDefForBatchCheckoutAndUpdate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCheckoutAndUpdateResponse), nil
	}
}

// BatchCheckoutAndUpdateInvoker XDM_批量检出并更新VersionModel
func (c *IDMEClassicAPIClient) BatchCheckoutAndUpdateInvoker(request *model.BatchCheckoutAndUpdateRequest) *BatchCheckoutAndUpdateInvoker {
	requestDef := GenReqDefForBatchCheckoutAndUpdate()
	return &BatchCheckoutAndUpdateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCheckoutUndo XDM_批量撤销检出VersionModel
//
// 根据主对象ID批量撤销检出版本对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchCheckoutUndo(request *model.BatchCheckoutUndoRequest) (*model.BatchCheckoutUndoResponse, error) {
	requestDef := GenReqDefForBatchCheckoutUndo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCheckoutUndoResponse), nil
	}
}

// BatchCheckoutUndoInvoker XDM_批量撤销检出VersionModel
func (c *IDMEClassicAPIClient) BatchCheckoutUndoInvoker(request *model.BatchCheckoutUndoRequest) *BatchCheckoutUndoInvoker {
	requestDef := GenReqDefForBatchCheckoutUndo()
	return &BatchCheckoutUndoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCheckoutUndoByAdmin XDM_管理员批量撤销检出VersionModel
//
// 管理员根据主对象ID批量撤销检出版本对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchCheckoutUndoByAdmin(request *model.BatchCheckoutUndoByAdminRequest) (*model.BatchCheckoutUndoByAdminResponse, error) {
	requestDef := GenReqDefForBatchCheckoutUndoByAdmin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCheckoutUndoByAdminResponse), nil
	}
}

// BatchCheckoutUndoByAdminInvoker XDM_管理员批量撤销检出VersionModel
func (c *IDMEClassicAPIClient) BatchCheckoutUndoByAdminInvoker(request *model.BatchCheckoutUndoByAdminRequest) *BatchCheckoutUndoByAdminInvoker {
	requestDef := GenReqDefForBatchCheckoutUndoByAdmin()
	return &BatchCheckoutUndoByAdminInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// BatchDeleteBranch XDM_批量删除VersionModel最新分支版本下所有小版本
//
// 根据主对象ID&amp;业务版本列表，批量删除最新分支版本下的所有小版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchDeleteBranch(request *model.BatchDeleteBranchRequest) (*model.BatchDeleteBranchResponse, error) {
	requestDef := GenReqDefForBatchDeleteBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteBranchResponse), nil
	}
}

// BatchDeleteBranchInvoker XDM_批量删除VersionModel最新分支版本下所有小版本
func (c *IDMEClassicAPIClient) BatchDeleteBranchInvoker(request *model.BatchDeleteBranchRequest) *BatchDeleteBranchInvoker {
	requestDef := GenReqDefForBatchDeleteBranch()
	return &BatchDeleteBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteLogicalBranch XDM_批量软删除VersionModel最新分支版本下所有小版本
//
// 批量软删除最新分支版本下的所有小版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchDeleteLogicalBranch(request *model.BatchDeleteLogicalBranchRequest) (*model.BatchDeleteLogicalBranchResponse, error) {
	requestDef := GenReqDefForBatchDeleteLogicalBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteLogicalBranchResponse), nil
	}
}

// BatchDeleteLogicalBranchInvoker XDM_批量软删除VersionModel最新分支版本下所有小版本
func (c *IDMEClassicAPIClient) BatchDeleteLogicalBranchInvoker(request *model.BatchDeleteLogicalBranchRequest) *BatchDeleteLogicalBranchInvoker {
	requestDef := GenReqDefForBatchDeleteLogicalBranch()
	return &BatchDeleteLogicalBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// BatchExecuteRevise XDM_批量修订VersionModel。
//
// 根据主对象ID批量修订对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchExecuteRevise(request *model.BatchExecuteReviseRequest) (*model.BatchExecuteReviseResponse, error) {
	requestDef := GenReqDefForBatchExecuteRevise()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchExecuteReviseResponse), nil
	}
}

// BatchExecuteReviseInvoker XDM_批量修订VersionModel。
func (c *IDMEClassicAPIClient) BatchExecuteReviseInvoker(request *model.BatchExecuteReviseRequest) *BatchExecuteReviseInvoker {
	requestDef := GenReqDefForBatchExecuteRevise()
	return &BatchExecuteReviseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// BatchUpdateAndCheckin XDM_批量更新并检入VersionModel
//
// 根据传入字段批量更新版本对象并根据主对象ID批量检入对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchUpdateAndCheckin(request *model.BatchUpdateAndCheckinRequest) (*model.BatchUpdateAndCheckinResponse, error) {
	requestDef := GenReqDefForBatchUpdateAndCheckin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateAndCheckinResponse), nil
	}
}

// BatchUpdateAndCheckinInvoker XDM_批量更新并检入VersionModel
func (c *IDMEClassicAPIClient) BatchUpdateAndCheckinInvoker(request *model.BatchUpdateAndCheckinRequest) *BatchUpdateAndCheckinInvoker {
	requestDef := GenReqDefForBatchUpdateAndCheckin()
	return &BatchUpdateAndCheckinInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateAndRevise XDM_批量修订且更新VersionModel。
//
// 根据主对象ID批量修订对象并根据传入字段更新主对象+版本对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchUpdateAndRevise(request *model.BatchUpdateAndReviseRequest) (*model.BatchUpdateAndReviseResponse, error) {
	requestDef := GenReqDefForBatchUpdateAndRevise()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateAndReviseResponse), nil
	}
}

// BatchUpdateAndReviseInvoker XDM_批量修订且更新VersionModel。
func (c *IDMEClassicAPIClient) BatchUpdateAndReviseInvoker(request *model.BatchUpdateAndReviseRequest) *BatchUpdateAndReviseInvoker {
	requestDef := GenReqDefForBatchUpdateAndRevise()
	return &BatchUpdateAndReviseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateByAdmin XDM_管理员批量更新VersionModel指定属性
//
// 以管理员身份批量更新指定版本实例上的基础信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchUpdateByAdmin(request *model.BatchUpdateByAdminRequest) (*model.BatchUpdateByAdminResponse, error) {
	requestDef := GenReqDefForBatchUpdateByAdmin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateByAdminResponse), nil
	}
}

// BatchUpdateByAdminInvoker XDM_管理员批量更新VersionModel指定属性
func (c *IDMEClassicAPIClient) BatchUpdateByAdminInvoker(request *model.BatchUpdateByAdminRequest) *BatchUpdateByAdminInvoker {
	requestDef := GenReqDefForBatchUpdateByAdmin()
	return &BatchUpdateByAdminInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// BatchUpdateVersion XDM_批量升版最新版本对象VersionModel的版本号
//
// 根据ID批量升版最新版本对象数据的版本号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchUpdateVersion(request *model.BatchUpdateVersionRequest) (*model.BatchUpdateVersionResponse, error) {
	requestDef := GenReqDefForBatchUpdateVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateVersionResponse), nil
	}
}

// BatchUpdateVersionInvoker XDM_批量升版最新版本对象VersionModel的版本号
func (c *IDMEClassicAPIClient) BatchUpdateVersionInvoker(request *model.BatchUpdateVersionRequest) *BatchUpdateVersionInvoker {
	requestDef := GenReqDefForBatchUpdateVersion()
	return &BatchUpdateVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Checkin XDM_检入VersionModel
//
// 根据主对象ID检入版本对象，按照设置的规则生成新的业务版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) Checkin(request *model.CheckinRequest) (*model.CheckinResponse, error) {
	requestDef := GenReqDefForCheckin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckinResponse), nil
	}
}

// CheckinInvoker XDM_检入VersionModel
func (c *IDMEClassicAPIClient) CheckinInvoker(request *model.CheckinRequest) *CheckinInvoker {
	requestDef := GenReqDefForCheckin()
	return &CheckinInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Checkout XDM_检出VersionModel
//
// 根据主对象ID检出版本对象，复制生成一条新的版本记录且状态为已检出。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) Checkout(request *model.CheckoutRequest) (*model.CheckoutResponse, error) {
	requestDef := GenReqDefForCheckout()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckoutResponse), nil
	}
}

// CheckoutInvoker XDM_检出VersionModel
func (c *IDMEClassicAPIClient) CheckoutInvoker(request *model.CheckoutRequest) *CheckoutInvoker {
	requestDef := GenReqDefForCheckout()
	return &CheckoutInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckoutAndUpdate XDM_检出并更新VersionModel
//
// 根据主对象ID检出对象并根据传入字段更新版本对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) CheckoutAndUpdate(request *model.CheckoutAndUpdateRequest) (*model.CheckoutAndUpdateResponse, error) {
	requestDef := GenReqDefForCheckoutAndUpdate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckoutAndUpdateResponse), nil
	}
}

// CheckoutAndUpdateInvoker XDM_检出并更新VersionModel
func (c *IDMEClassicAPIClient) CheckoutAndUpdateInvoker(request *model.CheckoutAndUpdateRequest) *CheckoutAndUpdateInvoker {
	requestDef := GenReqDefForCheckoutAndUpdate()
	return &CheckoutAndUpdateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckoutUndo XDM_撤销检出VersionModel
//
// 根据主对象ID撤销检出版本对象，删除新的版本记录且状态为已检入。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) CheckoutUndo(request *model.CheckoutUndoRequest) (*model.CheckoutUndoResponse, error) {
	requestDef := GenReqDefForCheckoutUndo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckoutUndoResponse), nil
	}
}

// CheckoutUndoInvoker XDM_撤销检出VersionModel
func (c *IDMEClassicAPIClient) CheckoutUndoInvoker(request *model.CheckoutUndoRequest) *CheckoutUndoInvoker {
	requestDef := GenReqDefForCheckoutUndo()
	return &CheckoutUndoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckoutUndoByAdmin XDM_管理员撤销检出VersionModel
//
// 管理员根据主对象ID撤销检出版本对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) CheckoutUndoByAdmin(request *model.CheckoutUndoByAdminRequest) (*model.CheckoutUndoByAdminResponse, error) {
	requestDef := GenReqDefForCheckoutUndoByAdmin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckoutUndoByAdminResponse), nil
	}
}

// CheckoutUndoByAdminInvoker XDM_管理员撤销检出VersionModel
func (c *IDMEClassicAPIClient) CheckoutUndoByAdminInvoker(request *model.CheckoutUndoByAdminRequest) *CheckoutUndoByAdminInvoker {
	requestDef := GenReqDefForCheckoutUndoByAdmin()
	return &CheckoutUndoByAdminInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CompareBusinessVersion XDM_VersionModel业务版本对比
//
// 根据主对象id，输入版本号（大版本+小版本）进行版本属性与关系对比。（建议使用新的接口instance-attrs-comparison和\\ instance-relation-comparison比较属性和关系）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) CompareBusinessVersion(request *model.CompareBusinessVersionRequest) (*model.CompareBusinessVersionResponse, error) {
	requestDef := GenReqDefForCompareBusinessVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CompareBusinessVersionResponse), nil
	}
}

// CompareBusinessVersionInvoker XDM_VersionModel业务版本对比
func (c *IDMEClassicAPIClient) CompareBusinessVersionInvoker(request *model.CompareBusinessVersionRequest) *CompareBusinessVersionInvoker {
	requestDef := GenReqDefForCompareBusinessVersion()
	return &CompareBusinessVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// DeleteBranch XDM_删除VersionModel最新分支版本下所有小版本
//
// 根据masterid&amp;version删除最新大版本下的所有小版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) DeleteBranch(request *model.DeleteBranchRequest) (*model.DeleteBranchResponse, error) {
	requestDef := GenReqDefForDeleteBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBranchResponse), nil
	}
}

// DeleteBranchInvoker XDM_删除VersionModel最新分支版本下所有小版本
func (c *IDMEClassicAPIClient) DeleteBranchInvoker(request *model.DeleteBranchRequest) *DeleteBranchInvoker {
	requestDef := GenReqDefForDeleteBranch()
	return &DeleteBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// DeleteLatestVersion XDM_删除VersionModel最新分支的最新版本
//
// 根据主对象ID入参，删除最新分支的最新版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) DeleteLatestVersion(request *model.DeleteLatestVersionRequest) (*model.DeleteLatestVersionResponse, error) {
	requestDef := GenReqDefForDeleteLatestVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLatestVersionResponse), nil
	}
}

// DeleteLatestVersionInvoker XDM_删除VersionModel最新分支的最新版本
func (c *IDMEClassicAPIClient) DeleteLatestVersionInvoker(request *model.DeleteLatestVersionRequest) *DeleteLatestVersionInvoker {
	requestDef := GenReqDefForDeleteLatestVersion()
	return &DeleteLatestVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLogicalBranch XDM_软删除VersionModel最新分支版本下所有小版本
//
// 软删除最新分支版本下的所有小版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) DeleteLogicalBranch(request *model.DeleteLogicalBranchRequest) (*model.DeleteLogicalBranchResponse, error) {
	requestDef := GenReqDefForDeleteLogicalBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLogicalBranchResponse), nil
	}
}

// DeleteLogicalBranchInvoker XDM_软删除VersionModel最新分支版本下所有小版本
func (c *IDMEClassicAPIClient) DeleteLogicalBranchInvoker(request *model.DeleteLogicalBranchRequest) *DeleteLogicalBranchInvoker {
	requestDef := GenReqDefForDeleteLogicalBranch()
	return &DeleteLogicalBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLogicalLatestVersion XDM_软删除VersionModel最新分支的最新版本
//
// 根据主对象ID入参，软删最新分支的最新版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) DeleteLogicalLatestVersion(request *model.DeleteLogicalLatestVersionRequest) (*model.DeleteLogicalLatestVersionResponse, error) {
	requestDef := GenReqDefForDeleteLogicalLatestVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLogicalLatestVersionResponse), nil
	}
}

// DeleteLogicalLatestVersionInvoker XDM_软删除VersionModel最新分支的最新版本
func (c *IDMEClassicAPIClient) DeleteLogicalLatestVersionInvoker(request *model.DeleteLogicalLatestVersionRequest) *DeleteLogicalLatestVersionInvoker {
	requestDef := GenReqDefForDeleteLogicalLatestVersion()
	return &DeleteLogicalLatestVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ExecuteRevise XDM_修订VersionModel
//
// 根据主对象ID修订对象，按照设置的规则生成新的业务版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ExecuteRevise(request *model.ExecuteReviseRequest) (*model.ExecuteReviseResponse, error) {
	requestDef := GenReqDefForExecuteRevise()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteReviseResponse), nil
	}
}

// ExecuteReviseInvoker XDM_修订VersionModel
func (c *IDMEClassicAPIClient) ExecuteReviseInvoker(request *model.ExecuteReviseRequest) *ExecuteReviseInvoker {
	requestDef := GenReqDefForExecuteRevise()
	return &ExecuteReviseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllVersions XDM_获取VersionModel版本列表
//
// 根据主对象ID，获取全量版本以及对应版本对象list属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ListAllVersions(request *model.ListAllVersionsRequest) (*model.ListAllVersionsResponse, error) {
	requestDef := GenReqDefForListAllVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllVersionsResponse), nil
	}
}

// ListAllVersionsInvoker XDM_获取VersionModel版本列表
func (c *IDMEClassicAPIClient) ListAllVersionsInvoker(request *model.ListAllVersionsRequest) *ListAllVersionsInvoker {
	requestDef := GenReqDefForListAllVersions()
	return &ListAllVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowVersionByMaster XDM_获取VersionModel对应版本信息
//
// 根据Masterid和版本号和小版本号，返回对应版本属性，小版本号为空则返回最新小版本属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowVersionByMaster(request *model.ShowVersionByMasterRequest) (*model.ShowVersionByMasterResponse, error) {
	requestDef := GenReqDefForShowVersionByMaster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVersionByMasterResponse), nil
	}
}

// ShowVersionByMasterInvoker XDM_获取VersionModel对应版本信息
func (c *IDMEClassicAPIClient) ShowVersionByMasterInvoker(request *model.ShowVersionByMasterRequest) *ShowVersionByMasterInvoker {
	requestDef := GenReqDefForShowVersionByMaster()
	return &ShowVersionByMasterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAndCheckin XDM_更新并检入VersionModel
//
// 根据传入字段更新版本对象并根据主对象ID检入对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) UpdateAndCheckin(request *model.UpdateAndCheckinRequest) (*model.UpdateAndCheckinResponse, error) {
	requestDef := GenReqDefForUpdateAndCheckin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAndCheckinResponse), nil
	}
}

// UpdateAndCheckinInvoker XDM_更新并检入VersionModel
func (c *IDMEClassicAPIClient) UpdateAndCheckinInvoker(request *model.UpdateAndCheckinRequest) *UpdateAndCheckinInvoker {
	requestDef := GenReqDefForUpdateAndCheckin()
	return &UpdateAndCheckinInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAndRevise XDM_修订且更新VersionModel。
//
// 根据主对象ID修订对象并根据传入字段更新主对象+版本对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) UpdateAndRevise(request *model.UpdateAndReviseRequest) (*model.UpdateAndReviseResponse, error) {
	requestDef := GenReqDefForUpdateAndRevise()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAndReviseResponse), nil
	}
}

// UpdateAndReviseInvoker XDM_修订且更新VersionModel。
func (c *IDMEClassicAPIClient) UpdateAndReviseInvoker(request *model.UpdateAndReviseRequest) *UpdateAndReviseInvoker {
	requestDef := GenReqDefForUpdateAndRevise()
	return &UpdateAndReviseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateByAdmin XDM_管理员更新对象VersionModel指定属性
//
// 以管理员身份更新指定版本实例上的基础信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) UpdateByAdmin(request *model.UpdateByAdminRequest) (*model.UpdateByAdminResponse, error) {
	requestDef := GenReqDefForUpdateByAdmin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateByAdminResponse), nil
	}
}

// UpdateByAdminInvoker XDM_管理员更新对象VersionModel指定属性
func (c *IDMEClassicAPIClient) UpdateByAdminInvoker(request *model.UpdateByAdminRequest) *UpdateByAdminInvoker {
	requestDef := GenReqDefForUpdateByAdmin()
	return &UpdateByAdminInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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
