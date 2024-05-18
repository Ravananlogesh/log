package common

// import (
// 	"bytes"
// 	"encoding/base64"
// 	"fmt"
// 	"html/template"
// 	"log"
// 	"math/rand"
// 	"api-log/ftdb"
// 	"strings"
// 	"time"

// 	"github.com/BurntSushi/toml"
// 	_ "github.com/denisenkom/go-mssqldb"
// 	_ "github.com/go-sql-driver/mysql"

// 	"github.com/google/uuid"
// 	util "github.com/mrlakshmanan/fcsutility2"
// )

// func ReadTomlConfig(filename string) interface{} {
// 	var f interface{}
// 	if _, err := toml.DecodeFile(filename, &f); err != nil {
// 		log.Println(err)
// 	}
// 	return f
// }

// func GenerateOTP() string {
// 	log.Println("GenerateOTP+")
// 	rand.Seed(time.Now().UnixNano())
// 	b := rand.Intn(1000000)
// 	otp := fmt.Sprintf("%06d", b)
// 	//log.Println(token)
// 	log.Println("GenerateOTP-")
// 	return otp
// }

// // --------------------------------------------------------------
// // method used to get values from coresettings
// // for given  input
// // --------------------------------------------------------
// func GetCoreSettingValue(Key string) (string, error) {
// 	log.Println("GetCoreSettingValue(+)")
// 	var Value string
// 	//open a db connection
// 	MariaEkycPRD, err := ftdb.LocalDbConnect(ftdb.MariaEKYCPRD)
// 	//if any error when opening db connection
// 	if err != nil {
// 		LogDebug("common.GetCoreSettingValue ", "(CGCV01)", err.Error())
// 		return Value, err
// 	} else {
// 		defer MariaEkycPRD.Close()
// 		// Get the value from CoreSetting.
// 		Value = util.GetCoreSettingValue(MariaEkycPRD, Key)
// 		log.Println(Value)

// 	}
// 	log.Println("GetCoreSettingValue(-)")
// 	return Value, nil
// }

// // ---------------------------------------------------------------
// // function Encrypts the given mailid
// // Returns the Encryted Mailid
// // ------------------------------------------------------------
// func GetEncryptedemail(emailId string) (string, error) {
// 	log.Println("GetEncryptedemail(+)")
// 	idxStart := 0
// 	idxMiddle := strings.Index(emailId, "@")
// 	log.Println(" idxStart:", idxStart)
// 	log.Println(" idxEnd:", idxMiddle)
// 	// for printinf first letter
// 	firstLetter := emailId[idxStart : idxStart+1]
// 	log.Println(firstLetter)

// 	firstHalf := emailId[idxStart+1 : idxMiddle]

// 	for i := range firstHalf {

// 		firstHalf = strings.Replace(firstHalf, string(firstHalf[i]), "*", 1)
// 	}
// 	// idxMiddle := strings.Index(emailId, "@")
// 	idxEnd := len(emailId)
// 	log.Println(" idxStart1:", idxMiddle)
// 	log.Println(" idxEnd1:", idxEnd)

// 	letterAfterAt := emailId[idxMiddle+2 : idxEnd]
// 	log.Println(letterAfterAt)
// 	// for printing first letter
// 	SecondHalf := emailId[idxMiddle : idxMiddle+2]

// 	for j := range letterAfterAt {
// 		//log.Println(j)
// 		if string(letterAfterAt[j]) == "." {
// 			break
// 		}
// 		letterAfterAt = strings.Replace(letterAfterAt, string(letterAfterAt[j]), "*", 1)

// 	}

// 	encrytedEmaiId := firstLetter + firstHalf + SecondHalf + letterAfterAt
// 	log.Println(encrytedEmaiId)
// 	log.Println("GetEncryptedemail(-)")
// 	return encrytedEmaiId, nil
// }

// // ---------------------------------------------------------------
// // function Encrypts the given MobileNumber
// // Returns the Encryted MobileNumber
// // ------------------------------------------------------------
// func GetEncryptedMobile(mobileNumber string) (string, error) {
// 	log.Println("GetEncryptedMobile(+)")
// 	for K := range mobileNumber {
// 		if K == len(mobileNumber)-4 {
// 			break
// 		}
// 		mobileNumber = strings.Replace(mobileNumber, string(mobileNumber[K]), "*", 1)
// 	}
// 	//log.Println(mobileNumber)
// 	log.Println("GetEncryptedMobile(-)")
// 	return mobileNumber, nil

// }

// func LogError(who string, ref string, msg string) {
// 	log.Println(who, "ERROR:("+ref+")", msg)
// }

// func LogDebug(who string, ref string, msg string) {
// 	log.Println("DEBUG: ", who, ref, msg)
// }

// func GetFileName_UUID_String() string {
// 	var id = uuid.New()
// 	return id.String()
// }

// // -------------------------------------
// // Method encodes the given input
// // Returns the data in string
// // --------------------------------------
// func EncodeToString(fileBody string) string {
// 	log.Println("EncodeToString(+)")

// 	encodedText := base64.StdEncoding.EncodeToString([]byte(fileBody))
// 	//log.Println("Encoded text:", encodedText)
// 	log.Println("EncodeToString(-)")

// 	return encodedText
// }

// // -------------------------------------
// // Method decode the given input
// // Returns the data in string
// // --------------------------------------
// func DecodeToString(fileBody string) (string, error) {
// 	log.Println("DecodeToString(+)")
// 	decodeText, err := base64.StdEncoding.DecodeString(fileBody)
// 	if err != nil {
// 		log.Println("Encoded text:", decodeText)
// 		LogDebug("common.DecodeToString ", "(CDS01)", err.Error())
// 		return string(decodeText), err
// 	}

// 	log.Println("DecodeToString(-)")
// 	return string(decodeText), nil
// }

// // ------------------------------------------
// // Method read html file and return the
// // file data as String
// // --------------------------------------------
// func HtmlFileToString(fileName string) (string, error) {
// 	log.Println("HtmlFileToString+")

// 	var htmlString string
// 	var tpl bytes.Buffer
// 	temp, err := template.ParseFiles(fileName) // change this
// 	if err != nil {
// 		LogDebug("common.HtmlFileToString ", "(CHFS01)", err.Error())
// 		return htmlString, err
// 	} else {
// 		temp.Execute(&tpl, "")
// 		htmlString = tpl.String()
// 	}

// 	log.Println("HtmlFileToString-")
// 	return htmlString, nil

// }

// func GetLoggedBy(ClientId string) string {
// 	log.Println("GetLoggedBy+")

// 	array := strings.Split(ClientId, ",")

// 	log.Println("GetLoggedBy-")
// 	return array[0]
// }

// func GetSetClient(ClientId string) string {
// 	log.Println("GetSetClient+")

// 	array := strings.Split(ClientId, ",")

// 	log.Println("GetSetClient-")
// 	return array[0]
// }