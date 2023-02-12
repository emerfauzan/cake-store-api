name = all
run:
	go run app/app.go
seed:
	go run main.go seed
migrate-create:
	go run main.go migrate create -n $(name)
migrate-up:
	go run main.go migrate up
migrate-down:
	go run main.go migrate down
migrate-status:
	go run main.go migrate status