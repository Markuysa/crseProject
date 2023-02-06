package cbgateway

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"ozonProjectmodule/internal/model/domain"
	"strings"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/text/encoding/charmap"
)

type Gateway struct {
	gtw *http.Client
}

func (g *Gateway) FetchRates(ctx context.Context, date time.Time) ([]domain.Rate, error) {

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	url := fmt.Sprintf("https://www.cbr-xml-daily.ru/daily_eng_utf8.xml?date_req=%s", date.Format("02/01/2006"))

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "request with context")
	}
	response, err := g.gtw.Do(request)

	if response.StatusCode != http.StatusOK {
		return nil, errors.Wrap(err, "request with context")
	}

	defer response.Body.Close()

	d := xml.NewDecoder(response.Body)
	//hz chto za fignya
	d.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("unknown charset: %s", charset)
		}
	}

	var cbRates Rates

	if err := d.Decode(&cbRates); err != nil {
		return nil, errors.Wrap(err, "decode the rates")
	}

	rates := make([]domain.Rate, len(cbRates.Currencies))
	for _, rate := range cbRates.Currencies{
		rates = append(rates, domain.Rate{
			Code:     rate.CharCode,
			Original: strings.Replace(rate.Value, ",", ".", 1),
			Nominal:  rate.Nominal,
		})
	}

	return rates, nil
}
