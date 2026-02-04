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

## 🚀 Features (v0.1)

- **Discovery Automático**: Varre sua conta AWS via SDK v2 e encontra recursos (Foco atual: RDS).
- **Geração de HCL**: Cria arquivos `.tf` formatados e prontos para uso.
- **State Binding Automático**: Gera blocos `import { ... }` (compatível com Terraform 1.5+) para evitar conflitos de criação.
- **SSO Nativo**: Suporte transparente para autenticação via AWS SSO (`aws sso login`).
- **Null Safety**: Camada de adaptação robusta que protege contra falhas de ponteiros da API da AWS.

## 📦 Instalação

```bash
# Via Go Install (Recomendado)
go install [github.com/joaomortani/exevo-terra@latest](https://github.com/joaomortani/exevo-terra@latest)

# Verifique a instalação
exevo-terra --help
```

## ⚡ Quick Start

1. Listar Recursos (Dry Run)
Veja o que o Exevo Terra consegue enxergar na sua conta:

```bash
exevo-terra rds list --region us-east-1 --profile meu-perfil-sso
```

2. Gerar Código e Imports
Gere os arquivos .tf e imports.tf para trazer os recursos para o seu state:

```bash
exevo-terra rds generate --filter "nome-do-app" --profile meu-perfil-sso
```

Isso criará:

rds.tf: A definição do módulo.

imports.tf: O mapeamento para o Terraform importar o state.

3. Aplicar
```bash
terraform init
terraform plan # Verifique se o plan indica "Importing..."
terraform apply
```

## 🗺️ Roadmap & Futuro
> O Exevo Terra está evoluindo de uma ferramenta "Opinionated" para um motor de IaC genérico.

- [x] **v0.1 (Atual):** Suporte focado em RDS com módulos padrão.
- [ ] **v0.2 (Em Desenvolvimento):** Arquitetura "Bring Your Own Module" (BYOM).
    - Suporte a configuração via YAML (`exevo.yaml`).
    - Mapeamento dinâmico de campos da AWS para Variáveis do Terraform.
    - Independência de Provider (suporte futuro a S3, ElastiCache, etc).
- [ ] **v1.0:** Plugin System e suporte a múltiplos Providers de Cloud.

## 🤝 Contribuindo
Pull Requests são bem-vindos! Para mudanças maiores, abra uma issue primeiro para discutir o que você gostaria de mudar.
