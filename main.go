package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type DBRecord struct {
	Id            string
	Codbar        string
	Descricao     string
	Categoria     string
	Peso          string
	Unidade       string
	Ncm           string
	Cest          string
	Foto          string
	Observacao    string
	AliqNaced     string
	AliqImpFed    string
	AliqEstatual  string
	AliqMunicipal string
	Gtin          string
	Preco         string
	PrecoMin      string
	PrecoMax      string
	PrecoMedio    string
}

type BlingRecord struct {
	Id                      string `csv:"ID"`
	Codigo                  string `csv:"Código"`
	Descricao               string `csv:"Descrição"`
	Unidade                 string `csv:"Unidade"`
	NCM                     string `csv:"NCM"`
	Origem                  string `csv:"Origem"`
	Preço                   string `csv:"Preco"`
	ValorIPIFix             string `csv:"Valor IPI fixo"`
	Observacoes             string `csv:"Observações"`
	Situacao                string `csv:"Situação"`
	Estoque                 string `csv:"Estoque"`
	PrecoCusto              string `csv:"Preço de custo"`
	CodFornecedor           string `csv:"Cód no fornecedor"`
	Fornecedor              string `csv:"Fornecedor"`
	Localização             string `csv:"Localização"`
	EstoqueMax              string `csv:"Estoque maximo"`
	EstoqueMin              string `csv:"Estoque minimo"`
	PesoLiqKg               string `csv:"Peso líquido (Kg)"`
	PesoBrutoKg             string `csv:"Peso bruto (Kg)"`
	Gtin                    string `csv:"GTIN/EAN"`
	GtinEmbalagem           string `csv:"GTIN/EAN da embalagem"`
	LarguraProd             string `csv:"Largura do Produto"`
	AlturaProduto           string `csv:"Altura do Produto"`
	ProfundProduto          string `csv:"Profundidade do produto"`
	DataValidade            string `csv:"Data Validade"`
	DescricaoProdFornecedor string `csv:"Descrição do Produto no Fornecedor"`
	DescricaoComplementar   string `csv:"Descrição Complementar"`
	UnidadePCaixa           string `csv:"Unidade por Caixa"`
	ProdutoVariacao         string `csv:"Produto Variação"`
	TipoProducao            string `csv:"Tipo Produção"`
	ClasseIPI               string `csv:"Classe de enquadramento do IPI"`
	CodListServico          string `csv:"Código da lista de serviços"`
	TipoItem                string `csv:"Tipo do item"`
	GrupoTags               string `csv:"Grupo de Tags/Tags"`
	Tributos                string `csv:"Tributos"`
	CodigoPai               string `csv:"Código Pai"`
	CodigoIntegracao        string `csv:"Código Integração"`
	GrpProdutos             string `csv:"Grupo de produtos"`
	Marca                   string `csv:"Marca"`
	Cest                    string `csv:"CEST"`
	Volumes                 string `csv:"Volumes"`
	DescrCurta              string `csv:"Descrição Curta"`
	CrossDocking            string `csv:"Cross-Docking"`
	URLImgExter             string `csv:"URL Imagens Externas"`
	LinkExt                 string `csv:"Link Externo"`
	MesesGarantiaFornecedor string `csv:"Meses Garantia no Fornecedor"`
	ClonarDadosPai          string `csv:"Clonar dados do pai"`
	CondicProduto           string `csv:"Condição do produto"`
	FreteGratis             string `csv:"Frete Grátis"`
	NumeroFCI               string `csv:"Número FCI"`
	Video                   string `csv:"Vídeo"`
	Departamento            string `csv:"Departamento"`
	UnidadeMedida           string `csv:"Unidade de medida"`
	PrecoCompra             string `csv:"Preço de compra"`
	ValorICMSRetencao       string `csv:"Valor base ICMS ST para retenção"`
	ValorICMSSTRetencao     string `csv:"Valor ICMS ST para retenção"`
	ValorICMSProprio        string `csv:"Valor ICMS próprio do substituto"`
	CategoriaProd           string `csv:"Categoria do produto"`
	InformacaoAdici         string `csv:"Informações Adicionais"`
}

func createBlingList(data []DBRecord) []BlingRecord {
	var blingList []BlingRecord
	for i := 0; i < len(data); i++ {
		var rec BlingRecord
		rec.Id = data[i].Id
		blingList = append(blingList, rec)

	}
	return blingList

}

func createDBList(data [][]string) []DBRecord {
	var shoppingList []DBRecord
	for i, line := range data {
		if i > 0 { // omit header line
			var rec DBRecord
			for j, field := range line {
				if j == 0 {
					rec.Id = field
				} else if j == 1 {
					rec.Codbar = field
				} else if j == 2 {
					rec.Descricao = field
				} else if j == 3 {
					rec.Categoria = field
				} else if j == 4 {
					rec.Peso = field
				} else if j == 5 {
					rec.Unidade = field
				} else if j == 6 {
					rec.Ncm = field
				} else if j == 7 {
					rec.Cest = field
				} else if j == 8 {
					rec.Foto = field
				} else if j == 9 {
					rec.Observacao = field
				} else if j == 10 {
					rec.AliqNaced = field
				} else if j == 11 {
					rec.AliqImpFed = field
				} else if j == 12 {
					rec.AliqEstatual = field
				} else if j == 13 {
					rec.AliqMunicipal = field
				} else if j == 14 {
					rec.Gtin = field
				} else if j == 15 {
					rec.Preco = field
				} else if j == 16 {
					rec.PrecoMin = field
				} else if j == 17 {
					rec.PrecoMax = field
				} else if j == 18 {
					rec.PrecoMedio = field
				}

			}
			shoppingList = append(shoppingList, rec)
		}
	}
	return shoppingList
}

func main() {
	// open file
	f, err := os.Open("sem_carne_produtos_ref13.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// convert records to array of structs
	DBList := createDBList(data)

	// print the array
	fmt.Printf("total count from DB:%d\n", len(DBList))

	productsBlingFile, err := os.OpenFile("produtos.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer productsBlingFile.Close()

	produtosBling := []*BlingRecord{}

	if err := gocsv.UnmarshalFile(productsBlingFile, &produtosBling); err != nil { // Load clients from file
		panic(err)
	}

	for _, produto := range produtosBling {
		fmt.Printf("Hello %s\n", produto.Descricao)
	}

	if _, err := productsBlingFile.Seek(0, 0); err != nil { // Go to the start of the file
		panic(err)
	}

	for i := 0; i < len(DBList); i++ {
		produtosBling = append(produtosBling,
			&BlingRecord{Id: DBList[i].Id,
				Codigo:         DBList[i].Codbar,
				Descricao:      DBList[i].Descricao,
				DescrCurta:     DBList[i].Descricao,
				Cest:           DBList[i].Cest,
				NCM:            DBList[i].Ncm,
				Unidade:        DBList[i].Unidade,
				URLImgExter:    DBList[i].Foto,
				Gtin:           DBList[i].Gtin,
				Preço:          DBList[i].PrecoMin,
				ClonarDadosPai: "Não",
				Situacao:       "Ativo",
				Estoque:        "0",
				FreteGratis:    "Não"}) // add products from DB to bling list
	}
	csvContent, err := gocsv.MarshalString(&produtosBling) // Get all clients as CSV string
	/*err = gocsv.MarshalFile(&clients, clientsFile) // Use this to save the CSV back to the file
	if err != nil {
		panic(err)
	}
	fmt.Println(csvContent) // Display all clients as CSV string*/

	fw, err := os.Create("sem_carne_bling_format.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer fw.Close()

	_, err2 := fw.WriteString(csvContent)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")

}
