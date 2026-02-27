# IPTV Besser - Gestión Canales TUI

## Activar githooks
```bash
# Dar permisos de ejecución a los scripts
chmod +x .githooks/pre-commit .githooks/commit-msg

# Decirle a git que use esta carpeta para los hooks
git config core.hooksPath .githooks
```

Instalar golangci-lint para el linter

`curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.10.1`

## Instalar gofumpt para formatear el código

```...```
