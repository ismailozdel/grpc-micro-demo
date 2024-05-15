package global

import (
	"crypto/md5"
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var params = make(map[string]interface{})

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type UpdateResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    any    `json:"data"`
}

type PatchResponse struct {
	Success []UpdateResponse `json:"success"`
	Fail    []UpdateResponse `json:"fail"`
}

func SetParam(key string, value interface{}) {
	params[key] = value
}

func GetParam(key string) interface{} {

	return params[key]
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDbConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Println(err.Error())
	}
}

// Test Git

func CloseGormConnection(ctx *fiber.Ctx) {

	db := ctx.Locals("db")
	if mdb, ok := db.(*gorm.DB); ok {
		_conn, _ := mdb.DB()
		CloseDbConnection(_conn)
	}
}

func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Println(err.Error())
	}
}

// verilen string değeri sha1 ile hashler
func Sha1(val *string) {
	h := sha1.New()
	h.Write([]byte(*val))
	*val = hex.EncodeToString(h.Sum(nil))

}

func IsNumeric(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

func ConvertToMap(structData interface{}) map[string]interface{} {
	jsonData, _ := json.Marshal(structData)
	var mapData map[string]interface{}
	_ = json.Unmarshal(jsonData, &mapData)

	return mapData
}

// Postgrenin döndürdüğü hata kodunu verir
func GetErrorCode(err error) string {
	re := regexp.MustCompile(`SQLSTATE (\w+)`)
	match := re.FindStringSubmatch(err.Error())
	if len(match) > 0 {
		return match[1]
	}

	return "-1"
}

func SendResponse(ctx *fiber.Ctx, err error, data interface{}, status ...int) error {

	return ctx.JSON(&Response{Code: 0, Message: "Ok", Data: data})
}

func UrlQueryToSqlCriteria(queries map[string]string) string {
	criteria := strings.Builder{}

	for key, value := range queries {
		if criteria.Len() > 0 {
			criteria.WriteString(" AND ")
		}

		if strings.Contains(value, ",") {
			v := "'" + strings.Join(strings.Split(value, ","), "','") + "'"
			criteria.WriteString(key + " IN (" + v + ")")
		} else if strings.Contains(key, "<>") {
			k := strings.Split(key, "<>")
			criteria.WriteString(k[0] + " <> '" + k[1] + "'")
		} else if strings.Contains(key, "<") {
			k := strings.Split(key, "<")
			criteria.WriteString(k[0] + " < '" + k[1] + "'")
		} else if strings.Contains(key, ">") {
			k := strings.Split(key, ">")
			criteria.WriteString(k[0] + " > '" + k[1] + "'")
		} else {
			criteria.WriteString(key + " = '" + value + "'")
		}
	}

	return criteria.String()
}

func Md5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// implement edilecek
func GenerateToken(id int64) string {
	return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWQiOiIxMiIsImlhdCI6MTUxNjIzOTAyMn0.-qcBUJh-U8SZFu9racf6P7lo3diw3RZAQWSSWkpY8-E"
}

var ErrRecordNotFound = errors.New("record not found")

// func HaciPhotoCheck(received, current string) (string, error) {
//     // received ve current boş değilse silme işlemi yapılabilir
//     shouldDelete := received != "" && current != ""

//     // Eğer received boş ise, current'ı döndür ve hata olmadığını belirt
//     if received == "" {
//         return current, nil
//     }

//     // Eğer received ve current farklı ise, yeni dosyayı kaydet
//     if received != current {
//         file := FileBase64{
//             File64: received,
//         }
//         // Yeni dosyayı kaydet
//         if err := file.SaveToStorage(); err != nil {
//             return "", err
//         }
// 		// Silme işlemi gerçekleştir
// 		if shouldDelete {
// 			DeleteFromStorage(current)
// 		}
//         // received değerini güncelle
//         received = file.FilePath
//     }

//     // received değerini döndür ve hata olmadığını belirt
//     return received, nil
// }
