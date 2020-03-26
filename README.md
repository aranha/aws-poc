# AWS - Rodando calculadora em golang usando packer, ansible e terraform

## Requisitos

- [Packer](https://www.packer.io/intro/getting-started/install.html)
- [Terraform](https://learn.hashicorp.com/terraform/getting-started/install.html)
- [Ansible](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html)

## Passos

### Adicionando as credenciais
Após os requisitos cumpridos, criar variáveis de ambiente com suas credencias do da AWS, para isso você precisa de sua Acess Key e sua Secret Key, as variáveis deverão ter os seguintes nomes:

- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY

para o Terraform conseguir usá-las crie mais duas:

- TF_VAR_aws_access_key
- TF_VAR_aws_secret_key

Para adicionar é bem simples abra seu arquivo no $HOME/./bashrc e no fim do arquivo coloque:

        export <nome_da_variável>="<valor>"

isso para as quatro.

Após isso execute o arquivo:

        . ./.bashrc


### Criando uma key pair na AWS

Na interface da AWS crie uma key pair, ela deverá ter o nome de **kp_tema13_devops**, após isso ela será baixada no seu computador, vá com o terminal até o local do download e execute o seguinte comando para dar permissões mínimas à key pair

        chmod 400 kp_tema13_devops.pem

### Criando a imagem AMI com Packer e Ansible

Essa etapa é bem simples, para conseguir criar uma imagem AMI e provisioná-la você basta rodar o seguinte comando na pasta raiz do projeto:

        packer build build.json

### Criando a instância da AMI usando Terraform

Após finalizado o processo do tópico anterior você deverá criar e rodar a instância da AMI, para isso use os seguintes comandos:

        terraform init
        terraform apply

caso deseja parar a execução da instância, use:

        terraform destroy

### Rodando a caluladora em go

Após ter criado a instância e ter sua key pair no computador você deverá rodar a calculadora dentro da sua instância na AWS, para isso faça um ssh no ip público disponível na interface sua da AWS, após ter o ip em mãos utilize o seguinte comando na pasta onde está sua key pair gerada anteriormente

        ssh -i kp_tema13_devops.pem ec2-user@<ip_publico>

Após logado no sistema rode os seguintes comandos:

        go build calculator.go
        ./calculator

Para parar o serviço mas não a instância da AMI, use:

        kill -9 <pid_id>

Para descobrir o pid id use:

        netstat -lpnd

## Endpoints

Após feito tudo isso basta somente acessar seu serviço, só funciona em sua rede por questões de segurança criadas no security groups. Os endpoints disponíveis são:

        <ip_publico>:5000/calc/<operação>/n1/n2

as operações disponíveis são: __sum, sub, mult e div__. O histórico está disponível em:

        <ip_publico>:5000/calc/hist