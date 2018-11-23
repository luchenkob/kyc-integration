package common

// KYCStatus2Status maps KYCStatus value to it's API representation.
var KYCStatus2Status = map[KYCStatus]string{
	Error:    "Error",
	Approved: "Approved",
	Denied:   "Denied",
	Unclear:  "Unclear",
}

// KYCFinality2Finality maps KYCFinality value to it's API representation.
var KYCFinality2Finality = map[KYCFinality]string{
	Unknown:  "Unknown",
	Final:    "Final",
	NonFinal: "NonFinal",
}

// CountryAlpha2ToAlpha3 maps country Alpha-2 ISO codes to their Alpha-3 ISO code counterparts.
var CountryAlpha2ToAlpha3 = map[string]string{
	"AF": "AFG",
	"AX": "ALA",
	"AL": "ALB",
	"DZ": "DZA",
	"AS": "ASM",
	"AD": "AND",
	"AO": "AGO",
	"AI": "AIA",
	"AQ": "ATA",
	"AG": "ATG",
	"AR": "ARG",
	"AM": "ARM",
	"AW": "ABW",
	"AU": "AUS",
	"AT": "AUT",
	"AZ": "AZE",
	"BS": "BHS",
	"BH": "BHR",
	"BD": "BGD",
	"BB": "BRB",
	"BY": "BLR",
	"BE": "BEL",
	"BZ": "BLZ",
	"BJ": "BEN",
	"BM": "BMU",
	"BT": "BTN",
	"BO": "BOL",
	"BQ": "BES",
	"BA": "BIH",
	"BW": "BWA",
	"BV": "BVT",
	"BR": "BRA",
	"IO": "IOT",
	"BN": "BRN",
	"BG": "BGR",
	"BF": "BFA",
	"BI": "BDI",
	"CV": "CPV",
	"KH": "KHM",
	"CM": "CMR",
	"CA": "CAN",
	"KY": "CYM",
	"CF": "CAF",
	"TD": "TCD",
	"CL": "CHL",
	"CN": "CHN",
	"CX": "CXR",
	"CC": "CCK",
	"CO": "COL",
	"KM": "COM",
	"CD": "COD",
	"CG": "COG",
	"CK": "COK",
	"CR": "CRI",
	"CI": "CIV",
	"HR": "HRV",
	"CU": "CUB",
	"CW": "CUW",
	"CY": "CYP",
	"CZ": "CZE",
	"DK": "DNK",
	"DJ": "DJI",
	"DM": "DMA",
	"DO": "DOM",
	"EC": "ECU",
	"EG": "EGY",
	"SV": "SLV",
	"GQ": "GNQ",
	"ER": "ERI",
	"EE": "EST",
	"SZ": "SWZ",
	"ET": "ETH",
	"FK": "FLK",
	"FO": "FRO",
	"FJ": "FJI",
	"FI": "FIN",
	"FR": "FRA",
	"GF": "GUF",
	"PF": "PYF",
	"TF": "ATF",
	"GA": "GAB",
	"GM": "GMB",
	"GE": "GEO",
	"DE": "DEU",
	"GH": "GHA",
	"GI": "GIB",
	"GR": "GRC",
	"GL": "GRL",
	"GD": "GRD",
	"GP": "GLP",
	"GU": "GUM",
	"GT": "GTM",
	"GG": "GGY",
	"GN": "GIN",
	"GW": "GNB",
	"GY": "GUY",
	"HT": "HTI",
	"HM": "HMD",
	"VA": "VAT",
	"HN": "HND",
	"HK": "HKG",
	"HU": "HUN",
	"IS": "ISL",
	"IN": "IND",
	"ID": "IDN",
	"IR": "IRN",
	"IQ": "IRQ",
	"IE": "IRL",
	"IM": "IMN",
	"IL": "ISR",
	"IT": "ITA",
	"JM": "JAM",
	"JP": "JPN",
	"JE": "JEY",
	"JO": "JOR",
	"KZ": "KAZ",
	"KE": "KEN",
	"KI": "KIR",
	"KP": "PRK",
	"KR": "KOR",
	"KW": "KWT",
	"KG": "KGZ",
	"LA": "LAO",
	"LV": "LVA",
	"LB": "LBN",
	"LS": "LSO",
	"LR": "LBR",
	"LY": "LBY",
	"LI": "LIE",
	"LT": "LTU",
	"LU": "LUX",
	"MO": "MAC",
	"MK": "MKD",
	"MG": "MDG",
	"MW": "MWI",
	"MY": "MYS",
	"MV": "MDV",
	"ML": "MLI",
	"MT": "MLT",
	"MH": "MHL",
	"MQ": "MTQ",
	"MR": "MRT",
	"MU": "MUS",
	"YT": "MYT",
	"MX": "MEX",
	"FM": "FSM",
	"MD": "MDA",
	"MC": "MCO",
	"MN": "MNG",
	"ME": "MNE",
	"MS": "MSR",
	"MA": "MAR",
	"MZ": "MOZ",
	"MM": "MMR",
	"NA": "NAM",
	"NR": "NRU",
	"NP": "NPL",
	"NL": "NLD",
	"NC": "NCL",
	"NZ": "NZL",
	"NI": "NIC",
	"NE": "NER",
	"NG": "NGA",
	"NU": "NIU",
	"NF": "NFK",
	"MP": "MNP",
	"NO": "NOR",
	"OM": "OMN",
	"PK": "PAK",
	"PW": "PLW",
	"PS": "PSE",
	"PA": "PAN",
	"PG": "PNG",
	"PY": "PRY",
	"PE": "PER",
	"PH": "PHL",
	"PN": "PCN",
	"PL": "POL",
	"PT": "PRT",
	"PR": "PRI",
	"QA": "QAT",
	"RE": "REU",
	"RO": "ROU",
	"RU": "RUS",
	"RW": "RWA",
	"BL": "BLM",
	"SH": "SHN",
	"KN": "KNA",
	"LC": "LCA",
	"MF": "MAF",
	"PM": "SPM",
	"VC": "VCT",
	"WS": "WSM",
	"SM": "SMR",
	"ST": "STP",
	"SA": "SAU",
	"SN": "SEN",
	"RS": "SRB",
	"SC": "SYC",
	"SL": "SLE",
	"SG": "SGP",
	"SX": "SXM",
	"SK": "SVK",
	"SI": "SVN",
	"SB": "SLB",
	"SO": "SOM",
	"ZA": "ZAF",
	"GS": "SGS",
	"SS": "SSD",
	"ES": "ESP",
	"LK": "LKA",
	"SD": "SDN",
	"SR": "SUR",
	"SJ": "SJM",
	"WE": "SWE",
	"CH": "CHE",
	"SY": "SYR",
	"TW": "TWN",
	"TJ": "TJK",
	"TZ": "TZA",
	"TH": "THA",
	"TL": "TLS",
	"TG": "TGO",
	"TK": "TKL",
	"TO": "TON",
	"TT": "TTO",
	"TN": "TUN",
	"TR": "TUR",
	"TM": "TKM",
	"TC": "TCA",
	"TV": "TUV",
	"UG": "UGA",
	"UA": "UKR",
	"AE": "ARE",
	"GB": "GBR",
	"UM": "UMI",
	"US": "USA",
	"UY": "URY",
	"UZ": "UZB",
	"VU": "VUT",
	"VE": "VEN",
	"VN": "VNM",
	"VG": "VGB",
	"VI": "VIR",
	"WF": "WLF",
	"EH": "ESH",
	"TE": "YEM",
	"ZM": "ZMB",
	"ZW": "ZWE",
	"XK": "XKX",
}

