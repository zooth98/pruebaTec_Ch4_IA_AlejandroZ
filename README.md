## PRUEBA TECNICA - RETO 4
Jovan Alejandro Zambrano Bello


Este proyecto implementa un motor de clasificación basado en inteligencia artificial utilizando Go y un modelo de lenguaje grande (LLM) compatible con OpenAI (Ollama + Llama3).

El motor clasifica muestras de texto según:
- categoría
- sensibilidad
- riesgo

Además, incluye una validación híbrida basada en reglas para documentos de identificación latinoamericanos de carácter sensible.

## ARQUITECTURA:

<img width="1357" height="387" alt="image" src="https://github.com/user-attachments/assets/4407cd02-b283-4ce1-9345-2ed6fa8ab7aa" />

## DECISIONES TÉCNICAS:

- Se eligió Go por su simplicidad, su compatibilidad con la ejecución concurrente y sus potentes capacidades HTTP.
- Se utilizó Ollama como proveedor local de modelos de lenguaje grande (LLM) compatibles con OpenAI.
- Las indicaciones se separaron en el archivo `system_prompt.md` para facilitar el mantenimiento y la revisión.
- Se implementó una estrategia de clasificación híbrida:
  - Validación basada en reglas para los documentos de identidad de Latinoamérica.
  - Clasificación semántica mediante modelos de lenguaje grande (LLM) para la comprensión contextual.

## REQUISITOS

- Go 1.25+
- Ollama installed
- llama3 model installed

## INSTALACIÓN

ollama run llama3

## EJECUCIÓN

go run demo.go classifier.go llm.go models.go
