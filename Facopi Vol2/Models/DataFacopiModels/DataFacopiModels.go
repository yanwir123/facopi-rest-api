package datafacopimodels

import "github.com/shopspring/decimal"

type DataFacopi struct {
	Kode       int64   `json:"Kode" from:"Kode"`
	Nama_Kopi  string  `json:"Nama Kopi" from:"Nama Kopi"`
	Varian     string  `json:"Varian" from:"Varian"`
	Keterangan string  `json:"Keterangan" from:"Keterangan"`
	Harga      float64 `json:"Harga" from:"Harga"`
	Stok       float64 `json:"Stok" from:"Stok"`
	Jumlah     float64 `json:"Jumlah" from:"Jumlah"`
	Sub_Total  decimal.Decimal `json:"Sub_Total" from:"Sub_Total"`
}