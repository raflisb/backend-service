module github.com/raflisb/backend-service

go 1.16

require (
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.2
	github.com/joho/godotenv v1.3.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.8
)

//replace ./github.com/biezhi/gorm-paginator/pagination => ./github.com/samratsuzil/gorm-paginator
