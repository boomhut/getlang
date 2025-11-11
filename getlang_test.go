package getlang

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyStringFromReader(t *testing.T) {
	info, _ := FromReader(strings.NewReader(""))
	assert.Equal(t, "und", info.LanguageCode())
}

func TestEnglishPhraseFromBigReader(t *testing.T) {
	largeText := ""
	for i := 0; i < 800; i++ {
		largeText += "this is more language as you can see "
	}
	info, _ := FromReader(strings.NewReader(largeText))
	assert.Equal(t, "en", info.LanguageCode())
	assert.Equal(t, true, info.Confidence() > 0.999)
}

func TestEnglishPhraseFromReader(t *testing.T) {
	info, _ := FromReader(strings.NewReader("this is the language we need to detect"))
	assert.Equal(t, "en", info.LanguageCode())
	assert.Equal(t, true, info.Confidence() > 0.7)
}

func TestEnglishPhraseTag(t *testing.T) {
	info, _ := FromReader(strings.NewReader("this is the language"))
	tag := info.Tag()

	assert.Equal(t, "en", tag.String())
	assert.Equal(t, false, tag.IsRoot())
	assert.Equal(t, true, tag.Parent().IsRoot())
}

func TestEnglishPhraseUSDI(t *testing.T) {
	text := "We hold these truths to be self-evident, that all men are created equal"
	ensureClassifiedWithConfidence(
		t,
		text,
		"en",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"English",
		"English")
}

func TestGermanPhraseUSDI(t *testing.T) {
	text := "Wir halten diese Wahrheiten für ausgemacht, daß alle Menschen gleich erschaffen worden"
	ensureClassifiedWithConfidence(
		t,
		text,
		"de",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"German",
		"Deutsch")
}

// // Test Latin text
// func TestLatinPhrase(t *testing.T) {
// 	// text := "Omnia homines aequales dignitate et iuribus nascuntur"
// 	text := "Omnium humanae gentis partium perspecto et cognito consensum fidemque propriae dignitatis atque iurium, quae omni tempore aequa et paria esse debent nec alienari possunt, totius terrae libertatis iustitiae pacis esse initium"
// 	lang := "Latin"

// 	ensureClassifiedWithConfidence(
// 		t,
// 		text,
// 		"la",
// 		0.95)

// 	ensureClassifiedTextNamed(
// 		t,
// 		text,
// 		"Latin",
// 		lang)
// }

func TestEnglishMixedGerman(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"If you wanted to greet someone in this language, you'd say 'wie geht es'",
		"en",
		0.35)
}

func TestEnglishMixedUkrainian(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"the best thing to say is своїй гідності in my opinon of this.",
		"en",
		0.55)
}

func TestSpanishPhraseUSDI(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Sostenemos como evidentes estas verdades: que los hombres son creados iguales",
		"es",
		0.75)
}

func TestPortuguesePhraseUSDI(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Consideramos estas verdades como autoevidentes, que todos os homens são criados iguais",
		"pt",
		0.95)
}

func TestPolishPhraseUDHR(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Wszyscy ludzie rodzą się wolni i równi w swojej godności i prawach",
		"pl",
		0.95)
}

func TestPunjabiPhrase(t *testing.T) {
	text := "ਮੇਰਾ ਨਾਮ ਭਰਤ ਹੈ."
	lang := "ਪੰਜਾਬੀ"

	ensureClassifiedWithConfidence(
		t,
		text,
		"pa",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Punjabi",
		lang)
}

func TestHungarianPhraseUDHR(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Minden emberi lény szabadon születik és egyenlő méltósága és joga van",
		"hu",
		0.95)
}

func TestItalianPhraseUDHR(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Tutti gli esseri umani nascono liberi ed eguali in dignità e diritti",
		"it",
		0.95)
}

func TestRussianPhraseUDHR(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"В сердце Шепчущего Леса жил гном по имени Бриндл. Его дом был в дупле дуба, наполненном светящимися грибами и баночками со звёздным светом. Каждое утро он чистил свой красный колпак, кормил воробьёв ягодным хлебом и шептал приветствия деревьям, охранявшим его дом. Бриндл отличался от других гномов — он любил исследовать мир за пределами мшистых троп. Однажды туманным утром он нашёл серебряное перо, запутавшееся в паутине. Поняв, что оно принадлежит лунной птице, редкому существу, он спрятал его за ухом и отправился в путь. Через ручьи и поющие корни Бриндл следовал за мерцанием лунного света, пока не добрался до кристального пруда. Там, под звёздным небом, приземлилась лунная птица, её перья сияли, как сны. Она поблагодарила его и оставила слезу света. Эта слеза стала фонарём, который никогда не гас. С той ночи лес мягко светился, помогая каждому путнику найти дорогу домой.",
		"ru",
		0.95)
}

