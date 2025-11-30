# Contributing to Shantilly Runtime

Thank you for your interest in contributing to Shantilly Runtime! This project is being developed autonomously by the AION framework, but human contributions are welcome.

## 🤖 Development Approach

This project uses **AION (AI Orchestration Native)** for autonomous development:
- **PM Agent**: Requirements and planning
- **Architect Agent**: System design and architecture
- **Developer Agent**: Code implementation
- **QA Agent**: Testing and validation
- **Release Agent**: Deployment and monitoring

## 📋 How to Contribute

### For Human Contributors

1. **Fork the repository**
2. **Create a feature branch**: `git checkout -b feature/amazing-feature`
3. **Make your changes** following the project patterns
4. **Add tests** for your changes
5. **Run the test suite**: `go test ./...`
6. **Commit your changes**: `git commit -m 'Add amazing feature'`
7. **Push to branch**: `git push origin feature/amazing-feature`
8. **Open a Pull Request**

### For AION Framework

AION agents will automatically:
- Analyze requirements and create PRDs
- Design architecture and technical specs
- Implement features with comprehensive testing
- Validate quality and performance
- Prepare releases with documentation

## 🛠️ Development Setup

### Prerequisites
- Go 1.21 or higher
- Git
- Make (optional, for build scripts)

### Setup
```bash
# Clone the repository
git clone https://github.com/helton-godoy/shantilly-runtime.git
cd shantilly-runtime

# Install dependencies
go mod download

# Run tests
go test ./...

# Build the project
go build ./...
```

## 📝 Code Style

- Follow Go conventions and best practices
- Use `gofmt` for code formatting
- Add comprehensive tests for new features
- Document public functions and types
- Keep functions small and focused

## 🧪 Testing

- Run `go test ./...` to run all tests
- Use `go test -race ./...` for race condition detection
- Maintain >80% test coverage
- Add integration tests for complex features

## 📚 Documentation

- Update README.md for user-facing changes
- Add code comments for complex logic
- Update API documentation for public interfaces
- Maintain bilingual support (pt-BR/en)

## 🚀 Release Process

Releases are managed by the AION Release Agent:
1. Automated testing and validation
2. Version bumping and changelog
3. GitHub release creation
4. Documentation updates
5. Deployment announcements

## 🤝 Community Guidelines

- Be respectful and constructive
- Focus on technical discussions
- Help others learn and contribute
- Follow the project's bilingual approach (pt-BR/en)

## 📞 Getting Help

- Open an **Issue** for bug reports or questions
- Start a **Discussion** for general topics
- Check existing issues before creating new ones
- Join our community discussions

---

*This contributing guide is maintained by both human contributors and the AION framework* 🚀