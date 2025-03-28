mock:
	mockgen -package=template -source=app/template/service.go -destination=app/template/service_mocks_test.go

run:
	go run .