func TestUkrainianPhraseUDHR(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Всі люди народжуються вільними і рівними у своїй гідності та правах ",
		"uk",
		0.7)
}

func TestUkrainianPhrase2UDHR(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"У серці Шепітного Лісу жив гном на ім’я Бріндл. Його дім був у дуплі дуба, наповненому сяючими грибами та баночками зоряного світла. Щоранку він полірував свій червоний ковпак, годував горобців крихтами ягідного хліба й вітав дерева, що стояли на варті його дому. Бріндл був не таким, як інші гноми — він любив подорожувати за межами мохових стежок. Одного туманного ранку він знайшов срібне перо, заплутане в павутині. Зрозумівши, що воно належить місячній пташці, він поклав його за вухо й вирушив у мандрівку.  Через потоки й співаючі корені Бріндл ішов за сяйвом місячного світла, доки не дійшов до кришталевого ставка. Там, під зоряним небом, приземлилася місячна пташка. Вона подякувала йому та залишила одну сльозу світла. Тая сльоза стала ліхтарем, що ніколи не гасне. Відтоді ліс м’яко світився, і кожна заблукла істота знаходила дорогу додому.",
		"uk",
		0.95)
}

func TestFrenchPhraseUDHR(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Tous les êtres humains naissent libres et égaux",
		"fr",
		0.9)
}

func TestKoreanPhrase(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"원래 AB형 사람이 똑똑해",
		"ko",
		0.95)
}

func TestJapanesePhrase(t *testing.T) {
	text := "何を食べますか"
	ensureClassifiedWithConfidence(
		t,
		text,
		"ja",
		0.90)

	ensureClassifiedTextNamed(
		t,
		text,
		"Japanese",
		"日本語")
}

func TestChinesePhrase(t *testing.T) {
	text := "球的采编网络,记者遍布"
	ensureClassifiedWithConfidence(
		t,
		text,
		"zh",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Chinese",
		"中文")
}

func TestArabicPhrase(t *testing.T) {
	text := "اهتمامًا بذلك المشروع. المجموعة الوحيدة التي "
	lang := "العربية"
	ensureClassifiedWithConfidence(
		t,
		text,
		"ar",
		0.9)

	ensureClassifiedTextNamed(
		t,
		text,
		"Arabic",
		lang)
}

func TestBanglaPhrase(t *testing.T) {
	text := "এই গবেষণায় রত, তাঁদেরকে বলা হয় ভাষাবিজ্ঞানী।ভাষাবিজ্ঞানীরা নৈর্ব্যক্তিক"
	lang := "বাংলা"

	ensureClassifiedWithConfidence(
		t,
		text,
		"bn",
		0.85)

	ensureClassifiedTextNamed(
		t,
		text,
		"Bangla",
		lang)
}

func TestHindiPhrase(t *testing.T) {
	text := "ब तक लगातार चल रहा है। इसका प्रसारण प्रत्येक शनिवार और रविवार को रात 10 बजे होता है। इसका पुनः प्रसारण सोनी पल चैनल पर रात 9 बजे होता"
	lang := "हिन्दी"

	ensureClassifiedWithConfidence(
		t,
		text,
		"hi",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Hindi",
		lang)
}

func TestMalayalamPhrase(t *testing.T) {
	text := "എന്റെ പേര് ഭാരത്"
	lang := "മലയാളം"

	ensureClassifiedWithConfidence(
		t,
		text,
		"ml",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Malayalam",
		lang)
}

// Luxembourgish
func TestLuxembourgishPhrase(t *testing.T) {
	text := "All Mënschen sinn gebuer fräi a gläich an der Wierde an den Rechter"
	lang := "Lëtzebuergesch"

	ensureClassifiedWithConfidence(
		t,
		text,
		"lb",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Luxembourgish",
		lang)
}
func TestCatalanPhrase(t *testing.T) {
	text := "Tots els éssers humans neixen lliures i iguals en dignitat i drets"
	lang := "català"

	ensureClassifiedWithConfidence(
		t,
		text,
		"ca",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Catalan",
		lang)
}

func TestGreekPhrase(t *testing.T) {
	text := "Ολοι οι άνθρωποι γεννιούνται ελεύθεροι και ίσοι στην αξιοπρέπεια και στα δικαιώματα"

	ensureClassifiedWithConfidence(
		t,
		text,
		"el",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Greek",
		"Ελληνικά")
}

