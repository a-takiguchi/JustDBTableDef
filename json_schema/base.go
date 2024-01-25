package json_schema

import (
	"encoding/json"
)

type Base struct {
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
    Record interface{} `json:"record"`
}

type Day struct {
    Days json.Number `json:"days"`
    Hours json.Number `json:"hours"`
    Minutes json.Number `json:"minutes"`
    WholeSeconds json.Number `json:"wholeSeconds"`
    MicorSeconds json.Number `json:"micorSeconds"`
}

