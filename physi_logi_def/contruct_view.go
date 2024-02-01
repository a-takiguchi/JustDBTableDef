package physi_logi_def
import "encoding/json"

type OutJsonContructView struct {
    ContructID json.Number `json:"ContructID"`
    KeiyakuID json.Number `json:"KeiyakuID"`
    KeiyakuName string `json:"KeiyakuName"`
    KeiyakuZACCode json.Number `json:"KeiyakuZACCode"`
    TeikyoID json.Number `json:"TeikyoID"`
    TeikyoName string `json:"TeikyoName"`
    TeikyoZACCode json.Number `json:"TeikyoZACCode"`
    ProjectName string `json:"ProjectName"`
    StartDate string `json:"StartDate"`
    EndDate string `json:"EndDate"`
    KeiyakuShoNumber string `json:"KeiyakuShoNumber"`
    ContructCommonID json.Number `json:"ContructCommonID"`
    ShohinID string `json:"ShohinID"`
    ShohinName string `json:"ShohinName"`
    PlanID string `json:"PlanID"`
    PlanName string `json:"PlanName"`
}

type OutCsvContructView struct {
    ContructID int64 `csv:"ContructID"`
    KeiyakuID int64 `csv:"KeiyakuID"`
    KeiyakuName string `csv:"KeiyakuName"`
    KeiyakuZACCode int64 `csv:"KeiyakuZACCode"`
    TeikyoID int64 `csv:"TeikyoID"`
    TeikyoName string `csv:"TeikyoName"`
    TeikyoZACCode int64 `csv:"TeikyoZACCode"`
    ProjectName string `csv:"ProjectName"`
    StartDate string `csv:"StartDate"`
    EndDate string `csv:"EndDate"`
    KeiyakuShoNumber string `csv:"KeiyakuShoNumber"`
    ContructCommonID int64 `csv:"ContructCommonID"`
    ShohinID string `csv:"ShohinID"`
    ShohinName string `csv:"ShohinName"`
    PlanID string `csv:"PlanID"`
    PlanName string `csv:"PlanName"`
}


const ContructID = "ContructID"
const KeiyakuID = "KeiyakuID"
const KeiyakuName = "KeiyakuName"
const KeiyakuZACCode = "KeiyakuZACCode"
const TeikyoID = "TeikyoID"
const TeikyoName = "TeikyoName"
const TeikyoZACCode = "TeikyoZACCode"
const ProjectName = "ProjectName"
const StartDate = "StartDate"
const EndDate = "EndDate"
const KeiyakuShoNumber = "KeiyakuShoNumber"
const ContructCommonID = "ContructCommonID"
const ShohinID = "ShohinID"
const ShohinName = "ShohinName"
const PlanID = "PlanID"
const PlanName = "PlanName"

func GetContructViewHeader() []string {
    return []string{
        ContructID,
        KeiyakuID,
        KeiyakuName,
        KeiyakuZACCode,
        TeikyoID,
        TeikyoName,
        TeikyoZACCode,
        ProjectName,
        StartDate,
        EndDate,
        KeiyakuShoNumber,
        ContructCommonID,
        ShohinID,
        ShohinName,
        PlanID,
        PlanName,
    }
}