func TestHebrewPhrase(t *testing.T) {
	text := "כראוי. בִּדקו את כותרת הדף"
	lang := "עברית"

	ensureClassifiedWithConfidence(
		t,
		text,
		"he",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Hebrew",
		lang)
}

func TestGujaratiPhrase(t *testing.T) {
	text := "ગુજરાતી"
	lang := "ગુજરાતી"

	ensureClassifiedWithConfidence(
		t,
		text,
		"gu",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Gujarati",
		lang)
}

func TestThaiPhrase(t *testing.T) {
	text := "ไทย ไทยไทย"
	lang := "ไทย"

	ensureClassifiedWithConfidence(
		t,
		text,
		"th",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Thai",
		lang)
}

func TestArmenianPhrase(t *testing.T) {
	text := "ըստ Գրիգորյան օրացույցի"
	lang := "հայերեն"

	ensureClassifiedWithConfidence(
		t,
		text,
		"hy",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Armenian",
		lang)
}

func TestSerbianLatinPhrase(t *testing.T) {
	text := "ljudi ne znaju jer me uglavnom vide"
	lang := "srpskohrvatski"

	ensureClassifiedWithConfidence(
		t,
		text,
		"sr",
		0.67)

	ensureClassifiedTextNamed(
		t,
		text,
		"Serbo-Croatian",
		lang)
}

func TestSerbianCyrillicPhrase(t *testing.T) {
	text := "Код животиња су ове реакције посебно важне при зарастању рана"
	lang := "српски"

	ensureClassifiedWithConfidence(
		t,
		text,
		"sr",
		0.8)

	ensureClassifiedTextNamed(
		t,
		text,
		"Serbian (Cyrillic)",
		lang)
}

func TestVietnamesePhrase(t *testing.T) {
	text := "Tôi là một người Việt Nam và tôi sống ở Việt Nam"
	lang := "Tiếng Việt"

	ensureClassifiedWithConfidence(
		t,
		text,
		"vi",
		0.4)

	ensureClassifiedTextNamed(
		t,
		text,
		"Vietnamese",
		lang)
}

func TestTeluguPhrase(t *testing.T) {
	text := "భారతదేశంలోని దక్షిణ"
	lang := "తెలుగు"

	ensureClassifiedWithConfidence(
		t,
		text,
		"te",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Telugu",
		lang)
}

func TestTamilPhrase(t *testing.T) {
	text := " நீளமான, கிளைக்காத"
	lang := "தமிழ்"

	ensureClassifiedWithConfidence(
		t,
		text,
		"ta",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Tamil",
		lang)
}

func TestTagalogPhrase(t *testing.T) {
	text := "ano ang nangyayari sa iyo at ang mah-ina mo ay hindi mo"
	lang := "Filipino"

	ensureClassifiedWithConfidence(
		t,
		text,
		"tl",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Filipino",
		lang)
}

func TestDutchPhrase(t *testing.T) {
	text := "Een ieder heeft, waar hij zich ook bevindt, het recht als persoon erkend te worden voor de wet"
	lang := "Nederlands"

	ensureClassifiedWithConfidence(
		t,
		text,
		"nl",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Dutch",
		lang)
}

func TestKannadaPhrase(t *testing.T) {
	text := "ನನ್ನ ಹೆಸರು ಭಾರತ್."
	lang := "ಕನ್ನಡ"

	ensureClassifiedWithConfidence(
		t,
		text,
		"kn",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Kannada",
		lang)
}

func TestNorwegianPhrase(t *testing.T) {
	// text := "Alle mennesker er født frie og med samme menneskeverd og rettigheter"
	// Other test prhases
	text := "Alle mennesker er født frie og med samme menneskeverd og rettigheter. De er utstyrt med fornuft og samvittighet og bør handle mot hverandre i broderlighet."
	lang := "norsk bokmål"

	ensureClassifiedWithConfidence(
		t,
		text,
		"no",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Norwegian Bokmål",
		lang)
}

// Another test for norwegian, but with a different text
func TestNorwegianPhrase2(t *testing.T) {
	text := "Erklæringens offisielle tekst foreligger på FNs seks arbeidsspråk: arabisk, engelsk, fransk, kinesisk, russisk og spansk. En lang rekke av FNs medlemsstater har fulgt Generalforsamlingens oppfordring og oversatt Erklæringen til de nasjonale språk."
	lang := "norsk bokmål"

	ensureClassifiedWithConfidence(
		t,
		text,
		"no",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Norwegian Bokmål",
		lang)
}

