DC = UID=$(shell id -u) docker compose

.PHONY: dev exe tidy shell clean
dev:   ## Arranca la app en http://localhost:8787
	$(DC) up --build
tidy:  ## Resuelve dependencias (go mod tidy)
	$(DC) run --rm dev go mod tidy
exe:   ## Compila dist/floreria.exe para Windows 10 (64 bits)
	$(DC) run --rm -e CGO_ENABLED=0 -e GOOS=windows -e GOARCH=amd64 dev \
		go build -trimpath -ldflags="-s -w" -o dist/floreria.exe .
shell: ## Terminal dentro del contenedor
	$(DC) run --rm dev bash
clean: ## Borra TODO: contenedores, volúmenes, imagen, datos y binarios
	$(DC) down -v --rmi local --remove-orphans
	rm -rf data dist
