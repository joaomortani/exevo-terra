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

# 🚀 Instalação

### Opção 1: Via Go (Para Desenvolvedores)
Se você já tem o Go instalado:

```bash
go install [github.com/joaomortani/exevo-terra@latest](https://github.com/joaomortani/exevo-terra@latest)
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
   exevo-terra docs --resource rds
   ```

4. **Gere o código:**
   Conecte na AWS e gere os arquivos Terraform:
   ```bash
   exevo-terra generate --resource rds --profile default
   ```

## 🗺️ Roadmap & Futuro
> O Exevo Terra está evoluindo de uma ferramenta "Opinionated" para um motor de IaC genérico.

- [x] **v0.1 (Atual):** Suporte focado em RDS com módulos padrão.
- [x] **v0.2 (Atual):** Arquitetura "Bring Your Own Module" (BYOM).
    - Suporte a configuração via YAML (`exevo.yaml`).
    - Mapeamento dinâmico de campos da AWS para Variáveis do Terraform.
    - Independência de Provider (suporte futuro a S3, ElastiCache, etc).
- [ ] **v1.0:** Plugin System e suporte a múltiplos Providers de Cloud.

## 🤝 Contribuindo
Pull Requests são bem-vindos! Para mudanças maiores, abra uma issue primeiro para discutir o que você gostaria de mudar.

## 📄 Licença
MIT