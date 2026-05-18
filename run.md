# AI Classification Engine

## Requirements

- Go installed
- Ollama installed
- llama3 model installed

---

# Install llama3

```powershell
ollama run llama3
```

---

# Run project

```powershell
go run demo.go llm.go classifier.go models.go
```

---

# Expected Output

```text
AI Classification Engine Demo
====================================

INPUT:
My password is 123456

CLASSIFICATION:
{
  "category": "credential",
  "sensitivity": "high",
  "risk": "critical",
  "reason": "Contains password information"
}
```