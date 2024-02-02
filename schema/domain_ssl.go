package schema

import "encoding/json"

type InDomainSSLRecord struct {
    ProjectName string `json:"field_1684834053_K"`
    Type string `json:"field_1670911377"`
    Customer string `json:"field_1670909964_K22"`
    AWSAccountID string `json:"field_1704780408"`
    Registrar string `json:"field_1704781328"`
    RemainDays DayStruct `json:"field_1704783209"`
    CommonName string `json:"field_1704780847"`
    ExpireDate string `json:"field_1704780864"`
    DomainName string `json:"field_1704780832"`
    PIC []string `json:"field_1704781767"`
    Status []string `json:"field_1704781365"`
    Note string `json:"field_1704781345"`
    ID []string `json:"field_1704782697"`
}

type OutDomainSSL struct {
    RecordId json.Number `json:"recordId" csv:"recordId"`
    ProjectName string `json:"ProjectName" csv:"ProjectName"`
    Type string `json:"Type" csv:"Type"`
    Customer string `json:"CustomerName" csv:"CustomerName"`
    AWSAccountID string `json:"AWSAccountID" csv:"AWSAccountID"`
    Registrar string `json:"Registrar" csv:"Registrar"`
    RemainDays json.Number `json:"RemainDays" csv:"RemainDays"`
    CommonName string `json:"CommonName" csv:"CommonName"`
    ExpireDate string `json:"ExpireDate" csv:"ExpireDate"`
    DomainName string `json:"DomainName" csv:"DomainName"`
    PIC string `json:"PIC" csv:"PIC"`
    Status string `json:"Status" csv:"Status"`
    Note string `json:"Note" csv:"Note"`
    ID string `json:"ID" csv:"ID"`
    CreatedBy string `json:"createdBy" csv:"createdBy"`
    CreatedAt string `json:"createdAt" csv:"createdAt"`
    UpdatedBy string `json:"updatedBy" csv:"updatedBy"`
    UpdatedAt string `json:"updatedAt" csv:"updatedAt"`
}




