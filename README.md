# Floristería · gestión local

App local (Go + SQLite) para pequeños negocios. Todo el desarrollo corre en Docker.

| Comando | Qué hace |
|---|---|
| `make dev` | Arranca en http://localhost:8787 |
| `make exe` | Compila `dist/floreria.exe` (Windows 64 bits) |
| `make tidy` | Actualiza dependencias |
| `make shell` | Terminal dentro del contenedor |
| `make clean` | Borra contenedores, volúmenes, imagen, `data/` y `dist/` |

Para eliminar todo rastro: `make clean` y borrar la carpeta del repo.
