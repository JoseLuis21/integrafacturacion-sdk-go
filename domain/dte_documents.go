package domain

// DTE root payloads for data_dte (without TED and TmstFirma).

type Dte33Data struct {
	Encabezado   Encabezado33   `json:"Encabezado"`
	Detalle      []Detalle      `json:"Detalle"`
	DscRcgGlobal []DscRcgGlobal `json:"DscRcgGlobal,omitempty"`
	Referencia   []Referencia   `json:"Referencia,omitempty"`
	Comisiones   []Comisiones   `json:"Comisiones,omitempty"`
	SubTotInfo   []SubTotInfo   `json:"SubTotInfo,omitempty"`
	Transporte   *Transporte    `json:"Transporte,omitempty"`
	OtraMoneda   *OtraMoneda    `json:"OtraMoneda,omitempty"`
}

type Encabezado33 struct {
	IdDoc    IdDocBase `json:"IdDoc"`
	Emisor   Emisor    `json:"Emisor"`
	Receptor Receptor  `json:"Receptor"`
	Totales  Totales   `json:"Totales"`
}

type Dte34Data struct {
	Encabezado   Encabezado34   `json:"Encabezado"`
	Detalle      []Detalle      `json:"Detalle"`
	DscRcgGlobal []DscRcgGlobal `json:"DscRcgGlobal,omitempty"`
	Referencia   []Referencia   `json:"Referencia,omitempty"`
	Comisiones   []Comisiones   `json:"Comisiones,omitempty"`
	SubTotInfo   []SubTotInfo   `json:"SubTotInfo,omitempty"`
	Transporte   *Transporte    `json:"Transporte,omitempty"`
	OtraMoneda   *OtraMoneda    `json:"OtraMoneda,omitempty"`
}

type Encabezado34 struct {
	IdDoc    IdDocBase `json:"IdDoc"`
	Emisor   Emisor    `json:"Emisor"`
	Receptor Receptor  `json:"Receptor"`
	Totales  Totales   `json:"Totales"`
}

type Dte39Data struct {
	Encabezado   Encabezado39   `json:"Encabezado"`
	Detalle      []Detalle      `json:"Detalle"`
	DscRcgGlobal []DscRcgGlobal `json:"DscRcgGlobal,omitempty"`
	Referencia   []Referencia   `json:"Referencia,omitempty"`
	SubTotInfo   []SubTotInfo   `json:"SubTotInfo,omitempty"`
	OtraMoneda   *OtraMoneda    `json:"OtraMoneda,omitempty"`
}

type Encabezado39 struct {
	IdDoc    IdDocBoleta    `json:"IdDoc"`
	Emisor   EmisorBoleta   `json:"Emisor"`
	Receptor ReceptorBoleta `json:"Receptor"`
	Totales  TotalesBoleta  `json:"Totales"`
}

type Dte41Data struct {
	Encabezado   Encabezado41   `json:"Encabezado"`
	Detalle      []Detalle      `json:"Detalle"`
	DscRcgGlobal []DscRcgGlobal `json:"DscRcgGlobal,omitempty"`
	Referencia   []Referencia   `json:"Referencia,omitempty"`
	SubTotInfo   []SubTotInfo   `json:"SubTotInfo,omitempty"`
	OtraMoneda   *OtraMoneda    `json:"OtraMoneda,omitempty"`
}

type Encabezado41 struct {
	IdDoc    IdDocBoleta    `json:"IdDoc"`
	Emisor   EmisorBoleta   `json:"Emisor"`
	Receptor ReceptorBoleta `json:"Receptor"`
	Totales  TotalesBoleta  `json:"Totales"`
}

type Dte46Data struct {
	Encabezado   Encabezado46   `json:"Encabezado"`
	Detalle      []Detalle      `json:"Detalle"`
	DscRcgGlobal []DscRcgGlobal `json:"DscRcgGlobal,omitempty"`
	Referencia   []Referencia   `json:"Referencia,omitempty"`
	Comisiones   []Comisiones   `json:"Comisiones,omitempty"`
	SubTotInfo   []SubTotInfo   `json:"SubTotInfo,omitempty"`
	Transporte   *Transporte    `json:"Transporte,omitempty"`
	OtraMoneda   *OtraMoneda    `json:"OtraMoneda,omitempty"`
}

type Encabezado46 struct {
	IdDoc    IdDocBase `json:"IdDoc"`
	Emisor   Emisor    `json:"Emisor"`
	Receptor Receptor  `json:"Receptor"`
	Totales  Totales   `json:"Totales"`
}

type Dte52Data struct {
	Encabezado   Encabezado52    `json:"Encabezado"`
	Detalle      []Detalle       `json:"Detalle"`
	DscRcgGlobal []DscRcgGlobal  `json:"DscRcgGlobal,omitempty"`
	Referencia   []Referencia    `json:"Referencia,omitempty"`
	Transporte   *TransporteGuia `json:"Transporte,omitempty"`
	Comisiones   []Comisiones    `json:"Comisiones,omitempty"`
	SubTotInfo   []SubTotInfo    `json:"SubTotInfo,omitempty"`
	OtraMoneda   *OtraMoneda     `json:"OtraMoneda,omitempty"`
}

type Encabezado52 struct {
	IdDoc    IdDocGuia   `json:"IdDoc"`
	Emisor   Emisor      `json:"Emisor"`
	Receptor Receptor    `json:"Receptor"`
	Totales  TotalesGuia `json:"Totales"`
}

type Dte56Data struct {
	Encabezado   Encabezado56   `json:"Encabezado"`
	Detalle      []Detalle      `json:"Detalle"`
	DscRcgGlobal []DscRcgGlobal `json:"DscRcgGlobal,omitempty"`
	Referencia   []Referencia   `json:"Referencia,omitempty"`
	Comisiones   []Comisiones   `json:"Comisiones,omitempty"`
	SubTotInfo   []SubTotInfo   `json:"SubTotInfo,omitempty"`
	Transporte   *Transporte    `json:"Transporte,omitempty"`
	OtraMoneda   *OtraMoneda    `json:"OtraMoneda,omitempty"`
}

type Encabezado56 struct {
	IdDoc    IdDocBase `json:"IdDoc"`
	Emisor   Emisor    `json:"Emisor"`
	Receptor Receptor  `json:"Receptor"`
	Totales  Totales   `json:"Totales"`
}

type Dte61Data struct {
	Encabezado   Encabezado61   `json:"Encabezado"`
	Detalle      []Detalle      `json:"Detalle"`
	DscRcgGlobal []DscRcgGlobal `json:"DscRcgGlobal,omitempty"`
	Referencia   []Referencia   `json:"Referencia,omitempty"`
	Comisiones   []Comisiones   `json:"Comisiones,omitempty"`
	SubTotInfo   []SubTotInfo   `json:"SubTotInfo,omitempty"`
	Transporte   *Transporte    `json:"Transporte,omitempty"`
	OtraMoneda   *OtraMoneda    `json:"OtraMoneda,omitempty"`
}

type Encabezado61 struct {
	IdDoc    IdDocBase `json:"IdDoc"`
	Emisor   Emisor    `json:"Emisor"`
	Receptor Receptor  `json:"Receptor"`
	Totales  Totales   `json:"Totales"`
}
