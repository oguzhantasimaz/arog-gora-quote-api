package models

import (
	"math/rand"
	"time"
)

type quotes []*quote

type quote struct {
	Text string `json:"quote"`
}

// Get returns a random quote.
func GetRandomQuote() *quote {
	qts := getQuoteList()
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	idx := rand.Intn(len(qts))
	return qts[idx]
}

func getQuoteList() quotes {
	return quotes{
		{Text: "Assikomee"},
		{Text: "Bi cisim yaklaşıyo"},
		{Text: "Kova yok kova, kova olsa"},
		{Text: "Adam sigaradan gitti usta"},
		{Text: "Turist mi çak beş bine çak"},
		{Text: "Her türlü halı kilim travel"},
		{Text: "Gemi biraz sağa mı çekiyor?"},
		{Text: "DÜNYALILARDAN TİSKİNİYORUM!"},
		{Text: "Ceku hazırsan çıkalım balım"},
		{Text: "Robot da olsa insan insandır"},
		{Text: "Hayır bu köyden olsam nolucak?"},
		{Text: "Götünüzden element uydurmayın lan"},
		{Text: "Uzaydayız, herkesin kafası karışık."},
		{Text: "He, she, it loves. Komplikedir yani"},
		{Text: "Toplamayın ulan, düşenlerden yiyin."},
		{Text: "Uzaylıların yanında kavga mı edelim?"},
		{Text: "Benim eğitim sistemimi yargılamayın "},
		{Text: "Karıya sorma karıya daha ne soruyosun"},
		{Text: "Ne zamandır kokluyorsunuz garavel bey?"},
		{Text: "Bu ok zehirli işte atmayın artık bunu ya"},
		{Text: "Abi çok iyi niyetli bi çalışma ama yemez"},
		{Text: "Yoldan geldik, koku seni yanıltmasın hocam"},
		{Text: "Arif bak oğlum robot ne diyo acıktım diyo."},
		{Text: "Bu grup daha genç, ananı tanıyor olamazlar"},
		{Text: "Benim adım Erşan Kuneri! Pornocu muyum ben?"},
		{Text: "Amir toca geri zekalisi 420 yaşında kuna 420"},
		{Text: "Maymun sahnesini at olmaz maymun sahnesi olmaz"},
		{Text: "İçindekini çıkarıcaz, sende olanı sana koyucaz"},
		{Text: "Amerikan baskanı dahil herkesi devreye sokun."},
		{Text: "Uzaylılar tarafından kaçırıldım. Evet tarafından."},
		{Text: "Petrol çıkartmayı öğretiyim de savaş çıkarın di mi"},
		{Text: "Komutan Logar mı? Onun ben amınakoyim. Yak bi tane robot"},
		{Text: "-Güç kalkanları devreye alındı mı? -Heee alındı alındı "},
		{Text: "-Bob Marley Faruk ben +Türk müsün Bob Marley? -Evet. +Yaşa be."},
		{Text: "-Nasıl tak diye burdayım saniyede! +Tamam da niye oradasın?"},
		{Text: "Bekleme süresini uzatırsan zamanda daha da ileriye gidersin"},
		{Text: "-Ateş su toprak tahta -Tahta mı? -Evet tahta zoruna mi gitti?"},
		{Text: "-Oğlum iki çay  kap gel -Tea? +Oh no -Çay söyleme  siktir et."},
		{Text: "Bana ninjalık yaptırmayın tamam mı? BANA NİNJALIK YAPTIRMAYIN!"},
		{Text: "-Robot olum o robot +Hee robot mutfak robotu al mutfağa yanına al."},
		{Text: "Üzümden yapıyorum, burdan böyle damıtıyorum, ihtiyacım kadar yapıyorum "},
		{Text: "Hocam o pelerinli ibneyi biraz dövsek ordan bi kanaat notu verir misiniz?"},
		{Text: "At, yaprak, kelebek, at, yaprak, kelebek gene, yaprak, yaprak, yaprak, yaprak"},
		{Text: "Ya benim uzaylıysa bire bir sevişmiş arkadaşım var ben sana onu da getiricem."},
		{Text: "Herkes türkçe konuşuyor kardeşim. Ben anlamıyorum ki hangisi türk hangisi uzaylı"},
		{Text: "Aşçı bahçıvana,bahçıvan şoföre,şoför uşağa sonra hepsi uşağa, filmin ismi grup indirimi"},
		{Text: "Normal bir erkek için cinsellik birinci plandayken benim için ikinci veyatta üçüncü planda"},
		{Text: "Ben de uzayında değilim ama yapılan muamele yanlış insan uzay deyince daha bir medeniyet bekliyor"},
		{Text: "Malı Arap Faik'ten alıyorduk, Karabük'te 2. yüklemeyi yapıyorduk, Adana'ya kadar da ben kullanıyordum kamyonu!"},
		{Text: "Speaking english? I live in english, it is not a language to me. It is totally the best way of expressing my own."},
		{Text: "Seni seçtim, çünkü sen farklısın. Japon'sun bir kere, akıllı adamsın. Güçlerimizi birleştirirsek kolayca dünyaya dönebiliriz."},
		{Text: "Mesela şöyle örnek vereyim; mesela Brad Pitt tarzı dediğimiz yani daha bi babyface ya da benim tarz; böyle kadınlara hitap eden."},
	}

}