// CountryAlpha2ToName maps country Alpha-2 ISO codes to their names.
var CountryAlpha2ToName = map[string]string{
	"AF": "Afghanistan",
	"AX": "Aland Islands",
	"AL": "Albania",
	"DZ": "Algeria",
	"AS": "American Samoa",
	"AD": "Andorra",
	"AO": "Angola",
	"AI": "Anguilla",
	"AQ": "Antarctica",
	"AG": "Antigua and Barbuda",
	"AR": "Argentina",
	"AM": "Armenia",
	"AW": "Aruba",
	"AU": "Australia",
	"AT": "Austria",
	"AZ": "Azerbaijan",
	"BS": "Bahamas",
	"BH": "Bahrain",
	"BD": "Bangladesh",
	"BB": "Barbados",
	"BY": "Belarus",
	"BE": "Belgium",
	"BZ": "Belize",
	"BJ": "Benin",
	"BM": "Bermuda",
	"BT": "Bhutan",
	"BO": "Bolivia",
	"BQ": "Bonaire",
	"BA": "Bosnia and Herzegovina",
	"BW": "Botswana",
	"BV": "Bouvet Island",
	"BR": "Brazil",
	"IO": "British Indian Ocean Territory",
	"BN": "Brunei Darussalam",
	"BG": "Bulgaria",
	"BF": "Burkina Faso",
	"BI": "Burundi",
	"CV": "Cabo Verde",
	"KH": "Cambodia",
	"CM": "Cameroon",
	"CA": "Canada",
	"KY": "Cayman Islands",
	"CF": "Central African Republic",
	"TD": "Chad",
	"CL": "Chile",
	"CN": "China",
	"CX": "Christmas Island",
	"CC": "Cocos (Keeling) Islands",
	"CO": "Colombia",
	"KM": "Comoros",
	"CD": "Democratic Republic of Congo",
	"CG": "Congo",
	"CK": "Cook Islands",
	"CR": "Costa Rica",
	"CI": "Côte d'Ivoire",
	"HR": "Croatia",
	"CU": "Cuba",
	"CW": "Curaçao",
	"CY": "Cyprus",
	"CZ": "Czech Republic ",
	"DK": "Denmark",
	"DJ": "Djibouti",
	"DM": "Dominica",
	"DO": "Dominican Republic",
	"EC": "Ecuador",
	"EG": "Egypt",
	"SV": "El Salvador",
	"GQ": "Equatorial Guinea",
	"ER": "Eritrea",
	"EE": "Estonia",
	"SZ": "Eswatini",
	"ET": "Ethiopia",
	"FK": "Falkland Islands",
	"FO": "Faroe Islands",
	"FJ": "Fiji",
	"FI": "Finland",
	"FR": "France",
	"GF": "French Guiana",
	"PF": "French Polynesia",
	"TF": "French Southern Territories",
	"GA": "Gabon",
	"GM": "Gambia",
	"GE": "Georgia",
	"DE": "Germany",
	"GH": "Ghana",
	"GI": "Gibraltar",
	"GR": "Greece",
	"GL": "Greenland",
	"GD": "Grenada",
	"GP": "Guadeloupe",
	"GU": "Guam",
	"GT": "Guatemala",
	"GG": "Guernsey",
	"GN": "Guinea",
	"GW": "Guinea-Bissau",
	"GY": "Guyana",
	"HT": "Haiti",
	"HM": "Heard Island and McDonald Islands",
	"VA": "Vatican",
	"HN": "Honduras",
	"HK": "Hong Kong",
	"HU": "Hungary",
	"IS": "Iceland",
	"IN": "India",
	"ID": "Indonesia",
	"IR": "Iran",
	"IQ": "Iraq",
	"IE": "Ireland",
	"IM": "Isle of Man",
	"IL": "Israel",
	"IT": "Italy",
	"JM": "Jamaica",
	"JP": "Japan",
	"JE": "Jersey",
	"JO": "Jordan",
	"KZ": "Kazakhstan",
	"KE": "Kenya",
	"KI": "Kiribati",
	"KP": "Democratic People's Republic of Korea",
	"KR": "Republic of Korea",
	"KW": "Kuwait",
	"KG": "Kyrgyzstan",
	"LA": "Lao People's Democratic Republic",
	"LV": "Latvia",
	"LB": "Lebanon",
	"LS": "Lesotho",
	"LR": "Liberia",
	"LY": "Libya",
	"LI": "Liechtenstein",
	"LT": "Lithuania",
	"LU": "Luxembourg",
	"MO": "Macao",
	"MK": "Macedonia",
	"MG": "Madagascar",
	"MW": "Malawi",
	"MY": "Malaysia",
	"MV": "Maldives",
	"ML": "Mali",
	"MT": "Malta",
	"MH": "Marshall Islands",
	"MQ": "Martinique",
	"MR": "Mauritania",
	"MU": "Mauritius",
	"YT": "Mayotte",
	"MX": "Mexico",
	"FM": "Micronesia",
	"MD": "Moldova",
	"MC": "Monaco",
	"MN": "Mongolia",
	"ME": "Montenegro",
	"MS": "Montserrat",
	"MA": "Morocco",
	"MZ": "Mozambique",
	"MM": "Myanmar",
	"NA": "Namibia",
	"NR": "Nauru",
	"NP": "Nepal",
	"NL": "Netherlands",
	"NC": "New Caledonia",
	"NZ": "New Zealand",
	"NI": "Nicaragua",
	"NE": "Niger",
	"NG": "Nigeria",
	"NU": "Niue",
	"NF": "Norfolk Island",
	"MP": "Northern Mariana Islands",
	"NO": "Norway",
	"OM": "Oman",
	"PK": "Pakistan",
	"PW": "Palau",
	"PS": "State of Palestine",
	"PA": "Panama",
	"PG": "Papua New Guinea",
	"PY": "Paraguay",
	"PE": "Peru",
	"PH": "Philippines",
	"PN": "Pitcairn",
	"PL": "Poland",
	"PT": "Portugal",
	"PR": "Puerto Rico",
	"QA": "Qatar",
	"RE": "Réunion",
	"RO": "Romania",
	"RU": "Russian Federation",
	"RW": "Rwanda",
	"BL": "Saint Barthélemy",
	"SH": "Saint Helena, Ascension and Tristan da Cunha",
	"KN": "Saint Kitts and Nevis",
	"LC": "Saint Lucia",
	"MF": "Saint Martin",
	"PM": "Saint Pierre and Miquelon",
	"VC": "Saint Vincent and the Grenadines",
	"WS": "Samoa",
	"SM": "San Marino",
	"ST": "Sao Tome and Principe",
	"SA": "Saudi Arabia",
	"SN": "Senegal",
	"RS": "Serbia",
	"SC": "Seychelles",
	"SL": "Sierra Leone",
	"SG": "Singapore",
	"SX": "Sint Maarten",
	"SK": "Slovakia",
	"SI": "Slovenia",
	"SB": "Solomon Islands",
	"SO": "Somalia",
	"ZA": "South Africa",
	"GS": "South Georgia and the South Sandwich Islands",
	"SS": "South Sudan",
	"ES": "Spain",
	"LK": "Sri Lanka",
	"SD": "Sudan",
	"SR": "Suriname",
	"SJ": "Svalbard and Jan Mayen",
	"WE": "Sweden",
	"CH": "Switzerland",
	"SY": "Syrian Arab Republic",
	"TW": "Taiwan",
	"TJ": "Tajikistan",
	"TZ": "Tanzania",
	"TH": "Thailand",
	"TL": "Timor-Leste",
	"TG": "Togo",
	"TK": "Tokelau",
	"TO": "Tonga",
	"TT": "Trinidad and Tobago",
	"TN": "Tunisia",
	"TR": "Turkey",
	"TM": "Turkmenistan",
	"TC": "Turks and Caicos Islands",
	"TV": "Tuvalu",
	"UG": "Uganda",
	"UA": "Ukraine",
	"AE": "United Arab Emirates",
	"GB": "United Kingdom of Great Britain and Northern Ireland",
	"UM": "United States Minor Outlying Islands",
	"US": "United States of America",
	"UY": "Uruguay",
	"UZ": "Uzbekistan",
	"VU": "Vanuatu",
	"VE": "Venezuela",
	"VN": "Viet Nam",
	"VG": "Virgin Islands (British)",
	"VI": "Virgin Islands (U.S.)",
	"WF": "Wallis and Futuna",
	"EH": "Western Sahara",
	"TE": "Yemen",
	"ZM": "Zambia",
	"ZW": "Zimbabwe",
	"XK": "Kosovo",
}

