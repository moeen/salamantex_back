package server


func InitServer() {
	router := Router()

	router.Run()
}
