package physi_logi_def

type AccessInfo struct {
    tableID string
    filterID string
    panelID string
}

var tables = map[string]AccessInfo{
    "domain-ssl": { 
        tableID: "table_1704780398",
        filterID: "filter_1704786328",
        panelID:  "panel_1704780386",
    },
}

func GetTableID(alias string) string {
    return tables[alias].tableID
}
func GetPanelID(alias string) string {
    return tables[alias].panelID
}
func GetFilterID(alias string) string {
    return tables[alias].filterID
}
