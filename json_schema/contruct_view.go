package json_schema

import (
	"encoding/json"
)

type ContructView struct {
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
    Record ContructViewRecord `json:"record"`
}

type ContructViewRecord struct {
    Project string `json:"field_1684834053_K"`
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
    RecordID json.Number `json:"field_1705998856"`
    ContructID json.Number `json:"field_1705998887"`
    TeikyoID json.Number `json:"field_1705998905"`
    TeikyoName string `json:"field_1705998924"`
    KeiyakuID json.Number `json:"field_1705998951"`
    KeiyakuName string `json:"field_1705998968"`
    TeikyoZACCode string `json:"field_1705998986"`
    KeiyakuZACCode string `json:"field_1705999004"`
    ProjectName string `json:"field_1705999024"`
    ContructDate string `json:"field_1705999056"`
    ContructEndDate string `json:"field_1705999080"`
    ContructNo string `json:"field_1705999098"`
    ShohinID string `json:"field_1705999118"`
    ShohinName string `json:"field_1705999133"`
    PlanID string `json:"field_1705999151"`
    PlanName string `json:"field_1705999162"`
}