// CountryName2ToAlpha3 maps ISO 3166-1 country names (including some common unofficial names) to their corresponding ISO 3166-1 alpha-3 codes.
var CountryName2ToAlpha3 = map[string]string{
	"AFGHANISTAN":            "AFG",
	"ALAND ISLANDS":          "ALA",
	"ALBANIA":                "ALB",
	"ALGERIA":                "DZA",
	"AMERICAN SAMOA":         "ASM",
	"ANDORRA":                "AND",
	"ANGOLA":                 "AGO",
	"ANGUILLA":               "AIA",
	"ANTARCTICA":             "ATA",
	"ANTIGUA AND BARBUDA":    "ATG",
	"ARGENTINA":              "ARG",
	"ARMENIA":                "ARM",
	"ARUBA":                  "ABW",
	"AUSTRALIA":              "AUS",
	"AUSTRIA":                "AUT",
	"AZERBAIJAN":             "AZE",
	"BAHAMAS":                "BHS",
	"BAHRAIN":                "BHR",
	"BANGLADESH":             "BGD",
	"BARBADOS":               "BRB",
	"BELARUS":                "BLR",
	"BELGIUM":                "BEL",
	"BELIZE":                 "BLZ",
	"BENIN":                  "BEN",
	"BERMUDA":                "BMU",
	"BHUTAN":                 "BTN",
	"BOLIVIA":                "BOL",
	"BONAIRE":                "BES",
	"BOSNIA AND HERZEGOVINA": "BIH",
	"BOTSWANA":               "BWA",
	"BOUVET ISLAND":          "BVT",
	"BRAZIL":                 "BRA",
	"BRITISH INDIAN OCEAN TERRITORY": "IOT",
	"BRUNEI DARUSSALAM":              "BRN",
	"BULGARIA":                       "BGR",
	"BURKINA FASO":                   "BFA",
	"BURUNDI":                        "BDI",
	"CABO VERDE":                     "CPV",
	"CAMBODIA":                       "KHM",
	"CAMEROON":                       "CMR",
	"CANADA":                         "CAN",
	"CAYMAN ISLANDS":                 "CYM",
	"CENTRAL AFRICAN REPUBLIC":       "CAF",
	"CHAD":                         "TCD",
	"CHILE":                        "CHL",
	"CHINA":                        "CHN",
	"CHRISTMAS ISLAND":             "CXR",
	"COCOS (KEELING) ISLANDS":      "CCK",
	"COLOMBIA":                     "COL",
	"COMOROS":                      "COM",
	"DEMOCRATIC REPUBLIC OF CONGO": "COD",
	"CONGO":                       "COG",
	"COOK ISLANDS":                "COK",
	"COSTA RICA":                  "CRI",
	"CÔTE D'IVOIRE":               "CIV",
	"CROATIA":                     "HRV",
	"CUBA":                        "CUB",
	"CURAÇAO":                     "CUW",
	"CYPRUS":                      "CYP",
	"CZECH REPUBLIC ":             "CZE",
	"DENMARK":                     "DNK",
	"DJIBOUTI":                    "DJI",
	"DOMINICA":                    "DMA",
	"DOMINICAN REPUBLIC":          "DOM",
	"ECUADOR":                     "ECU",
	"EGYPT":                       "EGY",
	"EL SALVADOR":                 "SLV",
	"EQUATORIAL GUINEA":           "GNQ",
	"ERITREA":                     "ERI",
	"ESTONIA":                     "EST",
	"ESWATINI":                    "SWZ",
	"ETHIOPIA":                    "ETH",
	"FALKLAND ISLANDS":            "FLK",
	"FAROE ISLANDS":               "FRO",
	"FIJI":                        "FJI",
	"FINLAND":                     "FIN",
	"FRANCE":                      "FRA",
	"FRENCH GUIANA":               "GUF",
	"FRENCH POLYNESIA":            "PYF",
	"FRENCH SOUTHERN TERRITORIES": "ATF",
	"GABON":         "GAB",
	"GAMBIA":        "GMB",
	"GEORGIA":       "GEO",
	"GERMANY":       "DEU",
	"GHANA":         "GHA",
	"GIBRALTAR":     "GIB",
	"GREECE":        "GRC",
	"GREENLAND":     "GRL",
	"GRENADA":       "GRD",
	"GUADELOUPE":    "GLP",
	"GUAM":          "GUM",
	"GUATEMALA":     "GTM",
	"GUERNSEY":      "GGY",
	"GUINEA":        "GIN",
	"GUINEA-BISSAU": "GNB",
	"GUYANA":        "GUY",
	"HAITI":         "HTI",
	"HEARD ISLAND AND MCDONALD ISLANDS":     "HMD",
	"VATICAN":                               "VAT",
	"HONDURAS":                              "HND",
	"HONG KONG":                             "HKG",
	"HUNGARY":                               "HUN",
	"ICELAND":                               "ISL",
	"INDIA":                                 "IND",
	"INDONESIA":                             "IDN",
	"IRAN":                                  "IRN",
	"IRAQ":                                  "IRQ",
	"IRELAND":                               "IRL",
	"ISLE OF MAN":                           "IMN",
	"ISRAEL":                                "ISR",
	"ITALY":                                 "ITA",
	"JAMAICA":                               "JAM",
	"JAPAN":                                 "JPN",
	"JERSEY":                                "JEY",
	"JORDAN":                                "JOR",
	"KAZAKHSTAN":                            "KAZ",
	"KENYA":                                 "KEN",
	"KIRIBATI":                              "KIR",
	"DEMOCRATIC PEOPLE'S REPUBLIC OF KOREA": "PRK",
	"REPUBLIC OF KOREA":                     "KOR",
	"KUWAIT":                                "KWT",
	"KYRGYZSTAN":                            "KGZ",
	"LAO PEOPLE'S DEMOCRATIC REPUBLIC":      "LAO",
	"LATVIA":                                       "LVA",
	"LEBANON":                                      "LBN",
	"LESOTHO":                                      "LSO",
	"LIBERIA":                                      "LBR",
	"LIBYA":                                        "LBY",
	"LIECHTENSTEIN":                                "LIE",
	"LITHUANIA":                                    "LTU",
	"LUXEMBOURG":                                   "LUX",
	"MACAO":                                        "MAC",
	"MACEDONIA":                                    "MKD",
	"MADAGASCAR":                                   "MDG",
	"MALAWI":                                       "MWI",
	"MALAYSIA":                                     "MYS",
	"MALDIVES":                                     "MDV",
	"MALI":                                         "MLI",
	"MALTA":                                        "MLT",
	"MARSHALL ISLANDS":                             "MHL",
	"MARTINIQUE":                                   "MTQ",
	"MAURITANIA":                                   "MRT",
	"MAURITIUS":                                    "MUS",
	"MAYOTTE":                                      "MYT",
	"MEXICO":                                       "MEX",
	"MICRONESIA":                                   "FSM",
	"MOLDOVA":                                      "MDA",
	"MONACO":                                       "MCO",
	"MONGOLIA":                                     "MNG",
	"MONTENEGRO":                                   "MNE",
	"MONTSERRAT":                                   "MSR",
	"MOROCCO":                                      "MAR",
	"MOZAMBIQUE":                                   "MOZ",
	"MYANMAR":                                      "MMR",
	"NAMIBIA":                                      "NAM",
	"NAURU":                                        "NRU",
	"NEPAL":                                        "NPL",
	"NETHERLANDS":                                  "NLD",
	"NEW CALEDONIA":                                "NCL",
	"NEW ZEALAND":                                  "NZL",
	"NICARAGUA":                                    "NIC",
	"NIGER":                                        "NER",
	"NIGERIA":                                      "NGA",
	"NIUE":                                         "NIU",
	"NORFOLK ISLAND":                               "NFK",
	"NORTHERN MARIANA ISLANDS":                     "MNP",
	"NORWAY":                                       "NOR",
	"OMAN":                                         "OMN",
	"PAKISTAN":                                     "PAK",
	"PALAU":                                        "PLW",
	"STATE OF PALESTINE":                           "PSE",
	"PANAMA":                                       "PAN",
	"PAPUA NEW GUINEA":                             "PNG",
	"PARAGUAY":                                     "PRY",
	"PERU":                                         "PER",
	"PHILIPPINES":                                  "PHL",
	"PITCAIRN":                                     "PCN",
	"POLAND":                                       "POL",
	"PORTUGAL":                                     "PRT",
	"PUERTO RICO":                                  "PRI",
	"QATAR":                                        "QAT",
	"RÉUNION":                                      "REU",
	"ROMANIA":                                      "ROU",
	"RUSSIAN FEDERATION":                           "RUS",
	"RUSSIA":                                       "RUS",
	"RWANDA":                                       "RWA",
	"SAINT BARTHÉLEMY":                             "BLM",
	"SAINT HELENA, ASCENSION AND TRISTAN DA CUNHA": "SHN",
	"SAINT KITTS AND NEVIS":                        "KNA",
	"SAINT LUCIA":                                  "LCA",
	"SAINT MARTIN":                                 "MAF",
	"SAINT PIERRE AND MIQUELON":                    "SPM",
	"SAINT VINCENT AND THE GRENADINES":             "VCT",
	"SAMOA":                                                "WSM",
	"SAN MARINO":                                           "SMR",
	"SAO TOME AND PRINCIPE":                                "STP",
	"SAUDI ARABIA":                                         "SAU",
	"SENEGAL":                                              "SEN",
	"SERBIA":                                               "SRB",
	"SEYCHELLES":                                           "SYC",
	"SIERRA LEONE":                                         "SLE",
	"SINGAPORE":                                            "SGP",
	"SINT MAARTEN":                                         "SXM",
	"SLOVAKIA":                                             "SVK",
	"SLOVENIA":                                             "SVN",
	"SOLOMON ISLANDS":                                      "SLB",
	"SOMALIA":                                              "SOM",
	"SOUTH AFRICA":                                         "ZAF",
	"SOUTH GEORGIA AND THE SOUTH SANDWICH ISLANDS":         "SGS",
	"SOUTH SUDAN":                                          "SSD",
	"SPAIN":                                                "ESP",
	"SRI LANKA":                                            "LKA",
	"SUDAN":                                                "SDN",
	"SURINAME":                                             "SUR",
	"SVALBARD AND JAN MAYEN":                               "SJM",
	"SWEDEN":                                               "SWE",
	"SWITZERLAND":                                          "CHE",
	"SYRIAN ARAB REPUBLIC":                                 "SYR",
	"SYRIA":                                                "SYR",
	"TAIWAN":                                               "TWN",
	"TAJIKISTAN":                                           "TJK",
	"TANZANIA":                                             "TZA",
	"THAILAND":                                             "THA",
	"TIMOR-LESTE":                                          "TLS",
	"TOGO":                                                 "TGO",
	"TOKELAU":                                              "TKL",
	"TONGA":                                                "TON",
	"TRINIDAD AND TOBAGO":                                  "TTO",
	"TUNISIA":                                              "TUN",
	"TURKEY":                                               "TUR",
	"TURKMENISTAN":                                         "TKM",
	"TURKS AND CAICOS ISLANDS":                             "TCA",
	"TUVALU":                                               "TUV",
	"UGANDA":                                               "UGA",
	"UKRAINE":                                              "UKR",
	"UNITED ARAB EMIRATES":                                 "ARE",
	"UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND": "GBR",
	"UNITED STATES MINOR OUTLYING ISLANDS":                 "UMI",
	"UNITED STATES OF AMERICA":                             "USA",
	"USA":                      "USA",
	"URUGUAY":                  "URY",
	"UZBEKISTAN":               "UZB",
	"VANUATU":                  "VUT",
	"VENEZUELA":                "VEN",
	"VIET NAM":                 "VNM",
	"VIRGIN ISLANDS (BRITISH)": "VGB",
	"VIRGIN ISLANDS (U.S.)":    "VIR",
	"WALLIS AND FUTUNA":        "WLF",
	"WESTERN SAHARA":           "ESH",
	"YEMEN":                    "YEM",
	"ZAMBIA":                   "ZMB",
	"ZIMBABWE":                 "ZWE",
	"KOSOVO":                   "XKX",
}