// Norwegian Nynorsk
func TestNorwegianNynorskPhrase(t *testing.T) {
	text := "Alle menneske er fødde frie og like i verd og rettar. Dei er utstyrte med fornuft og samvit og bør handle mot kvarandre i brorskap."
	lang := "nynorsk"

	ensureClassifiedWithConfidence(
		t,
		text,
		"nn",
		0.69)

	ensureClassifiedTextNamed(
		t,
		text,
		"Norwegian Nynorsk",
		lang)
}
func TestNorwegianNynorskPhrase2(t *testing.T) {
	text := "Erklæringa sin offisielle tekst ligg føre på FNs seks arbeidsspråk: arabisk, engelsk, fransk, kinesisk, russisk og spansk. Ei lang rekkje av FNs medlemsstatar har følgt Generalforsamlinga si oppmoding og omsett Erklæringa til dei nasjonale språka."
	lang := "nynorsk"

	ensureClassifiedWithConfidence(
		t,
		text,
		"nn",
		0.7)

	ensureClassifiedTextNamed(
		t,
		text,
		"Norwegian Nynorsk",
		lang)
}

// Western Frisian
func TestFrisianPhrase(t *testing.T) {
	text := "Alle minsken binne frij en gelyk yn wearde en rjochten"
	lang := "Frysk"

	ensureClassifiedWithConfidence(
		t,
		text,
		"fy",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Western Frisian",
		lang)
}

func TestSwedishPhrase(t *testing.T) {
	text := "Alla människor är födda fria och lika i värde och rättigheter"
	lang := "svenska"

	ensureClassifiedWithConfidence(
		t,
		text,
		"sv",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Swedish",
		lang)
}

func TestFinnishPhrase(t *testing.T) {
	text := "Kaikki ihmiset syntyvät vapaina ja tasa-arvoisina"
	lang := "suomi"

	ensureClassifiedWithConfidence(
		t,
		text,
		"fi",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Finnish",
		lang)
}

func TestDanishPhrase(t *testing.T) {
	text := "Umiddelbart efter denne historiske begivenhed henstillede generalforsamlingen til alle medlemslande, at de offentliggjorde erklæringens fulde tekst."
	lang := "dansk"

	ensureClassifiedWithConfidence(
		t,
		text,
		"da",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Danish",
		lang)
}

func TestCzechPhrase(t *testing.T) {
	text := "Všichni lidé se rodí svobodní a sobě rovní v důstojnosti a právech"
	lang := "čeština"

	ensureClassifiedWithConfidence(
		t,
		text,
		"cs",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Czech",
		lang)
}

func TestSlovenianPhrase(t *testing.T) {
	text := "Vsi ljudje se rodijo svobodni in enaki v svoji dostojanstvu in pravicah"
	lang := "slovenščina"

	ensureClassifiedWithConfidence(
		t,
		text,
		"sl",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Slovenian",
		lang)
}

func TestSlovakPhrase(t *testing.T) {
	text := "Všetci ľudia sa rodia slobodní a rovní v dôstojnosti a právach"
	lang := "slovenčina"

	ensureClassifiedWithConfidence(
		t,
		text,
		"sk",
		0.92)

	ensureClassifiedTextNamed(
		t,
		text,
		"Slovak",
		lang)
}

func TestAfrikaansPhrase(t *testing.T) {
	text := "Alle mense word vry en gelyk in waardigheid en regte gebore"
	lang := "Afrikaans"

	ensureClassifiedWithConfidence(
		t,
		text,
		"af",
		0.95)

	ensureClassifiedTextNamed(
		t,
		text,
		"Afrikaans",
		lang)
}

func TestNonsense(t *testing.T) {
	text := "wep lvna eeii vl jkk azc nmn iuah ppl zccl c%l aa1z"
	ensureClassifiedWithConfidence(
		t,
		text,
		"und",
		0.75)

	ensureClassifiedTextNamed(
		t,
		text,
		"Unknown language",
		"")
}

func ensureClassifiedWithConfidence(t *testing.T, text string, expectedLang string, minConfidence float64) {
	info := FromString(text)

	assert.Equal(t, expectedLang, info.LanguageCode(), "Misclassified text: "+text)
	assert.Equal(t, true, info.Confidence() > minConfidence)
}

func ensureClassifiedTextNamed(t *testing.T, text string, expectedEnglishName string, expectedSelfName string) {
	info := FromString(text)

	assert.Equal(t, expectedEnglishName, info.LanguageName(), "Wrong language name: "+text)
	assert.Equal(t, expectedSelfName, info.SelfName(), "Wrong self lang name: "+text)
}
