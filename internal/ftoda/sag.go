package ftoda

type Sag struct {
	Id                int    `gorm:"primaryKey" json:"id"`
	Titel             string `json:"titel"`
	TitelKort         string `gorm:"column:titelkort" json:"titelkort"`
	Offentlighedskode string `gorm:"column:offentlighedskode"`
	//Nummer                 string
	//NummerPrefix           string `gorm:"column:nummerprefix"`
	//NummerNumerisk         string `gorm:"column:nummernumerisk"`
	//NummerPostfix          string `gorm:"column:nummerpostfix"`
	Resume string `json:"resume"`
	//Afstemningskonklusion  string `gorm:"column:afstemningskonklusion"`
	//PeriodeId              int
	//AfgorelsesResultatKode string `gorm:"column:afgorelsesresultatkode"`
	//Baggrundsmateriale     string
	//Opdateringsdato        string
	//StatsbudgetSag         int
	//Begrundelse         string
	//Paragrafnummer      int
	//Paragraf            string
	//AfgorelsesDato      string
	//Afgorelse           string
	//RådsmodeDato        string
	//Lovnummer           string
	//LovnummerDato       string
	//Retsinformationsurl string
	//FremsatUnderSagId   int
	//DeltUnderSagId      int

	// Foreign types
	//Type string `json:"sagstype"` //Table Sagtype
	//Kategori string `json:"sagkategori"` //Table Sagkategori
	//Status     string //Table Sagsstatus
	//SagstrinId int    //Table Sagstrin, use-case when identifying sag by sagstrin (relation Afstemning)
}

/*
"id":10,"typeid":7,"kategoriid":null,"statusid":54,"titel":"Forslag til Europa-Parlamentets og R\u00e5dets direktiv om forel\u00f8big retshj\u00e6lp til mist\u00e6nkte eller tiltalte, der frihedsber\u00f8ves, og retshj\u00e6lp i sager ang\u00e5ende europ\u00e6iske arrestordrer","titelkort":"Forslag til Europa-Parlamentets og R\u00e5dets direktiv om forel\u00f8big retshj\u00e6lp til mist\u00e6nkte eller tiltalte, der frihedsber\u00f8ves, og retshj\u00e6lp i sager ang\u00e5ende europ\u00e6iske arrestordrer","offentlighedskode":"O","nummer":"KOM (2013) 0824","nummerprefix":"KOM","nummernumerisk":"824","nummerpostfix":"","resume":"<div>\n\n\n<p>\n<strong>Forslaget er modtaget</strong> af Folketinget i dansk sprogudgave \n den&nbsp;18. december&nbsp;2013, jf. <a href=\"http://www.euo.dk/upload/application/pdf/d83cf832/Protokol2_naerhed.pdf\">Lissabontraktatens \n protokol nr. 2</a> om kontrol med n\u00e6rhedsprincippet. <strong>Fristen \n</strong>p\u00e5  8 uger for fremsendelse af Folketingets begrundede udtalelse  \n<strong>udl\u00f8b&nbsp;if\u00f8lge IPEX&nbsp;12. februar&nbsp;2014.</strong>&nbsp; \n</p>\n\n\n<div>\n\n\n<div>\n\n\n<p>\n<a href=\"/164373.docid\">L\u00e6s</a> \nKommissionens fremsendelsesbrev   \n</p>\n\n</div>\n\n</div>\n\n</div>","afstemningskonklusion":null,"periodeid":14,"afg\u00f8relsesresultatkode":null,"baggrundsmateriale":null,"opdateringsdato":"2016-11-17T13:09:26.873","statsbudgetsag":false,"begrundelse":null,"paragrafnummer":null,"paragraf":null,"afg\u00f8relsesdato":null,"afg\u00f8relse":null,"r\u00e5dsm\u00f8dedato":null,"lovnummer":null,"lovnummerdato":null,"retsinformationsurl":null,"fremsatundersagid":null,"deltundersagid":null
*/
