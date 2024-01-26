package json_schema

import "encoding/json"

type DomainSSLUserConn struct {
    TableName string `json:"tableName"`
    RecordId json.Number `json:"recordId"`
    Reactions []string `json:"reactions"`
    Comments []string `json:"comments"`
    CreatedBy []string `json:"createdBy"`
    CreatedAt string `json:"createdAt"`
    UpdatedBy []string `json:"updatedBy"`
    UpdatedAt string `json:"updatedAt"`
    WorkflowName string `json:"workflowName"`
    InsertUser []string `json:"insertUser"`
    SubstituteUser []string `json:"substituteUser"`
    NextApprovalUsers []string `json:"nextApprovalUsers"`
    SubmittedAt string `json:"submittedAt"`
    LastApprovalUser []string `json:"lastApprovalUser"`
    ApproveAt []string `json:"approveAt"`
    ApprovalUser []string `json:"approvalUser"`
    ApprovalStatus string `json:"approvalStatus"`
    RecordLink string `json:"recordLink"`
    // Record json.RawMessage `json:"record"`
    Record DomainSSLUserConnRecord `json:"record"`
}

type DomainSSLUserConnRecord struct {
    ID json.Number `json:"field_1705543425"`
    Note string `json:"field_1705543479"`
    ContructID []string `json:"field_1706239175"`
    ProjectName string `json:"field_1684834053"`
    CustomerName string `json:"field_1670909964"`
    SlackMasterID []string `json:"field_1706239242"`
    User []string `json:"field_1706238968"`
    RoleName string `json:"field_1670910619"`
    SlackID string `json:"field_1706239002"`
}


