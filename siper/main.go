package main

import "fmt"

type ListPersetujuan struct {
	UUID             string `db:"uuid" json:"id"`
	TahunAnggaran    int    `db:"tahun_anggaran" json:"tahun_anggaran"`
	NotaStatusKode   string `db:"nota_status_kode" json:"-"`
	NotaStatusNama   string `db:"nota_status_nama" json:"-"`
	Process          string `db:"process_code" json:"process_code"`
	TanggalNotaDinas string `db:"tanggal_nota_dinas" json:"tanggal_nota_dinas"`
	DalamRangka      any    `db:"dalam_rangka" json:"dalam_rangka"`
	SubKegiatan      string `db:"sub_kegiatan" json:"sub_kegiatan"`
	TotalRows        int    `db:"total_rows" json:"-"`
}

func main() {
	var listKetuaTimkerja []ListPersetujuan
	var listKetuaKelompok []ListPersetujuan
	var listPPK []ListPersetujuan

	for _, d := range []string{"E", "D"} {
		switch d {
		case "E":
			listKetuaTimkerja = Data1()
		case "D":
			listKetuaKelompok = Data2()
		case "F":
			listPPK = Data3()
		}
	}

	dataStep := listKetuaTimkerja
	dataStep = CompareCurrentPersetujuan(dataStep, listKetuaKelompok)
	dataStep = CompareCurrentPersetujuan(dataStep, listPPK)

	for _, d := range dataStep {
		fmt.Println(d)
	}
}

func Data1() (data []ListPersetujuan) {
	data = []ListPersetujuan{
		{
			UUID:           "aa",
			NotaStatusKode: "A",
		},
		{
			UUID:           "bb",
			NotaStatusKode: "A",
		},
		{
			UUID:           "cc",
			NotaStatusKode: "A",
		},
		{
			UUID:           "dd",
			NotaStatusKode: "A",
		},
	}
	return
}

func Data2() (data []ListPersetujuan) {
	data = []ListPersetujuan{
		{
			UUID:           "aa",
			NotaStatusKode: "B",
		},
		{
			UUID:           "bb",
			NotaStatusKode: "B",
		},
	}
	return
}

func Data3() (data []ListPersetujuan) {
	data = []ListPersetujuan{
		{
			UUID:           "aa",
			NotaStatusKode: "C",
		},
	}
	return
}

func CompareCurrentPersetujuan(before []ListPersetujuan, after []ListPersetujuan) (merger []ListPersetujuan) {
	if len(before) == 0 {
		return after
	}
	if len(after) == 0 {
		return before
	}

	uniqueUUIDs := make(map[string]ListPersetujuan)
	for _, d := range before {
		uniqueUUIDs[d.UUID] = d
	}
	for _, d := range after {
		uniqueUUIDs[d.UUID] = d
	}
	for _, d := range uniqueUUIDs {
		merger = append(merger, d)
	}
	return
}
