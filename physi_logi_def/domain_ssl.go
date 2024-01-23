package physi_logi_def


var domainSSL = map[string]string{
    "field_1704780408": "AWSアカウントID",
    "field_1704780832": "ドメイン名",
    "field_1704780847": "コモンネーム",
    "field_1704780864": "有効期限",
    "field_1704781328": "レジストラ",
    "field_1704781345": "備考",
    "field_1704781365": "ステータス",
    "field_1704781767": "担当者",
    "field_1704783209": "残日数",
    "field_1704782697": "契約共通ID",
}


func GetJPName(colID string) string {
    return domainSSL[colID]
}
