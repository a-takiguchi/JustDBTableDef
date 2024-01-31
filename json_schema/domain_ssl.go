package json_schema

import "encoding/json"

type DomainSSL struct {
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
    Record DomainSSLRecord `json:"record"`
}

type DomainSSLRecord struct {
    Project string `json:"field_1684834053_K"`
    Type string `json:"field_1670911377"`
    Customer string `json:"field_1670909964_K22"`
    AWSAccountID string `json:"field_1704780408"`
    Registrar string `json:"field_1704781328"`
    RemainDays Day `json:"field_1704783209"`
    CommonName string `json:"field_1704780847"`
    ExpireDate string `json:"field_1704780864"`
    DomainName string `json:"field_1704780832"`
    PIC []string `json:"field_1704781767"`
    Status []string `json:"field_1704781365"`
    Note string `json:"field_1704781345"`
    ID []string `json:"field_1704782697"`
}


