-- Table for User
CREATE TABLE User (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nama_user VARCHAR(255),
    kata_sandi VARCHAR(255),
    notelp UNIQUE VARCHAR(255),
    tanggal_lahir DATE,
    jenis_kelamin VARCHAR(255),
    tentang_text TEXT,
    pekerjaan VARCHAR(255),
    email VARCHAR(255),
    id_provinsi VARCHAR(255),
    id_kota VARCHAR(255),
    isAdmin BOOLEAN,
    updated_at_date TIMESTAMP,
    created_at_date TIMESTAMP
);

-- Table for Alamat
CREATE TABLE alamat (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_user INT,
    judul_alamat VARCHAR(255),
    nama_penerima VARCHAR(255),
    no_telp VARCHAR(255),
    detail_alamat VARCHAR(255),
    updated_at_date TIMESTAMP,
    created_at_date TIMESTAMP,
    FOREIGN KEY (id_user) REFERENCES User(id)
);

-- Table for Toko
CREATE TABLE toko (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_user INT,
    nama_toko VARCHAR(255),
    url_foto VARCHAR(255),
    updated_at_date TIMESTAMP,
    created_at_date TIMESTAMP,
    FOREIGN KEY (id_user) REFERENCES User(id)
);

-- Table for Kategori
CREATE TABLE category (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nama_category VARCHAR(255),
    created_at_date TIMESTAMP,
    updated_at_date TIMESTAMP
);

-- Table for Produk
CREATE TABLE produk (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nama_produk VARCHAR(255),
    slug VARCHAR(255),
    harga_reseller VARCHAR(255),
    harga_konsumen VARCHAR(255),
    stok INT,
    deskripsi TEXT,
    created_at_date TIMESTAMP,
    updated_at_date TIMESTAMP,
    id_category INT,
    FOREIGN KEY (id_category) REFERENCES category(id)
);

-- Table for Foto Produk
CREATE TABLE foto_produk (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_produk INT,
    url VARCHAR(255),
    updated_at_date TIMESTAMP,
    created_at_date TIMESTAMP,
    FOREIGN KEY (id_produk) REFERENCES produk(id)
);

-- Table for Log Produk
CREATE TABLE log_produk (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_produk INT,
    nama_produk VARCHAR(255),
    slug VARCHAR(255),
    harga_reseller VARCHAR(255),
    harga_konsumen VARCHAR(255),
    stok INT,
    deskripsi TEXT,
    created_at_date TIMESTAMP,
    updated_at_date TIMESTAMP,
    id_toko INT,
    id_category INT,
    FOREIGN KEY (id_produk) REFERENCES produk(id),
    FOREIGN KEY (id_toko) REFERENCES toko(id),
    FOREIGN KEY (id_category) REFERENCES category(id)
);

-- Table for Transaction (trx)
CREATE TABLE trx (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_user INT,
    alamat_pengiriman INT,
    harga_total INT,
    kode_invoice VARCHAR(255),
    method_bayar VARCHAR(255),
    updated_at_date TIMESTAMP,
    created_at_date TIMESTAMP,
    FOREIGN KEY (id_user) REFERENCES User(id),
    FOREIGN KEY (alamat_pengiriman) REFERENCES alamat(id)
);

-- Table for Detail Transaction (detail_trx)
CREATE TABLE detail_trx (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_trx INT,
    id_produk INT,
    kuantitas INT,
    harga_total INT,
    updated_at_date TIMESTAMP,
    created_at_date TIMESTAMP,
    FOREIGN KEY (id_trx) REFERENCES trx(id),
    FOREIGN KEY (id_produk) REFERENCES produk(id)
);
