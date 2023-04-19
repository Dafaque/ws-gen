ver:
	@git describe --tags --abbrev=0 > assets/VERSION.txt

srun:
	@go run ./examples/server/main.go
crun:
	@go run ./examples/client/main.go