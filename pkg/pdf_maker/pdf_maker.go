package pdf_maker

import (
	"log"

	"github.com/cpprian/content_reader_cli/pkg/content_parser"
	"github.com/signintech/gopdf"
	"github.com/spf13/viper"
)

func CreatePdf(hrefParser *content_parser.ContentContainer) error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	err := pdf.AddTTFFont(viper.GetString("font.style"), "./config/font/"+viper.GetString("font.style")+".ttf")
	if err != nil {
		log.Printf("Cannot add font, error: %v\n", err)
		return err
	}
	err = pdf.SetFont(viper.GetString("font.style"), "", viper.GetInt("font.size"))
	if err != nil {
		log.Printf("Cannot set font, error: %v\n", err)
		return err
	}

	for _, href := range *hrefParser {
		for _, data := range href.Box {
			pdf.Cell(nil, data.Text)
		}
		pdf.Br(30)
	}
	
	return pdf.WritePdf("./store/" + viper.GetString("name") + ".pdf")
}
