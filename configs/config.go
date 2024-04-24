package configs

import "github.com/spf13/viper"

var cfg *config // cria a variável cfg que será o ponteiro de config, é privada para que apenas
//este package possa acessar as configurações diretamente, mas que sempre que eu quiser acessar
// as configurações eu não consiga alterar nada, eu faço a requisição e este package me retorna
// as informações que foram configuradas neste arquivo.

type config struct {
	API APIConfig // servidor
	DB  DBConfig  // banco de dados
}

type APIConfig struct {
	Port string // porta que será usada para receber as requisições
}

type DBConfig struct { // esta struct vai unir as outras para fazer um unico arquivo mapeando
	//cada dado dentro da sua struc e dentro de seu atributo
	Host     string // Host é o endereço do servidor de banco de dados.
	Port     string // Port é o número da porta do servidor de banco de dados.
	User     string // User é o nome de usuário para autenticação no banco de dados.
	Pass     string // Pass é a senha para autenticação no banco de dados.
	Database string // Database é o nome do banco de dados a ser acessado.
}

func init() { // init sempre é chamada no start das aplicações (faz parte do ciclo de vida do Golang)

	// viper é um pacote que será usado para definir valores padrões para nossas configurações.
	viper.SetDefault("api.port", "9000")           // define a chave api.port, e a porta padrao 9000 (porta onde a API fica escutando por requisições)
	viper.SetDefault("database.host", "localhost") // Define o valor padrão para o host do banco de dados como "localhost"
	viper.SetDefault("database.port", "5432")      // Define o valor padrão para a porta do banco de dados como "5432"
}

func Load() error { // retorna com error caso não consiga carregar as configurações
	viper.SetConfigName("config") // aqui o viper vai procurar pelo arquivo de nome "config"
	viper.SetConfigType("toml")   // Define o tipo de configuração como TOML (Tom's Obvious, Minimal Language)
	viper.AddConfigPath(".")      // Define o diretório atual como o local para procurar arquivos de configuração
	err := viper.ReadInConfig()   // o viper vai ler o arquivo de configuração e guardar o resultado em err
	if err != nil {               // Verifica se o erro é diferente de nil(sem erro) para fazer o if abaixo, se for nill continua depois do return.
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok { //verifica se o erro err não é do tipo viper.ConfigFileNotFoundError.
			//Se não for, o bloco de código dentro do if será executado.
			return err // Retorna o erro.
		}
	}
	cfg = new(config) // cria um ponteiro da struct

	cfg.API = APIConfig{ // configura o campo API da variavel cfg com uma nova atribuição na estrutura APIConfig
		Port: viper.GetString("api.port"), // Obtém a porta do arquivo de configuração usando Viper
	}

	cfg.DB = DBConfig{ // configura o campo DB da variavel cfg com uma novas atribuições na estrutura DBConfig
		Host:     viper.GetString("database.host"), // Obtém o host do arquivo de configuração usando Viper
		Port:     viper.GetString("database.port"), // Obtém a porta do arquivo de configuração usando Viper
		User:     viper.GetString("database.user"), // Obtém o usuário do arquivo de configuração usando Viper
		Pass:     viper.GetString("database.pass"), // Obtém a senha do arquivo de configuração usando Viper
		Database: viper.GetString("database.name"), // Obtém o nome do banco de dados do arquivo de configuração usando Viper
	}

	return nil

}

func GetDb() DBConfig { // retorna a configuração do banco de dados DBConfig
	return cfg.DB
}

func GetServerPort() string { // retorna a porta do servidor armazenada em cfg.API.Port
	return cfg.API.Port
}
