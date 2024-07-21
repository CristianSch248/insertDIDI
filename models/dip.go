package models

type Diploma struct {
	Data struct {
		Cursos []struct {
			Nome     string `json:"Nome"`
			IndNSF   string `json:"IndNSF"`
			Endereco struct {
				UF              string `json:"UF"`
				CEP             string `json:"CEP"`
				Bairro          string `json:"Bairro"`
				Numero          string `json:"Numero"`
				Logradouro      string `json:"Logradouro"`
				NomeMunicipio   string `json:"NomeMunicipio"`
				CodigoMunicipio int    `json:"CodigoMunicipio"`
			} `json:"Endereco"`
			Diplomados []struct {
				ID string `json:"ID"`
				RG struct {
					UF             string `json:"UF"`
					Numero         string `json:"Numero"`
					OrgaoExpedidor string `json:"OrgaoExpedidor"`
				} `json:"RG"`
				CPF      string `json:"CPF"`
				Nome     string `json:"Nome"`
				Sexo     string `json:"Sexo"`
				Filiacao []struct {
					Nome string `json:"Nome"`
					Sexo string `json:"Sexo"`
				} `json:"Filiacao"`
				NumMascara   string `json:"NumMascara"`
				ViaDiploma   int    `json:"ViaDiploma"`
				FormaAcesso  string `json:"FormaAcesso"`
				DataIngresso string `json:"DataIngresso"`
				DtNascimento string `json:"DtNascimento"`
				Naturalidade struct {
					UF              string `json:"UF"`
					Estrangeiro     bool   `json:"Estrangeiro"`
					NomeMunicipio   string `json:"NomeMunicipio"`
					CodigoMunicipio int    `json:"CodigoMunicipio"`
				} `json:"Naturalidade"`
				DadosRegistro struct {
					NroLivro      int    `json:"NroLivro"`
					NumFolha      int    `json:"NumFolha"`
					NumSerie      int    `json:"NumSerie"`
					DtAbertura    string `json:"DtAbertura"`
					NumDiploma    int    `json:"NumDiploma"`
					DtExpedicao   string `json:"DtExpedicao"`
					NumRegistro   int    `json:"NumRegistro"`
					DataConclusao string `json:"DataConclusao"`
					DtColacaoGrau string `json:"DtColacaoGrau"`
					LivroRegistro struct {
						NumLivroReg int `json:"NumLivroReg"`
					} `json:"LivroRegistro"`
					DtRegistroDiploma   string `json:"DtRegistroDiploma"`
					ProcessoDoDiploma   string `json:"ProcessoDoDiploma"`
					ResponsavelRegistro struct {
						CPF  string `json:"CPF"`
						Nome string `json:"Nome"`
					} `json:"ResponsavelRegistro"`
					DocumentacaoComprobatoria []struct {
						Documento        string `json:"Documento"`
						TipoDocumentacao string `json:"TipoDocumentacao"`
					} `json:"DocumentacaoComprobatoria"`
				} `json:"DadosRegistro"`
				DataConclusao  string `json:"DataConclusao"`
				Nacionalidade  string `json:"Nacionalidade"`
				SituacaoAluno  string `json:"SituacaoAluno"`
				SituacoesENADE []struct {
					Edicao   string `json:"Edicao"`
					Motivo   string `json:"Motivo"`
					Condicao string `json:"Condicao"`
					Situacao string `json:"Situacao"`
				} `json:"SituacoesENADE"`
				HistoricoEscolar struct {
					CodigoCurriculo  string `json:"CodigoCurriculo"`
					MatrizCurricular struct {
						DisciplinasCursadas []struct {
							Nota     float64 `json:"Nota"`
							Codigo   string  `json:"Codigo"`
							Status   string  `json:"Status"`
							Periodo  string  `json:"Periodo"`
							Docentes []struct {
								Nome      string `json:"Nome"`
								Titulacao string `json:"Titulacao"`
							} `json:"Docentes"`
							Disciplina   string `json:"Disciplina"`
							StatusExtra  string `json:"StatusExtra"`
							CargaHoraria int    `json:"CargaHoraria"`
						} `json:"DisciplinasCursadas"`
					} `json:"MatrizCurricular"`
				} `json:"HistoricoEscolar"`
				RegistroAcademico              int    `json:"RegistroAcademico"`
				DtEmissaoHistorico             string `json:"DtEmissaoHistorico"`
				HrEmissaoHistorico             string `json:"HrEmissaoHistorico"`
				CargaHorariaCursoIntegralizada int    `json:"CargaHorariaCursoIntegralizada"`
			} `json:"Diplomados"`
			Modalidade  string `json:"Modalidade"`
			Autorizacao struct {
				Data              string `json:"Data"`
				Tipo              string `json:"Tipo"`
				Numero            string `json:"Numero"`
				DataPublicacao    string `json:"DataPublicacao"`
				TramitacaoEMEC    bool   `json:"TramitacaoEMEC"`
				SecaoPublicacao   int    `json:"SecaoPublicacao"`
				VeiculoPublicacao string `json:"VeiculoPublicacao"`
			} `json:"Autorizacao"`
			CargaHoraria int `json:"CargaHoraria"`
			Habilitacoes []struct {
				Data string `json:"Data"`
				Nome string `json:"Nome"`
			} `json:"Habilitacoes"`
			GrauConferido  string `json:"GrauConferido"`
			NomeAbreviado  string `json:"NomeAbreviado"`
			Reconhecimento struct {
				Data              string `json:"Data"`
				Tipo              string `json:"Tipo"`
				Numero            string `json:"Numero"`
				DataPublicacao    string `json:"DataPublicacao"`
				TramitacaoEMEC    bool   `json:"TramitacaoEMEC"`
				SecaoPublicacao   int    `json:"SecaoPublicacao"`
				PaginaPublicacao  int    `json:"PaginaPublicacao"`
				VeiculoPublicacao string `json:"VeiculoPublicacao"`
			} `json:"Reconhecimento"`
			CodigoCursoEmec         int    `json:"CodigoCursoEmec"`
			TituloConferido         string `json:"TituloConferido"`
			RenovacaoReconhecimento struct {
				Data              string `json:"Data"`
				Tipo              string `json:"Tipo"`
				Numero            string `json:"Numero"`
				NumeroDou         int    `json:"NumeroDou"`
				DataPublicacao    string `json:"DataPublicacao"`
				TramitacaoEMEC    bool   `json:"TramitacaoEMEC"`
				SecaoPublicacao   int    `json:"SecaoPublicacao"`
				PaginaPublicacao  int    `json:"PaginaPublicacao"`
				VeiculoPublicacao string `json:"VeiculoPublicacao"`
			} `json:"RenovacaoReconhecimento"`
		} `json:"Cursos"`
	} `json:"data"`
}
