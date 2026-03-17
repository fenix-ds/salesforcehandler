package salesforcehandler_test

import (
	"os"
	"testing"

	"github.com/fenix-ds/salesforcehandler"
	"github.com/joho/godotenv"
)

func Test_Get(t *testing.T) {
	err := loadfileEnv(t)
	if err != nil {
		t.Error(err)
	}

	sf, err := salesforcehandler.NewSalesForceHandler(salesforcehandler.SalesForceParam{
		Urls: salesforcehandler.SalesForceUrls{
			Api:           os.Getenv("CRM_API_URL"),
			Autentication: os.Getenv("CRM_LOGIN_URL"),
		},
		Autentication: salesforcehandler.SalesForceAutentication{
			GrantType:    os.Getenv("CRM_LOGIN_FORM_GRANTTYPE"),
			ClientId:     os.Getenv("CRM_LOGIN_FORM_CLIENTID"),
			ClientSecret: os.Getenv("CRM_LOGIN_FORM_CLIENTSECRET"),
			UserName:     os.Getenv("CRM_LOGIN_FORM_USERNAME"),
			Password:     os.Getenv("CRM_LOGIN_FORM_PASSWORD"),
		},
	})

	if err != nil {
		t.Error(err)
	}

	data, err := sf.Get("SELECT+Id,+OwnerId,+IsDeleted,+Name,+CurrencyIsoCode,+RecordTypeId,+CreatedDate,+CreatedById,+LastModifiedDate,+LastModifiedById,+SystemModstamp,+LastActivityDate,+LastViewedDate,+LastReferencedDate,+Aprovado__c,+Hor_rio_de_Inscri_o__c,+Universidade_de_interesse__c,+motivoDesistencia__c,+Status_Primeira_Etapa__c,+Tipo_de_CTA__c,+Interessado_em_Bolsas_Menores__c,+Bolsas_selecionadas__c,+Carreira__c,+Codigo_Processo_Seletivo2__c,+Codigo_unico__c,+Comentario_geral__c,+Conta__c,+Data_Curso_Fim__c,+Data_Curso_inicio__c,+Data_da_entrevista__c,+Data_de_nascimento__c,+Data_inscricao__c,+Email__c,+Email_profissional__c,+Empregador__c,+Endereco__c,+English_Listening__c,+English_Speaking__c,+Formulario_Ativo__c,+Ja_aplicou_anteriormente__c,+Lead__c,+Nivel_de_ingles__c,+Nome_completo__c,+Nome_do_Curso__c,+Nota_geral__c,+Numero__c,+Perfil_do_Linkedin__c,+Por_que_deve_ser_selecionado__c,+Telefone__c,+Universidade__c,+Visto_B1_B2__c,+utm_campaign__c,+utm_content__c,+utm_medium__c,+utm_source__c,+utm_term__c,+Leva__c,+Nome_do_formulario__c,+Curso_de_interesse__c,+Etnia__c,+Fase__c,+Bolsa_desconto__c,+Data_de_Interesse__c,+Pagina_de_Cadastro__c,+Leva_de_Aprova_o__c,+Detalhes_do_Interesse__c,+Link_VA__c,+Score__c,+Subsidio_Corporativo_Empresa__c,+Propriet_rio_do_Lead__c,+Experiencia_internacional__c,+Area_de_trabalho__c,+Descricao_Experiencia_Internacional__c,+Nacionalidade__c,+Total_Subsdio_Corporativo__c,+Nivel_Educacional__c,+Nivel_Profissional__c,+Qual_sua_ocupacao__c,+Years_of_Managing_People__c,+Years_of_Professional_Experience__c,+Hierarchical_Level__c,+Qual_a_sua_atual_posicao_profissional__c,+Genero__c+FROM+Formulario__c+ORDER+BY+Id+ASC+LIMIT+200")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(data)
	}
}
func Test_Update(t *testing.T) {
	err := loadfileEnv(t)
	if err != nil {
		t.Error(err)
	}

	sf, err := salesforcehandler.NewSalesForceHandler(salesforcehandler.SalesForceParam{
		Urls: salesforcehandler.SalesForceUrls{
			Api:           os.Getenv("CRM_API_URL"),
			Autentication: os.Getenv("CRM_LOGIN_URL"),
		},
		Autentication: salesforcehandler.SalesForceAutentication{
			GrantType:    os.Getenv("CRM_LOGIN_FORM_GRANTTYPE"),
			ClientId:     os.Getenv("CRM_LOGIN_FORM_CLIENTID"),
			ClientSecret: os.Getenv("CRM_LOGIN_FORM_CLIENTSECRET"),
			UserName:     os.Getenv("CRM_LOGIN_FORM_USERNAME"),
			Password:     os.Getenv("CRM_LOGIN_FORM_PASSWORD"),
		},
	})

	if err != nil {
		t.Error(err)
	}

	err = sf.Patch(&salesforcehandler.SalesForcePatchObject{
		Name: os.Getenv("TEST_OBJ_NAME"),
		Id:   os.Getenv("TEST_OBJ_ID"),
		Data: struct {
			Name string `json:"Name"`
		}{
			Name: "TEST",
		},
	})
	if err != nil {
		t.Error(err)
	}
}

func Test_DownloadDocumento(t *testing.T) {
	err := loadfileEnv(t)
	if err != nil {
		t.Error(err)
	}

	sf, err := salesforcehandler.NewSalesForceHandler(salesforcehandler.SalesForceParam{
		Urls: salesforcehandler.SalesForceUrls{
			Api:           os.Getenv("CRM_API_URL"),
			Autentication: os.Getenv("CRM_LOGIN_URL"),
		},
		Autentication: salesforcehandler.SalesForceAutentication{
			GrantType:    os.Getenv("CRM_LOGIN_FORM_GRANTTYPE"),
			ClientId:     os.Getenv("CRM_LOGIN_FORM_CLIENTID"),
			ClientSecret: os.Getenv("CRM_LOGIN_FORM_CLIENTSECRET"),
			UserName:     os.Getenv("CRM_LOGIN_FORM_USERNAME"),
			Password:     os.Getenv("CRM_LOGIN_FORM_PASSWORD"),
		},
	})

	if err != nil {
		t.Error(err)
	}

	if document, err := sf.DownloadFile(&salesforcehandler.SalesForceDownloadFilesParam{
		Name: os.Getenv("TEST_OBJ_NAME"),
		Id:   os.Getenv("TEST_OBJ_ID"),
	}); err != nil {
		t.Error(err)
	} else if document == nil {
		t.Error("not contend")
	} else {
		err = os.WriteFile("arquivo.pdf", document, 0644)
		if err != nil {
			panic(err)
		}
	}
}

func loadfileEnv(t *testing.T) error {
	if envs, err := godotenv.Read(".env"); err != nil {
		t.Error(err)
		return err
	} else {
		for key, value := range envs {
			t.Setenv(key, value)
		}
	}

	return nil
}
