package scan_history

import (
	"qrapi/g/x/web"
	"qrapi/x/logger"
	"qrapi/x/mongodb"
	"qrapi/x/validator"
)

var scanHistoryLog = logger.NewLogger("tbl_scan_history")
var scanHistoryTable = mongodb.NewTable("scan_history", "sh")

type ScanHistory struct {
	mongodb.Model `bson:",inline"`
	CustomerID    string `bson:"customer_id" json:"customer_id"`
	ProductID     string `bson:"product_id" json:"product_id"`
	NumberOfScan  string `bson:"number_of_scan" json:"number_of_scan"`
}

func (scanHistory *ScanHistory) Create() error {
	var existScanHistory *ScanHistory
	scanHistoryTable.FindId(scanHistory.ID).One(&existScanHistory)
	if existScanHistory != nil {

	}
	err := validator.Struct(scanHistory)
	if err != nil {
		scanHistoryLog.Error(err)
		return web.WrapBadRequest(err, "")
	}
	return scanHistoryTable.Create(scanHistory)
}