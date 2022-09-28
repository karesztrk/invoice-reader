package main

import (
	"fmt"
    "encoding/json"
	"github.com/ledongthuc/pdf"
)

type Invoice struct {
    InvoiceNo string `json:"invoiceNo"`
    Partner string `json:"partner"`
    Amount string `json:"amount"`
}

func main() {
	pdf.DebugOn = true
	content, err := readPdf("test.pdf") // Read local pdf file
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
	return
}

const InvoiceNoRow = 788
const InvoiceNoCol = 0

const PartnerRow = 699
const PartnerCol = 1

const AmountRow = 490
const AmountCol = 1

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		return "", err
	}
	totalPage := r.NumPage()

    var invoice Invoice

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		rows, _ := p.GetTextByRow()
		for _, row := range rows {
            for j, word := range row.Content {
                //fmt.Println(word.S)
                if (row.Position == InvoiceNoRow && j == InvoiceNoCol) {
                    invoice.InvoiceNo = word.S;
                } else if (row.Position == PartnerRow && j == PartnerCol) {
                    invoice.Partner = word.S;
                } else if (row.Position == AmountRow && j == AmountCol) {
                    invoice.Amount = word.S;
                }
            }
		}
	}

    jsonData, _ := json.Marshal(&invoice)
    fmt.Println(string(jsonData))

	return "", nil
}