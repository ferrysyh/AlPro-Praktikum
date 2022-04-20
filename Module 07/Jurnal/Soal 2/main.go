package main

import "fmt"

func main() {
	var hari, bulan string
	var tanggal, tahun int

	fmt.Scanln(&hari, &tanggal, &bulan, &tahun)
	var nomorBulan int = angkaBulan(bulan)
	pengambilan(tanggal, angkaBulan(bulan), tahun, hari, &tanggal, &nomorBulan, &tahun, &hari)
}

func kabisat(tahun int) bool {
	if tahun%4 == 0 && tahun%100 != 0 || tahun%400 == 0 {
		return true
	} else {
		return false
	}
}

func angkaBulan(bulan string) int {
	if bulan == "januari" {
		return 1
	} else if bulan == "februari" {
		return 2
	} else if bulan == "maret" {
		return 3
	} else if bulan == "april" {
		return 4
	} else if bulan == "mei" {
		return 5
	} else if bulan == "juni" {
		return 6
	} else if bulan == "juli" {
		return 7
	} else if bulan == "agustus" {
		return 8
	} else if bulan == "september" {
		return 9
	} else if bulan == "oktober" {
		return 10
	} else if bulan == "november" {
		return 11
	} else if bulan == "desember" {
		return 12
	} else {
		return -1
	}
}

func bulanAngka(angka int) string {
	switch angka {
	case 1:
		return "januari"
	case 2:
		return "februari"
	case 3:
		return "maret"
	case 4:
		return "april"
	case 5:
		return "mei"
	case 6:
		return "juni"
	case 7:
		return "juli"
	case 8:
		return "agustus"
	case 9:
		return "september"
	case 10:
		return "oktober"
	case 11:
		return "november"
	case 12:
		return "desember"
	default:
		return "invalid"
	}
}

func jumlahHari(bulan, tahun int) int {
	switch bulan {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if kabisat(tahun) {
			return 29
		} else {
			return 28
		}
	default:
		return -1
	}
}

func mencariDurasi(hari1 string, hari2 *string, durasi *int) {
	switch hari1 {
	case "senin":
		*hari2 = "rabu"
		*durasi = 2
	case "selasa":
		*hari2 = "kamis"
		*durasi = 2
	case "rabu":
		*hari2 = "jumat"
		*durasi = 2
	case "kamis":
		*hari2 = "senin"
		*durasi = 4
	case "jumat":
		*hari2 = "selasa"
		*durasi = 4
	case "sabtu":
		*hari2 = "selasa"
		*durasi = 3
	case "minggu":
		*hari2 = "selasa"
		*durasi = 2
	default:
		*hari2 = "invalid"
		*durasi = -1
	}
}

func pengambilan(tanggal1, bulan1, tahun1 int, hari1 string, tanggal2, bulan2, tahun2 *int, hari2 *string) {
	var durasi int
	var hari string
	mencariDurasi(hari1, &hari, &durasi)
	*tanggal2 = tanggal1 + durasi
	*bulan2 = bulan1
	*tahun2 = tahun1
	if *tanggal2 > jumlahHari(*bulan2, *tahun2) {
		*tanggal2 = *tanggal2 - jumlahHari(*bulan2, *tahun2)
		*bulan2 = *bulan2 + 1
		if *bulan2 > 12 {
			*bulan2 = *bulan2 - 12
			*tahun2 = *tahun2 + 1
		}
	}
	*hari2 = hari
	fmt.Println(*hari2, *tanggal2, bulanAngka(*bulan2), *tahun2)
}
