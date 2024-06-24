package keuanganfacopirepository

import (
	"Facopi-Vol2/Models"
	connection "Facopi-Vol2/Models/Connection"
	KeuanganFacopiModels "Facopi-Vol2/Models/KeuanganFacopiModels"
)

func InsertKeuanganFacopi(KeuanganFakopi KeuanganFacopiModels.KeuanganFakopi) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connection.DB

	query = "INSERT INTO KeuanganFakopi (Kode, Sub_Day, Sub_Week, Sub_Month, Sub_Year, Penjualan_Masuk, Penjualan_Keluar, Penjualan_Bersih, Total_Penjualan, Jumlah_Cup_Terjual, Pemasukan_Week, Pemasukan_Month, Pengeluaran_Week, Pengeluaran_Month, Total_Week, Total_Month, Keterangan) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	tempResult := db.Exec(query, KeuanganFakopi.Kode, KeuanganFakopi.Sub_Day, KeuanganFakopi.Sub_Week, KeuanganFakopi.Sub_Month, KeuanganFakopi.Sub_Year, KeuanganFakopi.Penjualan_Masuk, KeuanganFakopi.Penjualan_Keluar, KeuanganFakopi.Penjualan_Bersih, KeuanganFakopi.Total_Penjualan, KeuanganFakopi.Jumlah_Cup_Terjual, KeuanganFakopi.Pemasukan_Week, KeuanganFakopi.Pemasukan_Month, KeuanganFakopi.Pengeluaran_Week, KeuanganFakopi.Pengeluaran_Month, KeuanganFakopi.Total_Week, KeuanganFakopi.Total_Month, KeuanganFakopi.Keterangan)

	if tempResult.Error != nil {
		result = Models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		result = Models.BaseResponseModels{
			CodeResponse:  200,
			HeaderMessage: "Success",
			Message:       "Data berhasil ditambahkan.",
			Data:          nil,
		}
	}
	return result
}

func UpdateKeuanganFacopi(KeuanganFakopi KeuanganFacopiModels.KeuanganFakopi) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connection.DB

	query = "UPDATE KeuanganFakopi SET Sub_Day = ?, Sub_week = ?, Sub_Month = ?, Sub_Year = ?, Penjualan_Masuk = ?, Penjualan_Keluar = ?, Penjualan_Bersih = ?, Total_Penjualan = ?, Jumlah_Cup_Terjual = ?, Pemasukan_Week = ?, Pemasukan_Month = ?, Pengeluaran_Week = ?, Pengeluaran_Month = ?, Total_Week = ?, Total_Month = ?, Keterangan = ? WHERE Kode = ?"

	tempResult := db.Exec(query, KeuanganFakopi.Sub_Day, KeuanganFakopi.Sub_Week, KeuanganFakopi.Sub_Month, KeuanganFakopi.Sub_Year, KeuanganFakopi.Penjualan_Masuk, KeuanganFakopi.Penjualan_Keluar, KeuanganFakopi.Penjualan_Bersih, KeuanganFakopi.Total_Penjualan, KeuanganFakopi.Jumlah_Cup_Terjual, KeuanganFakopi.Pemasukan_Week, KeuanganFakopi.Pemasukan_Month, KeuanganFakopi.Pengeluaran_Week, KeuanganFakopi.Pengeluaran_Month, KeuanganFakopi.Total_Week, KeuanganFakopi.Total_Month, KeuanganFakopi.Keterangan, KeuanganFakopi.Kode)

	if tempResult.Error != nil {
		result = Models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		rowsAffected := tempResult.RowsAffected
		if rowsAffected == 0 {
			result = Models.BaseResponseModels{
				CodeResponse:  404,
				HeaderMessage: "Not Found",
				Message:       "Data dengan ID tersebut tidak ditemukan.",
				Data:          nil,
			}
		} else {
			result = Models.BaseResponseModels{
				CodeResponse:  200,
				HeaderMessage: "Success",
				Message:       "Data berhasil diubah.",
				Data:          nil,
			}
		}
	}

	return result
}

func DeleteKeuanganFacopi(request KeuanganFacopiModels.KeuanganFakopi) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connection.DB

	query = "DELETE FROM KeuanganFakopi WHERE Kode = ?"

	tempResult := db.Exec(query, request.Kode)

	if tempResult.Error != nil {
		result = Models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {
		// Periksa apakah ada baris yang terpengaruh oleh perintah DELETE
		rowsAffected := tempResult.RowsAffected
		if rowsAffected == 0 {
			result = Models.BaseResponseModels{
				CodeResponse:  404,
				HeaderMessage: "Not Found",
				Message:       "Data dengan ID tersebut tidak ditemukan.",
				Data:          nil,
			}
		} else {
			result = Models.BaseResponseModels{
				CodeResponse:  200,
				HeaderMessage: "Success",
				Message:       "Data berhasil dihapus.",
				Data:          nil,
			}
		}
	}

	return result
}

func GetKeuanganFacopiByID(request KeuanganFacopiModels.KeuanganFakopi) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	var KeuanganFacopi []KeuanganFacopiModels.KeuanganFakopi
	db := connection.DB

	if request.Kode != 0 {
		query = "SELECT * FROM KeuanganFakopi WHERE Kode = ?"
	} else {
		query = "SELECT * FROM KeuanganFakopi"
	}

	tempResult := db.Raw(query).Find(& KeuanganFacopi)

	if tempResult.Error != nil {
		result = Models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Error",
			Message:       tempResult.Error.Error(),
			Data:          nil,
		}
	} else {	
		result = Models.BaseResponseModels{
			CodeResponse:  200,
			HeaderMessage: "Success",
			Message:       "Data retrieved successfully",
			Data:         	KeuanganFacopi,
		}
	}

	return result
}



