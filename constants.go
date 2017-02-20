package ovchipapi

const clientId = "nmOIiEJO5khvtLBK9xad3UkkS8Ua"
const clientSecret = "FE8ef6bVBiyN0NeyUJ5VOWdelvQa"

const loginUrl = "https://login.ov-chipkaart.nl/oauth2/token"

const api2BaseUrl = "https://api2.ov-chipkaart.nl/femobilegateway/v1"
const authorizeUrl = api2BaseUrl + "/api/authorize"
const cardsUrl = api2BaseUrl + "/cards/list"
const cardUrl = api2BaseUrl + "/card/"
const transactionsUrl = api2BaseUrl + "/transactions"

type Locale string

const (
	NL_NL Locale = "nl-NL"
	EN_US Locale = "en-US"
)
