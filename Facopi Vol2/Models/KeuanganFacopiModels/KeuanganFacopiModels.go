package keuanganfacopimodels

import "github.com/shopspring/decimal"

type KeuanganFakopi struct {
	Kode               int64   `json:"Kode" from:"Kode"`
	Sub_Day            float64 `json:"Sub_Day" from:"Sub_Day"`
	Sub_Week           float64 `json:"Sub_Week" from:"Sub_Week"`
	Sub_Month          float64 `json:"Sub_Month" from:"Sub_Month"`
	Sub_Year           float64 `json:"Sub_Year" from:"Sub_Year"`
	Penjualan_Masuk    int64   `json:"Penjulan Masuk" from:"Penjualan Masuk"`
	Penjualan_Keluar   int64   `json:"Penjualan Keluar" from:"Penjualan Keluar"`
	Penjualan_Bersih   int64   `json:"Penjualan Bersih" from:"Penjualan Bersih"`
	Total_Penjualan    int64   `json:"Total Penjualan" from:"Total Penjualan"`
	Jumlah_Cup_Terjual float64 `json:"Jumlah Cup Terjual" from:"Jumlah Cup Terjual"`
	Pemasukan_Week     float64 `json:"Pemasukan Week" from:"Pemasukan_Week"`
	Pemasukan_Month    float64 `json:"Pemasukan Month" from:"Pemasukan_Month"`
	Pengeluaran_Week   float64 `json:"Pengeluaran_Week" From:"Pengeluaran Week"`
	Pengeluaran_Month  float64 `json:"Pengeluaran_Month" from:"Pengeluaran Month"`
	Total_Week         decimal.Decimal `json:"Total Week" from:"Total Week"`
	Total_Month        decimal.Decimal `json:"Total Month" from:"Total Month"`
	Keterangan 		   string `json:"Keterangan" from:"Keterangan"`
}