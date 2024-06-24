create database Facopi;

create table DataFacopi (
    Kode int primary key
    ,Nama_Kopi varchar(100)
    ,Varian varchar(100)
    ,Keterangan varchar(300)
    ,Harga float
    ,Stok float
    ,Jumlah float
    ,Sub_Total DECIMAL
);

create table KeuanganFakopi (
    Kode int primary key
    ,Sub_Day float
    ,Sub_Week float
    ,Sub_Month float
    ,Sub_Year float
    ,Penjualan_Masuk int
    ,Penjualan_Keluar int
    ,Penjualan_Bersih int
    ,Total_Penjualan int
    ,Jumlah_Cup_Terjual float
    ,Pemasukan_Week float
    ,Pemasukan_Month float
    Pengeluaran_Week float
    ,Pengeluaran_Month float
    ,Total_Week DECIMAL
    ,Total_Month DECIMAL
    ,Keterangan varchar(300)
);