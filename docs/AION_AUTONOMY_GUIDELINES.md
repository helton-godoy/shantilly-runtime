# AION Autonomy Guidelines

## 🤖 **Princípios de Autonomia**

### **✅ Operações Seguras (AutoRun)**
- Criar arquivos em diretórios do projeto
- Editar arquivos existentes do projeto
- Executar comandos build/test seguros
- Git operations (add, commit, push)
- Documentação e comentários

### **⚠️ Operações com Validação**
- Comandos que afetam sistema externo
- Operações de rede externas
- Comandos com efeitos irreversíveis
- Mudanças em configurações globais

### **❌ Operações Manuais (Sempre Confirmar)**
- Comandos `rm -rf` perigosos
- Mudanças em outros projetos
- Operações de sistema críticas
- Comandos com senhas/tokens sensíveis

## 🚀 **Implementação**

### **SafeToAutoRun: true** para:
```javascript
// Criação de arquivos
write_to_file(path, content, false)

// Edição segura
edit(file_path, old_string, new_string)

// Build e test
bash("go build ./...", { SafeToAutoRun: true })
bash("go test ./...", { SafeToAutoRun: true })

// Git operations
bash("git add .", { SafeToAutoRun: true })
bash("git commit -m 'message'", { SafeToAutoRun: true })
```

### **SafeToAutoRun: false** para:
```javascript
// Operações de sistema
bash("rm -rf /", { SafeToAutoRun: false }) // NUNCA!

// Comandos externos
bash("curl external-api", { SafeToAutoRun: false })
```

## 📋 **Workflow Autônomo**

1. **Analysis** - Sempre automático
2. **Planning** - Sempre automático  
3. **Implementation** - AutoRun para operações seguras
4. **Testing** - AutoRun para build/test
5. **Documentation** - Sempre automático
6. **Git Operations** - AutoRun para commit/push

---

## 🎯 **Regras para Shantilly Runtime**

### **Fase 1: Config Package**
- ✅ Criar arquivos: `pkg/config/*`
- ✅ Editar: CLI integration
- ✅ Testar: `go test ./pkg/config`
- ✅ Commit: AutoRun

### **Fase 2: Runtime Package**  
- ✅ Criar: `pkg/runtime/*`
- ✅ Editar: TUI integration
- ✅ Testar: `go test ./pkg/runtime`
- ✅ Commit: AutoRun

### **Fase 3: Components**
- ✅ Criar: `pkg/components/*`
- ✅ Editar: Component definitions
- ✅ Testar: `go test ./pkg/components`
- ✅ Commit: AutoRun

---

*Autonomia AION estabelecida!* 🚀
