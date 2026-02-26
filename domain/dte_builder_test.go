package domain

import (
	"strings"
	"testing"
)

func TestDTE33ToRequestBuildsDataDTEWithoutTEDOrTmstFirma(t *testing.T) {
	dte := Dte33Data{
		Encabezado: Encabezado33{
			IdDoc: IdDocBase{TipoDTE: 33, FchEmis: "2026-02-03"},
			Emisor: Emisor{
				RUTEmisor:  "76326028-3",
				RznSoc:     "SERVICIOS INFORMATICOS SIGANET LIMITADA",
				GiroEmis:   "Servicios de desarrollo de software",
				DirOrigen:  "Av. Apoquindo 3000",
				CmnaOrigen: "Las Condes",
			},
			Receptor: Receptor{RUTRecep: "18923860-6", RznSocRecep: "Cliente"},
			Totales:  Totales{MntNeto: 100000, IVA: 19000, MntTotal: 119000},
		},
		Detalle: []Detalle{{NroLinDet: 1, NmbItem: "Servicio", MontoItem: 100000}},
	}

	req, err := DTE33ToRequest("u1", "b1", "idem-1", dte)
	if err != nil {
		t.Fatalf("DTE33ToRequest error: %v", err)
	}

	if req.CodeSII != "33" {
		t.Fatalf("expected code 33, got %s", req.CodeSII)
	}
	if req.DataDTE == "" {
		t.Fatalf("expected non-empty data_dte")
	}
	if strings.Contains(req.DataDTE, "TED") {
		t.Fatalf("data_dte must not contain TED")
	}
	if strings.Contains(req.DataDTE, "TmstFirma") {
		t.Fatalf("data_dte must not contain TmstFirma")
	}
	if !strings.Contains(req.DataDTE, "Encabezado") {
		t.Fatalf("data_dte must contain Encabezado")
	}
}
