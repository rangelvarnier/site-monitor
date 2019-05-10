**Introdução o Golang**

**Introduçao ao Go**

* Go Workspace
    * A pasta bin para os arquivos binários, compilados
    * A pasta src para o código fonte
    * A pasta pkg para os pacotes compartilhados entre as aplicações
* Instalação da extensão do Go no Visual Studio Code
    * Com isso temos autocomplete, detecção de erros, etc
* Convenções da linguagem
* Implementação do Hello World
* Compilando e executando um programa em Go

**Trabalhando com variáveis**

* Declaração de variáveis
    * Valor padrão das variáveis sem valor inicial
* Inferência de tipos de variáveis
* Descobrir o tipo da variável
    * Através da função TypeOf, do pacote reflect
* Declaração curta de variáveis
* Ler dados digitados do usuário
    * Através das funções Scanf e Scan, do pacote fmt
* Mais convenções do Go
    * Variáveis e imports não utilizados são deletados

**Controlando o fluxo do script**

* Controle de fluxo com if
    * Sua condição não fica entre parênteses e deve sempre retornar um booleano
* Controle de fluxo com switch
    * Se os casos não forem atendidos, será executado o código do caso default
* Introdução às funções
* Pacote os, para encerrar o programa

**Fazendo requisições para a web**

* Pacote net/http, com funcionalidades de acesso à internet, de requisições web
    * Entre elas, a função http.Get, para fazer uma requisição GET para um site
* Funções com múltiplos retornos
    * Identificador em branco, para ignorar um ou mais retornos de uma função
* Status de uma requisição
    * Uma requisição de sucesso possui status code 200
* A instrução for, para deixar o nosso programa em loop eterno

**As principais coleções do Go**

* Coleções: Arrays e Slices
* Trabalhar com Arrays e Slices
* Diferenças entre Arrays e Slices
* A estrutura de repetição for
* Operador de iteração range
    * O range nos dá acesso a cada item da coleção, nos retornando a posição do item iterado e o próprio item daquela posição
* Constantes
* Trabalhando com o pacote time

**Lendo arquivos**

* Detectar erros
* Trabalhar com arquivos
    * Inclusive abrindo e fechando arquivos
* Ler um arquivo por completo com o pacote io/ioutil
* Ler linha a linha com o pacote bufio
    * Criação de leitores
* Limpando strings com TrimSpace

**Escrever em arquivos**

* Abrir arquivos
* Se o arquivo não existir, nós o criamos
* As flags da função os.OpenFile
* Escrever em arquivos
* Converter tipos para string
* Trabalhar com tempo e formatá-lo
* Exibir o conteúdo de um arquivo