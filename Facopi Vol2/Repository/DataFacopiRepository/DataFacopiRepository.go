package datafacopirepository

import (
	"Facopi-Vol2/Models"
	connections "Facopi-Vol2/Models/Connection"
	DataFacopiModels "Facopi-Vol2/Models/DataFacopiModels"
)

func InsertDataFacopi(DataFacopi DataFacopiModels.DataFacopi) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "INSERT INTO DataFacopi (Kode, Nama_Kopi, Varian, Keterangan, Harga, Stok, Jumlah, Sub_Total) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	tempResult := db.Exec(query, DataFacopi.Kode, DataFacopi.Nama_Kopi, DataFacopi.Varian, DataFacopi.Keterangan, DataFacopi.Harga, DataFacopi.Stok, DataFacopi.Jumlah, DataFacopi.Sub_Total)

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


func UpdateDataFacopi(DataFacopi DataFacopiModels.DataFacopi) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "UPDATE DataFacopi SET Nama_Kopi = ?, Varian = ?, Keterangan = ?, Harga = ?, Stok = ?, Jumlah = ?, Sub_Total = ? WHERE Kode = ?"

	tempResult := db.Exec(query,DataFacopi.Nama_Kopi, DataFacopi.Varian, DataFacopi.Keterangan, DataFacopi.Harga, DataFacopi.Stok, DataFacopi.Jumlah, DataFacopi.Sub_Total, DataFacopi.Kode)

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


func DeleteDataFacopi(request DataFacopiModels.DataFacopi) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	db := connections.DB
	query = "DELETE FROM DataFacopi WHERE Kode = ?"

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

func GetDataFacopiByID(request DataFacopiModels.DataFacopi) Models.BaseResponseModels {
	var query string
	var result Models.BaseResponseModels
	var DataFacopi []DataFacopiModels.DataFacopi
	db := connections.DB

	if request.Kode != 0 {
		query = "SELECT * FROM DataFacopi WHERE Kode = ?"
	} else {
		query = "SELECT * FROM DataFacopi"
	}

	tempResult := db.Raw(query).Find(& DataFacopi)

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
			Data:         	DataFacopi,
		}
	}

	return result
}
