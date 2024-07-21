package utilitarios

import (
	"bytes"
	"encoding/json"
	"fmt"
	"insertDiploma/models"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type TokenDidi struct {
	Token       string `json:"token"`
	NivelPerfil string `json:"nivel_perfil"`
}

func GetToken(usuario string, senha string) string {
	tokenDidi := TokenDidi{}

	urlDip := "https://backhomo.astendiploma.com.br/api/v2/login"
	body, _ := json.Marshal(map[string]string{
		"schema": "avmb",
		"login":  usuario,
		"senha":  senha,
	})

	payload := bytes.NewBuffer(body)

	resp, err := http.Post(urlDip, "application/json", payload)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &tokenDidi)
	return string(tokenDidi.Token)
}

func GetCPF() string {
	form := url.Values{}
	form.Add("acao", "gerar_cpf")
	form.Add("pontuacao", "N")

	payload := bytes.NewBufferString(form.Encode())

	resp, err := http.Post("https://www.4devs.com.br/ferramentas_online.php", "application/x-www-form-urlencoded", payload)
	if err != nil {
		log.Fatalln("HTTP - Post - err", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("io.ReadAll - err", err)
	}

	return string(body)
}

func Insert(cpfs []string, token string) {

	jsonFile, err := os.Open(`pacote.json`)
	if err != nil {
		//Caso tenha tido erro, ele é apresentado na tela
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValuesJSON, _ := io.ReadAll(jsonFile)

	var data models.Diploma

	err = json.Unmarshal(byteValuesJSON, &data)
	if err != nil {
		log.Panic(err)
	}

	for i := range cpfs {
		for j := range data.Data.Cursos {
			for k := range data.Data.Cursos[j].Diplomados {
				data.Data.Cursos[j].Diplomados[k].CPF = cpfs[i]
			}
		}

		modifiedJSON, err := json.MarshalIndent(data.Data, "", "  ")
		if err != nil {
			log.Panic(err)
		}

		urlDip := "https://backhomo.astendiploma.com.br/api/v2/diploma/validar"
		body, _ := json.Marshal(map[string]interface{}{
			"data":           json.RawMessage(modifiedJSON),
			"descricao":      "Insert Go" + cpfs[i],
			"id_repositorio": 106,
		})

		payload := bytes.NewBuffer(body)

		req, err := http.NewRequest("POST", urlDip, payload)
		if err != nil {
			log.Fatalln(err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-access-token", token)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(string(responseBody))

		urlDip = "https://backhomo.astendiploma.com.br/api/inserirpacote"
		body, _ = json.Marshal(map[string]interface{}{
			"data":           json.RawMessage(modifiedJSON),
			"descricao":      "Insert Go" + cpfs[i],
			"id_repositorio": 106,
		})

		payload = bytes.NewBuffer(body)

		req, err = http.NewRequest("POST", urlDip, payload)
		if err != nil {
			log.Fatalln(err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-access-token", token)

		client = &http.Client{}
		resp, err = client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		responseBody, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(string(responseBody))
	}
}
