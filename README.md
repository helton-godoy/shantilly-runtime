# Shantilly Runtime

[![Build and Test](https://github.com/helton-godoy/shantilly-runtime/actions/workflows/ci.yml/badge.svg)](https://github.com/helton-godoy/shantilly-runtime/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.21-00ADD8.svg?logo=go)](https://golang.org/doc/devel/release)
[![AION Orchestration](https://img.shields.io/badge/AION-Orchestrated-blue)](https://github.com/helton-godoy/aion)

> **Runtime TUI Declarativo - Versão limpa e otimizada com orquestração AION**

Shantilly Runtime é uma implementação focada e limpa do runtime TUI declarativo, desenvolvida 100% pelo framework AION (AI Orchestration Native).

## 🌐 Escolha seu idioma · Choose your language

- 🇧🇷 **Português (PT-BR)**: [Documentação abaixo](#documentação-em-português)
- 🌍 **English (EN)**: [English Documentation](#english-documentation)

---

## 📋 Documentação em Português

### 🎯 Propósito

Este projeto é uma reimplementação limpa e otimizada do Shantilly, focada em:

- ✨ **Abordagem Declarativa**: YAML-driven para criação de TUI
- 🏗️ **Ecossistema Moderno**: Go + Bubbletea ecosystem
- 🤖 **Orquestração AION**: Desenvolvimento autônomo completo
- 📦 **Binário Único**: Portátil para Linux, macOS, Windows
- 🛡️ **Segurança por Padrão**: Validação JIT e sandbox

### 🚀 Características Principais

- **Declarativo**: Defina UIs complexas com YAML simples
- **Componentes Ricos**: Forms, seletores, layouts, mouse support
- **Event-Driven**: Sistema de eventos robusto e extensível
- **Shell Integration**: Perfeito para scripts e automação
- **Cross-Platform**: Linux, macOS, Windows

### 🔄 Desenvolvimento AION

Este projeto é desenvolvido autonomamente pelo framework AION:

1. **Product Manager Agent**: Análise de requisitos e planejamento
2. **Solution Architect Agent**: Design de arquitetura e especificações técnicas
3. **Developer Agent**: Implementação de código e testes
4. **Quality Assurance Agent**: Validação e garantia de qualidade
5. **Release Manager Agent**: Deploy e monitoramento

### 📊 Status do Projeto

- **Framework**: AION v1.0.0-alpha
- **Orquestração**: 100% autônoma
- **Stack**: Go 1.24+ com Charm ecosystem
- **Target**: Production-ready em 6 semanas
- **Quality**: Test-driven development com CI/CD

---

## 📚 English Documentation

### 🎯 Purpose

This project is a clean, focused reimplementation of Shantilly, emphasizing:

- ✨ **Declarative Approach**: YAML-driven TUI creation
- 🏗️ **Modern Ecosystem**: Go + Bubbletea ecosystem
- 🤖 **AION Orchestration**: Fully autonomous development
- 📦 **Single Binary**: Portable for Linux, macOS, Windows
- 🛡️ **Secure by Default**: JIT validation and sandbox

### 🚀 Key Features

- **Declarative**: Define complex UIs with simple YAML
- **Rich Components**: Forms, selectors, layouts, mouse support
- **Event-Driven**: Robust and extensible event system
- **Shell Integration**: Perfect for scripts and automation
- **Cross-Platform**: Linux, macOS, Windows

### 🔄 AION Development

This project is autonomously developed by the AION framework:

1. **Product Manager Agent**: Requirements analysis and planning
2. **Solution Architect Agent**: Architecture design and technical specs
3. **Developer Agent**: Code implementation and testing
4. **Quality Assurance Agent**: Validation and quality assurance
5. **Release Manager Agent**: Deployment and monitoring

### 📊 Project Status

- **Framework**: AION v1.0.0-alpha
- **Orchestration**: 100% autonomous
- **Stack**: Go 1.24+ with Charm ecosystem
- **Target**: Production-ready in 6 weeks
- **Quality**: Test-driven development with CI/CD

---

## 🚀 Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/helton-godoy/shantilly-runtime.git
cd shantilly-runtime

# Build the project
go build ./...

# Run the application
./shantilly-runtime --help
```

### Basic Usage

```yaml
# example.yaml
app:
  name: "My TUI App"
  version: "1.0.0"

layout:
  type: "vertical"
  components:
    - type: "text"
      content: "Welcome to Shantilly Runtime!"
    
    - type: "button"
      label: "Click Me"
      on_click:
        type: "run"
        command: "echo 'Button clicked!'"
```

```bash
# Run with YAML configuration
./shantilly-runtime -f example.yaml
```

---

## 🛠️ Development

### Prerequisites

- Go 1.21 or higher
- Git
- Make (optional)

### Setup

```bash
# Install dependencies
go mod download

# Run tests
go test ./...

# Build
go build ./...

# Run with coverage
go test -race -coverprofile=coverage.out ./...
```

### Contributing

Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on contributing to this project.

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🤝 Community

- **Issues**: [Report bugs and request features](https://github.com/helton-godoy/shantilly-runtime/issues)
- **Discussions**: [General topics and questions](https://github.com/helton-godoy/shantilly-runtime/discussions)
- **Contributing**: [See contributing guidelines](CONTRIBUTING.md)

---

*Generated by AION Autonomous Development System* 🚀