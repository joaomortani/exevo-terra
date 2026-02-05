# Exevo Terra 🪄

<p align="center">
  <img src="https://img.shields.io/github/go-mod/go-version/joaomortani/exevo-terra?style=for-the-badge&logo=go&color=00ADD8" alt="Go Version" />
  <img src="https://img.shields.io/badge/Terraform-%3E%3D1.5-623CE4?style=for-the-badge&logo=terraform" alt="Terraform Version" />
  <img src="https://img.shields.io/github/license/joaomortani/exevo-terra?style=for-the-badge&color=blue" alt="License" />
  <img src="https://img.shields.io/badge/status-beta-orange?style=for-the-badge" alt="Status" />
</p>

> **"Exevo Terra"**: Do latim tibiano *"Invocar Terra"*.
> Traga sua infraestrutura legada da AWS para o mundo do Código (HCL) instantaneamente.

O **Exevo Terra** é uma CLI escrita em Go projetada para engenheiros de SRE e DevOps que precisam importar recursos existentes da AWS para o Terraform sem escrever HCL manualmente e sem sofrer com `terraform import` linha por linha.

---

## 🚀 Features (v0.3)

- **Motor Genérico & Configurável**: Arquitetura 100% guiada por arquivo de configuração (`exevo.yaml`). Adicione novos recursos ou altere mapeamentos sem tocar no código Go.
- **Suporte Multi-Recurso**: Suporte nativo implementado para **RDS** e **S3**, com arquitetura pronta para escalar para EC2, Lambda, etc.
- **Discovery & Import Automático**: Varre sua conta AWS, encontra recursos e gera blocos `import { ... }` (Terraform 1.5+), eliminando a necessidade de `terraform import` manual.
- **Infraestrutura "Production-Ready"**:
  - Gera `main.tf` limpo e formatado (HCL).
  - Gera `versions.tf` com configuração de **Backend S3 Dinâmico** e versionamento de providers.
  - Gera `imports.tf` para bind imediato do estado.
- **Developer Experience (DX)**:
  - `init`: Cria o scaffolding do projeto com templates prontos.
  - `inspect`: Varre um recurso real na nuvem e gera uma tabela Markdown com todos os campos disponíveis para mapeamento (Auto-Discovery de Schema).
- **Isolamento de Estado**: O código é gerado em pastas isoladas (`infra/rds`, `infra/s3`), garantindo que o `terraform.tfstate` não vire um monólito.
- **SSO Nativo**: Integração transparente com credenciais `aws sso`.

# 🚀 Instalação

### Opção 1: Via Go (Para Desenvolvedores)
Se você já tem o Go instalado:

```bash
go install [github.com/joaomortani/exevo-terra@latest] (https://github.com/joaomortani/exevo-terra@latest)
```

### Opção 2: Binário (Para todos)
1. Vá na aba [Releases](../../releases) deste repositório.
2. Baixe a versão compatível com seu sistema (Ex: `Linux_x86_64` ou `Darwin_arm64` para Mac M1/M2).
3. Descompacte e mova para o seu path:
   ```bash
   tar -xvf exevo-terra_*.tar.gz
   sudo mv exevo-terra /usr/local/bin/
   ```

## ⚡ Como Usar (Quickstart)

1. **Inicialize o projeto:**
   ```bash
   exevo-terra init
   ```

2. **Edite o arquivo gerado:**
   Abra o `exevo.yaml` e ajuste o nome do bucket e as configurações do S3/RDS.

3. **Descubra os campos (Opcional):**
   Descubra quais campos da AWS você pode mapear no seu YAML:
   ```bash
   exevo-terra inspect --resource rds
   ```

4. **Gere o código:**
   Conecte na AWS e gere os arquivos Terraform:
   ```bash
   exevo-terra generate --resource rds --profile default
   ```

## 🗺️ Roadmap & Futuro

> O Exevo Terra evoluiu de um script simples para um **Framework de IaC Multi-Cloud** orientado a configuração.

- [x] **v0.1:** MVP focado em RDS (Hardcoded).
- [x] **v0.2:** Motor Genérico "Bring Your Own Module" (BYOM) e Suporte a S3.
- [x] **v0.3 (Atual):** Experiência de Produto Completa.
    - Comandos de DX: `init` (Onboarding) e `inspect` (Schema Discovery).
    - Gestão de Estado Global: Geração automática de `versions.tf` e Backend S3 dinâmico.
    - Isolamento de Outputs: Estrutura organizada em `infra/{resource}/`.
- [ ] **v0.4 (Próximo):** Expansão de Cobertura (The "Big Five").
    - Adicionar suporte nativo (Fetchers) para:
      1. **EC2** (Compute Instances)
      2. **ECS** (Serverless Functions)
      3. **ElastiCache** (Redis/Memcached)
      4. **SQS** (Message Queues)
      5. **VPC** (Network Rules)
- [ ] **v0.5:** Engenharia Reversa Total.
    - Gerar o `exevo.yaml` automaticamente a partir do comando `inspect`.
    - Importar infraestrutura legada inteira com um único comando.
- [ ] **v1.0:** Plugin System (Go Plugins) e suporte a Azure/GCP.

## 🤝 Contribuindo
Pull Requests são bem-vindos! Para mudanças maiores, abra uma issue primeiro para discutir o que você gostaria de mudar.

## 📄 Licença
MIT