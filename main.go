package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/joeljunstrom/go-luhn"
	"io/fs"
	"net/http"
	"regexp"
	"strconv"
)

type card struct {
	Number    string
	Valid    bool
	Issuer    string
	MII    string
	PAN    string
}

//go:embed static/dist/*
var static embed.FS

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func api(w http.ResponseWriter, req *http.Request) {
	if req.FormValue("card") == "" {
		fmt.Fprint(w, "error: null")
		return
	}
	card := getCardInfo(req.FormValue("card"))
	js, err := json.Marshal(card)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	setupResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getMII(pan string) string {
	miiDigit, err := strconv.Atoi(string(pan[0]))
	_ = err
	switch miiDigit {
	case 0:
		return "ISO/TC 68"
	case 1, 2:
		return "airlines"
	case 3:
		return "travel and entertainment"
	case 4, 5:
		return "banking and financial"
	case 6:
		return "merchandising and banking/financial"
	case 7:
		return "petroleum"
	case 8:
		return "healthcare and telecommunications"
	case 9:
		return "national assignment"
	default:
		return "no information available"
	}
}

func getIssuer(pan string) string {
	visa, _ := regexp.MatchString("^4\\d{12}(\\d{3})?$", pan)
	mastercard, _ := regexp.MatchString("^(5[1-5]\\d{4}|677189)\\d{10}$|^(222[1-9]|2[3-6]\\d{2}|27[0-1]\\d|2720)\\d{12}$", pan)
	amex, _ := regexp.MatchString("^3[47]\\d{13}$", pan)
	discover, _ := regexp.MatchString("^(6011|65\\d{2})\\d{12}$", pan)
	dankort, _ := regexp.MatchString("^(5019)\\d{12}$", pan)
	jcb, _ := regexp.MatchString("^(?:2131|1800|35\\d{3})\\d{11}$", pan)
	maestro, _ := regexp.MatchString("^(?:5[0678]\\d\\d|6304|6390|67\\d\\d)\\d{8,15}$", pan)
	diners, _ := regexp.MatchString("^3(?:0[0-5]|[68][0-9])[0-9]{11}$", pan)
	if visa {
		return "visa"
	} else if mastercard{
		return "mastercard"
	} else if amex {
		return "amex"
	} else if discover {
		return "discover"
	} else if dankort {
		return "dankort"
	} else if jcb {
		return "jcb"
	} else if maestro {
		return "maestro"
	} else if diners {
		return "diners"
	} else {
		return "unknown"
	}
}


func getPAN(pan string) string {
	if len(pan) > 8 {
		return pan[6 : len(pan)-1]
	} else {
		return "unknown"
	}
}


func getCardInfo(pan string) card {
 card := card{
	 Number: pan,
	 Valid:  luhn.Valid(pan),
	 Issuer: getIssuer(pan),
	 MII:    getMII(pan),
	 PAN:    getPAN(pan),
 }
 return card
}

func serveWeb(port string) {
	color.Cyan("Running on port %s", port)
	embedded, err := fs.Sub(static, "static/dist")
	if err != nil {
		panic(err)
	}

	http.Handle("/",  http.FileServer(http.FS(embedded)))
	http.HandleFunc("/api", api)
	http.ListenAndServe(":" + port, nil)

}

func main() {
	// Print with default helper functions
	color.Cyan("Prints text in cyan.")

	// A newline will be appended automatically
	color.Blue("Prints %s in blue.", "text")

	// These are using the default foreground colors
	color.Red("We have red")
	color.Magenta("And many others ..")
	serveWeb("8090")
}