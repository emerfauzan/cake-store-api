name = all
run:
	go run main.go run
migrate-create:
	go run main.go migrate create -n $(name)
migrate-up:
	go run main.go migrate up
migrate-down:
	go run main.go migrate down
migrate-status:
	go run main.go migrate status