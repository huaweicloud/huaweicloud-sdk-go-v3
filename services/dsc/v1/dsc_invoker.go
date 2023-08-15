package v1

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/dsc/v1/model"
)

type AddBucketsInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddBucketsInvoker) Invoke() (*model.AddBucketsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddBucketsResponse), nil
	}
}

type AddRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddRuleInvoker) Invoke() (*model.AddRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddRuleResponse), nil
	}
}

type AddRuleGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddRuleGroupInvoker) Invoke() (*model.AddRuleGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddRuleGroupResponse), nil
	}
}

type AddScanJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddScanJobInvoker) Invoke() (*model.AddScanJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddScanJobResponse), nil
	}
}

type BatchAddDataMaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddDataMaskInvoker) Invoke() (*model.BatchAddDataMaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddDataMaskResponse), nil
	}
}

type ChangeDbTemplateOperationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeDbTemplateOperationInvoker) Invoke() (*model.ChangeDbTemplateOperationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeDbTemplateOperationResponse), nil
	}
}

type ChangeRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeRuleInvoker) Invoke() (*model.ChangeRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeRuleResponse), nil
	}
}

type CreateDatabaseWaterMarkInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDatabaseWaterMarkInvoker) Invoke() (*model.CreateDatabaseWaterMarkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDatabaseWaterMarkResponse), nil
	}
}

type CreateDocWatermarkInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDocWatermarkInvoker) Invoke() (*model.CreateDocWatermarkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDocWatermarkResponse), nil
	}
}

type CreateDocWatermarkByAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDocWatermarkByAddressInvoker) Invoke() (*model.CreateDocWatermarkByAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDocWatermarkByAddressResponse), nil
	}
}

type CreateImageWatermarkInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateImageWatermarkInvoker) Invoke() (*model.CreateImageWatermarkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateImageWatermarkResponse), nil
	}
}

type CreateImageWatermarkByAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateImageWatermarkByAddressInvoker) Invoke() (*model.CreateImageWatermarkByAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateImageWatermarkByAddressResponse), nil
	}
}

type CreateProductOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProductOrderInvoker) Invoke() (*model.CreateProductOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProductOrderResponse), nil
	}
}

type DeleteBucketInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBucketInvoker) Invoke() (*model.DeleteBucketResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBucketResponse), nil
	}
}

type DeleteRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRuleInvoker) Invoke() (*model.DeleteRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRuleResponse), nil
	}
}

type DeleteRuleGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRuleGroupInvoker) Invoke() (*model.DeleteRuleGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRuleGroupResponse), nil
	}
}

type ListBucketsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBucketsInvoker) Invoke() (*model.ListBucketsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBucketsResponse), nil
	}
}

type ListDbMaskTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDbMaskTaskInvoker) Invoke() (*model.ListDbMaskTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDbMaskTaskResponse), nil
	}
}

type ListRelationBucketsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRelationBucketsInvoker) Invoke() (*model.ListRelationBucketsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRelationBucketsResponse), nil
	}
}

type ListRelationColumnInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRelationColumnInvoker) Invoke() (*model.ListRelationColumnResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRelationColumnResponse), nil
	}
}

type ListRelationDbInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRelationDbInvoker) Invoke() (*model.ListRelationDbResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRelationDbResponse), nil
	}
}

type ListRelationFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRelationFileInvoker) Invoke() (*model.ListRelationFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRelationFileResponse), nil
	}
}

type ListRelationTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRelationTableInvoker) Invoke() (*model.ListRelationTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRelationTableResponse), nil
	}
}

type ListRuleGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRuleGroupsInvoker) Invoke() (*model.ListRuleGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRuleGroupsResponse), nil
	}
}

type ShowDatabaseWaterMarkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDatabaseWaterMarkInvoker) Invoke() (*model.ShowDatabaseWaterMarkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDatabaseWaterMarkResponse), nil
	}
}

type ShowDocWatermarkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDocWatermarkInvoker) Invoke() (*model.ShowDocWatermarkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDocWatermarkResponse), nil
	}
}

type ShowDocWatermarkByAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDocWatermarkByAddressInvoker) Invoke() (*model.ShowDocWatermarkByAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDocWatermarkByAddressResponse), nil
	}
}

type ShowImageWatermarkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageWatermarkInvoker) Invoke() (*model.ShowImageWatermarkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageWatermarkResponse), nil
	}
}

type ShowImageWatermarkByAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageWatermarkByAddressInvoker) Invoke() (*model.ShowImageWatermarkByAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageWatermarkByAddressResponse), nil
	}
}

type ShowImageWatermarkWithImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageWatermarkWithImageInvoker) Invoke() (*model.ShowImageWatermarkWithImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageWatermarkWithImageResponse), nil
	}
}

type ShowImageWatermarkWithImageByAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageWatermarkWithImageByAddressInvoker) Invoke() (*model.ShowImageWatermarkWithImageByAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageWatermarkWithImageByAddressResponse), nil
	}
}

type ShowRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRulesInvoker) Invoke() (*model.ShowRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRulesResponse), nil
	}
}

type ShowScanJobResultsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScanJobResultsInvoker) Invoke() (*model.ShowScanJobResultsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScanJobResultsResponse), nil
	}
}

type ShowScanJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScanJobsInvoker) Invoke() (*model.ShowScanJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScanJobsResponse), nil
	}
}

type ShowSpecificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSpecificationInvoker) Invoke() (*model.ShowSpecificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSpecificationResponse), nil
	}
}

type ShowTopicsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTopicsInvoker) Invoke() (*model.ShowTopicsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTopicsResponse), nil
	}
}

type UpdateAssetNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAssetNameInvoker) Invoke() (*model.UpdateAssetNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAssetNameResponse), nil
	}
}

type UpdateDefaultTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDefaultTopicInvoker) Invoke() (*model.UpdateDefaultTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDefaultTopicResponse), nil
	}
}

type ShowOpenApiCalledRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOpenApiCalledRecordsInvoker) Invoke() (*model.ShowOpenApiCalledRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOpenApiCalledRecordsResponse), nil
	}
}
